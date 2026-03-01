package sarvamaierrors

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// APIError represents a structured error returned by Sarvam AI.
// It implements the error interface.
type APIError struct {
	StatusCode int    // HTTP status code (e.g., 400, 401, 500)
	Message    string `json:"message"` // Human-readable error message
	Code       string `json:"code"`    // Optional API-specific error code
	RawBody    string // Raw response body for debugging purposes
}

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// Error satisfies the error interface.
func (e *APIError) Error() string {
	msg := e.Message
	if msg == "" {
		msg = "unknown error"
	}

	if e.Code != "" {
		return fmt.Sprintf("sarvamai API error (%d) [%s]: %s", e.StatusCode, e.Code, msg)
	}
	return fmt.Sprintf("sarvamai API error (%d): %s", e.StatusCode, msg)
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

	// Try to parse the body as JSON to extract message and code.
	// Many APIs use different formats, so we try a few common ones.
	var generic map[string]any
	if err := json.Unmarshal(bodyBytes, &generic); err == nil {
		// 1. Check for "detail" (FastAPI/Common)
		if detail, ok := generic["detail"].(string); ok {
			apiErr.Message = detail
		} else if detailObj, ok := generic["detail"].([]any); ok {
			// Handle validation list
			var details []string
			for _, d := range detailObj {
				details = append(details, fmt.Sprintf("%v", d))
			}
			apiErr.Message = strings.Join(details, "; ")
		}

		// 2. Check for "message"
		if msg, ok := generic["message"].(string); ok && apiErr.Message == "" {
			apiErr.Message = msg
		}

		// 3. Check for "error" (OpenAI style)
		if errVal, ok := generic["error"]; ok {
			if errMsg, ok := errVal.(string); ok && apiErr.Message == "" {
				apiErr.Message = errMsg
			} else if errMap, ok := errVal.(map[string]any); ok {
				if msg, ok := errMap["message"].(string); ok && apiErr.Message == "" {
					apiErr.Message = msg
				}
				if code, ok := errMap["code"].(string); ok {
					apiErr.Code = code
				}
			}
		}

		// 4. Check for top-level "code"
		if code, ok := generic["code"].(string); ok && apiErr.Code == "" {
			apiErr.Code = code
		}
	}

	// Even if JSON decoding fails,
	// we still return a structured error with RawBody populated.
	return apiErr
}
