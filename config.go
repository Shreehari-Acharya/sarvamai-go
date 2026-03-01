// Package sarvamai provides a Go client for the Sarvam AI API.
//
// # Usage
//
// Create a client with your API key:
//
//	client, err := sarvamai.NewClient(sarvamai.Config{
//	    APIKey: "your-api-key",
//	})
//
// # API Key
//
// You can obtain an API key from the Sarvam dashboard:
// https://dashboard.sarvam.ai/key-management
package sarvamai

import (
	"net/http"
	"time"
)

// Config holds the configuration for the Sarvam AI client.
//
// # Fields
//
//   - APIKey: Your Sarvam AI API key. Required. Get one at https://dashboard.sarvam.ai/key-management
//   - BaseURL: The base URL for the API. Defaults to "https://api.sarvam.ai" if empty.
//   - HTTPClient: The HTTP client to use for requests. Defaults to a 30-second timeout if nil.
//   - MaxRetries: Maximum number of retries for transient errors (429, 5xx). Defaults to 3 if <= 0.
type Config struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
	MaxRetries int
}

func defaultConfig() Config {
	return Config{
		BaseURL: "https://api.sarvam.ai",
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		MaxRetries: 3,
	}
}
