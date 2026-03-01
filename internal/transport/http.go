package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"mime/multipart"
	"net/http"
	"reflect"
	"time"

	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/sarvamaierrors"
)

const (
	defaultMaxRetries = 3
	defaultMinWait    = 1 * time.Second
	defaultMaxWait    = 10 * time.Second
)

type Transport struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
	MaxRetries int
}

// doWithRetry is a helper that executes a function with exponential backoff retries.
func (t *Transport) doWithRetry(ctx context.Context, operation func() (*http.Response, error)) (*http.Response, error) {
	maxRetries := t.MaxRetries
	if maxRetries <= 0 {
		maxRetries = defaultMaxRetries
	}

	var lastResp *http.Response
	var lastErr error

	for i := 0; i <= maxRetries; i++ {
		// If this is a retry, wait before trying again
		if i > 0 {
			wait := time.Duration(math.Pow(2, float64(i-1))) * defaultMinWait
			if wait > defaultMaxWait {
				wait = defaultMaxWait
			}

			timer := time.NewTimer(wait)
			select {
			case <-ctx.Done():
				timer.Stop()
				if lastResp != nil {
					lastResp.Body.Close()
				}
				return nil, ctx.Err()
			case <-timer.C:
			}
		}

		resp, err := operation()
		if err != nil {
			// Always retry on network errors
			lastErr = err
			continue
		}

		// Check if we should retry based on status code
		if resp.StatusCode == http.StatusTooManyRequests || (resp.StatusCode >= 500 && resp.StatusCode <= 599) {
			lastResp = resp
			lastErr = sarvamaierrors.ParseAPIError(resp)
			continue
		}

		// Success or non-retryable error
		return resp, nil
	}

	return lastResp, lastErr
}

func (t *Transport) DoRequest(
	ctx context.Context,
	method string,
	path string,
	body any,
	result any,
	contentType string,
) error {

	if err := ctx.Err(); err != nil {
		return err
	}

	var bodyBytes []byte
	if body != nil {
		var err error
		bodyBytes, err = json.Marshal(body)
		if err != nil {
			return fmt.Errorf("marshal request body: %w", err)
		}
	}

	resp, err := t.doWithRetry(ctx, func() (*http.Response, error) {
		var reader io.Reader
		if bodyBytes != nil {
			reader = bytes.NewReader(bodyBytes)
		}

		req, err := http.NewRequestWithContext(ctx, method, t.BaseURL+path, reader)
		if err != nil {
			return nil, err
		}

		req.Header.Set("api-subscription-key", t.APIKey)
		if contentType == "" {
			contentType = "application/json"
		}
		req.Header.Set("Content-Type", contentType)

		return t.HTTPClient.Do(req)
	})

	if err != nil {
		return err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("failed to close response body: %v\n", err)
		}
	}()

	if resp.StatusCode >= 400 {
		return sarvamaierrors.ParseAPIError(resp)
	}

	if result != nil && resp.StatusCode != http.StatusNoContent {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("decode response: %w", err)
		}
	}

	return nil
}

func (t *Transport) DoStreamRequest(
	ctx context.Context,
	method string,
	path string,
	body any,
) (*http.Response, error) {

	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	var bodyBytes []byte
	if body != nil {
		var err error
		bodyBytes, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal request body: %w", err)
		}
	}

	// We don't retry streaming requests because they are long-lived
	// and retrying might result in duplicate connections.
	var reader io.Reader
	if bodyBytes != nil {
		reader = bytes.NewReader(bodyBytes)
	}

	req, err := http.NewRequestWithContext(ctx, method, t.BaseURL+path, reader)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("api-subscription-key", t.APIKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Cache-Control", "no-cache")

	resp, err := t.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("execute request: %w", err)
	}

	if resp.StatusCode >= 400 {
		err := sarvamaierrors.ParseAPIError(resp)
		resp.Body.Close()
		return nil, err
	}

	return resp, nil
}

func (t *Transport) DoMultipartRequest(
	ctx context.Context,
	path string,
	fileFieldName string,
	fileName string,
	file io.Reader,
	formStruct any,
	result any,
) error {

	// Note: We don't automatically retry multipart requests because the file reader
	// might not be seekable (e.g., a pipe or a socket), so we can't rewind it for a retry.
	// Production SDKs typically require an io.ReadSeeker if retries are needed for uploads.

	pr, pw := io.Pipe()
	writer := multipart.NewWriter(pw)

	go func() {
		defer pw.Close()
		defer writer.Close()

		part, err := writer.CreateFormFile(fileFieldName, fileName)
		if err != nil {
			pw.CloseWithError(err)
			return
		}

		if _, err := io.Copy(part, file); err != nil {
			pw.CloseWithError(err)
			return
		}

		if formStruct != nil {
			v := reflect.ValueOf(formStruct)
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}

			tp := v.Type()
			for i := 0; i < v.NumField(); i++ {
				fieldValue := v.Field(i)
				fieldType := tp.Field(i)

				tag := fieldType.Tag.Get("form")
				if tag == "" {
					continue
				}

				val := fieldValue
				if val.Kind() == reflect.Ptr || val.Kind() == reflect.Interface {
					if val.IsNil() {
						continue
					}
					val = val.Elem()
				}

				value := fmt.Sprintf("%v", val.Interface())
				if err := writer.WriteField(tag, value); err != nil {
					pw.CloseWithError(err)
					return
				}
			}
		}
	}()

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		t.BaseURL+path,
		pr,
	)
	if err != nil {
		return err
	}

	req.Header.Set("api-subscription-key", t.APIKey)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := t.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return sarvamaierrors.ParseAPIError(resp)
	}

	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return err
		}
	}

	return nil
}
