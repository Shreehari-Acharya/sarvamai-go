package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/errors"
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
		return errors.ParseAPIError(resp)
	}

	if result != nil && resp.ContentLength != 0 {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("decode response: %w", err)
		}
	}

	return nil
}
