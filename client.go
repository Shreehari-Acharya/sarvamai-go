package sarvamai

import (
	"errors"

	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/transport"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/stt"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/text"
)

// Client provides access to Sarvam AI services.
//
// # Services
//
// The client exposes the following services:
//
//   - Text: For translation, transliteration, and language detection
//   - SpeechToText: For speech-to-text transcription (REST and streaming)
//
// # Example
//
//	client, err := sarvamai.NewClient(sarvamai.Config{
//	    APIKey: "your-api-key",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Use text translation
//	resp, err := client.Text.Translate(ctx, translate.Request{...})
//
//	// Use speech-to-text
//	resp, err := client.SpeechToText.Transcribe(ctx, stt.TranscribeRequest{...})
type Client struct {
	transport *transport.Transport

	Text         *text.Client
	SpeechToText *stt.Client
}

// NewClient creates a new Sarvam AI client with the given configuration.
//
// # Parameters
//
//	cfg: Configuration containing API key and optional settings. APIKey is required.
//
// # Returns
//
//	A new Client instance or an error if the API key is missing.
//
// # Example
//
//	client, err := sarvamai.NewClient(sarvamai.Config{
//	    APIKey:     "your-api-key",
//	    BaseURL:    "https://api.sarvam.ai",  // optional, default: https://api.sarvam.ai
//	    HTTPClient: &http.Client{Timeout: 30 * time.Second}, // optional
//	})
func NewClient(cfg Config) (*Client, error) {
	if cfg.APIKey == "" {
		return nil, errors.New("SARVAM API key is required. Create one at https://dashboard.sarvam.ai/key-management")
	}

	def := defaultConfig()

	if cfg.BaseURL == "" {
		cfg.BaseURL = def.BaseURL
	}

	if cfg.HTTPClient == nil {
		cfg.HTTPClient = def.HTTPClient
	}

	t := &transport.Transport{
		APIKey:     cfg.APIKey,
		BaseURL:    cfg.BaseURL,
		HTTPClient: cfg.HTTPClient,
	}

	c := &Client{
		transport: t,
	}

	c.Text = text.NewClient(t)
	c.SpeechToText = stt.NewClient(t)

	return c, nil
}
