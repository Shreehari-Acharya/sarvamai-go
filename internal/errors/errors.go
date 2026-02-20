package errors

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// APIError represents a structured error returned by Sarvam AI.
// It implements the error interface.
type APIError struct {
	StatusCode int    // HTTP status code (e.g., 400, 401, 500)
	Message    string `json:"message"` // Human-readable error message
	Code       string `json:"code"`    // Optional API-specific error code
	RawBody    string // Raw response body for debugging purposes
}

type DecodedAPIError struct {
	Detail string `json:"detail"`
}

// Error satisfies the error interface.
func (e *APIError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("sarvamai API error (%d): %s", e.StatusCode, e.Message)
	}
	return fmt.Sprintf("sarvamai API error (%d)", e.StatusCode)
}

// ParseAPIError reads and parses an error HTTP response.
// It ensures the response body is fully read so that
// the underlying HTTP connection can be safely reused.
func ParseAPIError(resp *http.Response) error {

	// Read the entire body to ensure connection reuse.
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read error response body: %w", err)
	}

	apiErr := &APIError{
		StatusCode: resp.StatusCode,
		RawBody:    string(bodyBytes),
	}

	var decoded DecodedAPIError

	if err := json.Unmarshal(bodyBytes, &decoded); err == nil {
		if decoded.Detail != "" {
			apiErr.Message = decoded.Detail
		}
	}

	// Even if JSON decoding fails,
	// we still return a structured error with RawBody populated.
	return apiErr
}
