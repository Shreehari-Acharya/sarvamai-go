package sarvamai

import (
	"net/http"
	"time"
)

type Config struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
}

func defaultConfig() Config {
	return Config{
		BaseURL: "https://api.sarvam.ai",
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}
