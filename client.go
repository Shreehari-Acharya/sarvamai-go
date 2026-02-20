package sarvamai

import (
	"errors"

	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/transport"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/text"
)

type Client struct {
	transport *transport.Transport

	Text *text.Service
}

func NewSarvamAIClient(cfg Config) (*Client, error) {
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

	c.Text = text.NewService(t)

	return c, nil
}
