package stt

import (
	"context"

	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/transport"
)

type Client struct {
	transport *transport.Transport
}

func NewClient(t *transport.Transport) *Client {
	return &Client{
		transport: t,
	}
}

// Transcribe performs speech-to-text transcription on the provided audio file using the specified model and parameters.
// It validates the request parameters, constructs a multipart/form-data request, and sends it to the /speech-to-text endpoint.

func (c *Client) Transcribe(ctx context.Context, req TranscribeRequest) (*TranscribeResponse, error) {
	var resp TranscribeResponse

	fields := map[string]string{}

	if req.Model != nil {
		fields["model"] = string(*req.Model)
	}
	if req.Mode != nil {
		fields["mode"] = string(*req.Mode)
	}
	if req.Language != nil {
		fields["language_code"] = string(*req.Language)
	}
	if req.AudioCodec != nil {
		fields["input_audio_codec"] = string(*req.AudioCodec)
	}

	// Validate request parameters before making API call
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	err = c.transport.DoMultipartRequest(
		ctx,
		"/speech-to-text",
		"file",
		req.FileName,
		req.File,
		fields,
		&resp,
	)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
