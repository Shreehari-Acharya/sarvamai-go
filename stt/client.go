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

// Client provides speech-to-text transcription services.
type Client struct {
	transport *transport.Transport
}

// NewClient creates a new STT client.
func NewClient(t *transport.Transport) *Client {
	return &Client{
		transport: t,
	}
}

// Transcribe performs speech-to-text transcription on the provided audio file.
//
// # Parameters
//
//	ctx: Context for the request
//	req: Transcription request containing audio file and options
//
// # Returns
//
//	TranscribeResponse containing the transcribed text, or an error
//
// # Example
//
//	audio, err := os.Open("recording.wav")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer audio.Close()
//
//	resp, err := client.Transcribe(ctx, stt.TranscribeRequest{
//	    File:     audio,
//	    FileName: "recording.wav",
//	    Model:    stt.ModelPtr(stt.ModelSaarika),
//	})
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
	err := req.validate()
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

// TranscribeStream starts a streaming speech-to-text transcription session.
//
// # Parameters
//
//	ctx: Context for the WebSocket connection
//	cfg: Stream configuration
//
// # Returns
//
//	Stream for sending audio and receiving transcriptions, or an error
//
// # Streaming Workflow
//
//  1. Call TranscribeStream to get a Stream instance
//  2. Send audio using stream.SendAudio() in a loop
//  3. Call stream.Flush() to finalize current segment and get results
//  4. Read transcriptions from stream.Messages()
//  5. Handle errors from stream.Errors()
//  6. Call stream.Close() when done
//
// # Example
//
//	stream, err := client.TranscribeStream(ctx, stt.StreamConfig{
//	    Language:   languages.CodePtr(languages.Code("en-IN")),
//	    Model:     stt.ModelPtr(stt.ModelSaarika),
//	    SampleRate: stt.SampleRate16000,
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer stream.Close()
//
//	// Send audio in a goroutine
//	go func() {
//	    for chunk := range audioChunks {
//	        if err := stream.SendAudio(chunk); err != nil {
//	            return
//	        }
//	    }
//	    stream.Flush()
//	}()
//
//	// Receive transcriptions
//	for resp := range stream.Messages() {
//	    if resp.Type == stt.TypeData {
//	        var data stt.TranscriptionData
//	        resp.UnmarshalData(&data)
//	        fmt.Println(data.Transcript)
//	    }
//	}
func (c *Client) TranscribeStream(ctx context.Context, cfg StreamConfig) (*Stream, error) {

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

	// validate config parameters before dialing WebSocket
	if err := cfg.validate(); err != nil {
		return nil, err
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
		sampleRate: cfg.SampleRate,
	}

	go stream.readLoop()

	return stream, nil
}

// SendAudio sends audio data to the streaming transcription.
// Audio should be PCM format at the sample rate specified in StreamConfig.
func (s *Stream) SendAudio(pcm []byte) error {
	payload := map[string]any{
		"audio": map[string]any{
			"data":        base64.StdEncoding.EncodeToString(pcm),
			"sample_rate": strconv.Itoa(int(s.sampleRate)),
			"encoding":    "audio/wav",
		},
	}

	return s.ws.WriteJSON(payload)
}

// Flush signals the end of the current audio segment.
// Call this after sending audio to get the final transcription for that segment.
func (s *Stream) Flush() error {
	return s.ws.WriteJSON(map[string]string{
		"type": "flush",
	})
}

// Messages returns a channel of streaming responses.
func (s *Stream) Messages() <-chan StreamResponse {
	return s.messages
}

// Errors returns a channel for errors during streaming.
func (s *Stream) Errors() <-chan error {
	return s.errs
}

// Close closes the streaming session.
func (s *Stream) Close() error {
	close(s.done)
	return s.ws.Close()
}

// Transcript returns the accumulated transcript from all received messages.
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
