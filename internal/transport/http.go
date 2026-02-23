package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

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

	fmt.Printf("Making request to %s with body: %s\n", req.URL, reader)

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
	fields map[string]string,
	result any,
) error {

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

		for key, value := range fields {
			if err := writer.WriteField(key, value); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
	}()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, t.BaseURL+path, pr)
	if err != nil {
		return fmt.Errorf("create multipart request: %w", err)
	}

	req.Header.Set("api-subscription-key", t.APIKey)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := t.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("execute multipart request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return sarvamaierrors.ParseAPIError(resp)
	}

	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil && err != io.EOF {
			return fmt.Errorf("decode response: %w", err)
		}
	}

	return nil
}
