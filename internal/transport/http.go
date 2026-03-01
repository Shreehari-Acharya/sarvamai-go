package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"reflect"

	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/sarvamaierrors"
)

type Transport struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
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

	var reader io.Reader

	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("marshal request body: %w", err)
		}
		reader = bytes.NewBuffer(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, t.BaseURL+path, reader)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("api-subscription-key", t.APIKey)

	if contentType == "" {
		contentType = "application/json"
	}
	req.Header.Set("Content-Type", contentType)

	resp, err := t.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("execute request: %w", err)
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("failed to close response body: %v\n", err)
		}
	}()

	if resp.StatusCode >= 400 {
		return sarvamaierrors.ParseAPIError(resp)
	}

	if result != nil && resp.ContentLength != 0 {
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

	var reader io.Reader

	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal request body: %w", err)
		}
		reader = bytes.NewBuffer(b)
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

	pr, pw := io.Pipe()
	writer := multipart.NewWriter(pw)

	go func() {
		defer pw.Close()
		defer writer.Close()

		// Write file part
		part, err := writer.CreateFormFile(fileFieldName, fileName)
		if err != nil {
			pw.CloseWithError(err)
			return
		}

		if _, err := io.Copy(part, file); err != nil {
			pw.CloseWithError(err)
			return
		}

		// Write struct fields
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

				// Handle pointer and interface types safely
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
