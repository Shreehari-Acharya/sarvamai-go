package stt

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/transport"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/languages"
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

func (c *Client) TranscribeStream(ctx context.Context, cfg StreamConfig) (*Stream, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	query := url.Values{}

	if cfg.Language == nil {
		lang := languages.Code("unknown")
		cfg.Language = &lang
	}
	query.Set("language-code", string(*cfg.Language))

	if cfg.Model != nil {
		query.Set("model", string(*cfg.Model))
	}
	if cfg.Mode != nil {
		query.Set("mode", string(*cfg.Mode))
	}
	if cfg.SampleRate != 0 {
		query.Set("sample_rate", strconv.Itoa(int(cfg.SampleRate)))
	}
	if cfg.InputAudioCodec != nil {
		query.Set("input_audio_codec", string(*cfg.InputAudioCodec))
	}

	query.Set("high_vad_sensitivity", strconv.FormatBool(cfg.HighVADSensitivity))
	query.Set("vad_signals", strconv.FormatBool(cfg.VADSignals))
	query.Set("flush_signal", strconv.FormatBool(cfg.FlushSignal))

	wsConn, err := c.transport.DialWebSocket(ctx, "/speech-to-text/ws", query)
	if err != nil {
		return nil, err
	}

	stream := &Stream{
		ws:         wsConn,
		messages:   make(chan StreamResponse),
		errs:       make(chan error, 1), // buffered to prevent blocking on error
		done:       make(chan struct{}),
		transcript: "",
	}

	go stream.readLoop()

	return stream, nil
}

func (s *Stream) SendAudio(pcm []byte) error {
	payload := map[string]any{
		"audio": map[string]any{
			"data":        base64.StdEncoding.EncodeToString(pcm),
			"sample_rate": "16000",
			"encoding":    "audio/wav",
		},
	}

	return s.ws.WriteJSON(payload)
}

func (s *Stream) Flush() error {
	return s.ws.WriteJSON(map[string]string{
		"type": "flush",
	})
}

func (s *Stream) Messages() <-chan StreamResponse {
	return s.messages
}

func (s *Stream) Errors() <-chan error {
	return s.errs
}

func (s *Stream) Close() error {
	close(s.done)
	return s.ws.Close()
}

func (s *Stream) Transcript() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.transcript
}

func (s *Stream) readLoop() {
	defer close(s.messages)
	defer close(s.errs)

	for {
		_, data, err := s.ws.ReadMessage()
		if err != nil {
			select {
			case <-s.done:
				// Stream was closed, exit gracefully
			default:
				s.errs <- err
			}
			return
		}

		var resp StreamResponse
		if err := json.Unmarshal(data, &resp); err != nil {
			s.errs <- err
			continue
		}

		if resp.Type == TypeData {
			var transcriptionData TranscriptionData
			if err := resp.UnmarshalData(&transcriptionData); err == nil {
				s.mu.Lock()
				if s.transcript != "" {
					s.transcript += " "
				}
				s.transcript += transcriptionData.Transcript
				s.mu.Unlock()
			}
		}

		s.messages <- resp
	}
}
