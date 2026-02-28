package tts

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/transport"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/languages"
)

// Client provides text-to-speech synthesis services.
type TTSClient struct {
	transport *transport.Transport
}

// NewTTSClient creates a new TTS client.
func NewTTSClient(t *transport.Transport) *TTSClient {
	return &TTSClient{
		transport: t,
	}
}

type ttsRequest struct {
	Text                string            `json:"text"`
	TargetLanguageCode  string            `json:"target_language_code"`
	Model               *Model            `json:"model,omitempty"`
	SpeakerVoice        *SpeakerVoice     `json:"speaker,omitempty"`
	Pitch               *float64          `json:"pitch,omitempty"`
	Pace                *float64          `json:"pace,omitempty"`
	Loudness            *float64          `json:"loudness,omitempty"`
	SpeechSampleRate    *SpeechSampleRate `json:"speech_sample_rate,omitempty"`
	EnablePreprocessing *bool             `json:"enable_preprocessing,omitempty"`
	OutputAudioCodec    *AudioCodec       `json:"output_audio_codec,omitempty"`
	Temperature         *float64          `json:"temperature,omitempty"`
}

// Convert performs text-to-speech conversion on the provided text.
//
// # Parameters
//
//	ctx: Context for the request
//	text: The text to synthesize
//	targetLanguage: Language code for synthesis (e.g., "hi-IN", "en-IN")
//	opts: Optional functional options
//
// # Returns
//
//	ConvertResponse containing the base64 encoded audio data, or an error
//
// # Example
//
//	resp, err := client.TextToSpeech.Convert(
//	    ctx,
//	    "Hello, yeh ek sarvam ai text to speech conversion ka example hai.",
//	    "hi-IN",
//	    tts.WithSpeakerVoice(tts.SpeakerShubh),
//	    tts.WithOutputAudioCodec(tts.AudioCodecMP3),
//	)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	for i, audio := range resp.Audios {
//	    decodedAudio, err := base64.StdEncoding.DecodeString(audio)
//	    if err != nil {
//	        log.Printf("Failed to decode audio %d: %v", i, err)
//	        continue
//	    }
//	    os.WriteFile(fmt.Sprintf("output_%d.mp3", i), decodedAudio, 0644)
//	}
//
// # API Reference
//
// https://docs.sarvam.ai/api-reference-docs/text-to-speech/convert
func (c *TTSClient) Convert(
	ctx context.Context,
	text string,
	targetLanguage languages.Code,
	opts ...option,
) (*ConvertResponse, error) {

	req := &ttsRequest{
		Text:               text,
		TargetLanguageCode: targetLanguage.String(),
	}

	// Apply options
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, err
		}
	}

	// Validate
	if err := validateTTSRequest(req); err != nil {
		return nil, err
	}

	var resp ConvertResponse
	err := c.transport.DoRequest(
		ctx,
		"POST",
		"/text-to-speech",
		req,
		&resp,
		"application/json",
	)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

type ttsStreamRequest struct {
	Model               *Model            `json:"model,omitempty"`
	TargetLanguageCode  string            `json:"target_language_code,omitempty"`
	Speaker             SpeakerVoice      `json:"speaker,omitempty"`
	SendCompletionEvent *bool             `json:"send_completion_event,omitempty"`
	Pitch               *float64          `json:"pitch,omitempty"`
	Pace                *float64          `json:"pace,omitempty"`
	Loudness            *float64          `json:"loudness,omitempty"`
	Temperature         *float64          `json:"temperature,omitempty"`
	SpeechSampleRate    *SpeechSampleRate `json:"speech_sample_rate,omitempty"`
	EnablePreprocessing *bool             `json:"enable_preprocessing,omitempty"`
	OutputAudioCodec    *AudioCodec       `json:"output_audio_codec,omitempty"`
	OutputAudioBitrate  *Bitrate          `json:"output_audio_bitrate,omitempty"`
	MinBufferSize       *int              `json:"min_buffer_size,omitempty"`
	MaxChunkLength      *int              `json:"max_chunk_length,omitempty"`
}

// StreamConvert starts a streaming text-to-speech synthesis session.
//
// # Parameters
//
//	ctx: Context for the WebSocket connection
//	targetLanguage: Language code for synthesis
//	opts: Optional streaming options
//
// # Returns
//
//	*AudioStream for iterating over audio chunks, or an error
//
// # Example
//
//	stream, err := client.TextToSpeech.StreamConvert(
//	    ctx,
//	    "hi-IN",
//	    tts.WithStreamSpeaker(tts.SpeakerPriya),
//	    tts.WithStreamModel(tts.BulbulV2),
//	)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer stream.Close()
//
//	// Send text
//	if err := stream.SendText("Hello, how are you?"); err != nil {
//	    log.Fatal(err)
//	}
//	stream.Flush()
//
//	// Iterate through audio chunks
//	for stream.Next() {
//	    audio := stream.Current()
//	    decodedAudio, err := base64.StdEncoding.DecodeString(audio.Audio)
//	    if err != nil {
//	        log.Printf("Failed to decode audio: %v", err)
//	        continue
//	    }
//	    // Process audio data
//	    _ = decodedAudio
//	}
//
// # API Reference
//
// https://docs.sarvam.ai/api-reference-docs/text-to-speech/stream
func (c *TTSClient) StreamConvert(
	ctx context.Context,
	targetLanguage languages.Code,
	opts ...streamOption,
) (*AudioStream, error) {

	cfg := &ttsStreamRequest{
		TargetLanguageCode: string(targetLanguage),
	}

	// Apply options
	for _, opt := range opts {
		if err := opt(cfg); err != nil {
			return nil, err
		}
	}

	// Validate
	if err := validateTTSStreamRequest(cfg); err != nil {
		return nil, err
	}

	// Build WebSocket query params
	query := url.Values{}
	if cfg.Model != nil {
		query.Set("model", string(*cfg.Model))
	}

	sendCompletion := true
	if cfg.SendCompletionEvent != nil {
		sendCompletion = *cfg.SendCompletionEvent
	}
	query.Set("send_completion_event", strconv.FormatBool(sendCompletion))

	// Dial WebSocket
	wsConn, err := c.transport.DialWebSocket(ctx, "/text-to-speech/ws", query)
	if err != nil {
		return nil, err
	}

	configPayload := map[string]any{
		"type": "config",
		"data": cfg,
	}

	if err := wsConn.WriteJSON(configPayload); err != nil {
		wsConn.Close()
		return nil, err
	}

	stream := &AudioStream{
		ws:     wsConn,
		audio:  make(chan AudioData, 100),
		errs:   make(chan error, 1),
		events: make(chan EventData, 1),
		done:   make(chan struct{}),
	}

	go stream.readLoop()
	return stream, nil
}

// Next advances the iterator to the next audio chunk.
// Returns true if there is a chunk available, false if the stream is done or errored.
func (s *AudioStream) Next() bool {
	select {
	case chunk, ok := <-s.audio:
		if !ok {
			s.doneFlag = true
			return false
		}
		s.current = chunk
		return true
	case err := <-s.errs:
		s.err = err
		s.doneFlag = true
		return false
	}
}

// Current returns the current audio chunk.
// Valid only after Next returns true.
func (s *AudioStream) Current() AudioData {
	return s.current
}

// Err returns the error encountered during streaming, if any.
func (s *AudioStream) Err() error {
	return s.err
}

// Close closes the streaming session.
func (s *AudioStream) Close() error {
	close(s.done)
	return s.ws.Close()
}

// SendText sends a chunk of text to be synthesized into speech.
func (s *AudioStream) SendText(text string) error {
	return s.ws.WriteJSON(map[string]any{
		"type": "text",
		"data": map[string]string{"text": text},
	})
}

// Flush signals the end of the current text segment.
func (s *AudioStream) Flush() error {
	return s.ws.WriteJSON(map[string]string{"type": "flush"})
}

// Ping sends a ping to keep the connection alive.
func (s *AudioStream) Ping() error {
	return s.ws.WriteJSON(map[string]string{"type": "ping"})
}

// Events returns a channel for event notifications (e.g., completion).
func (s *AudioStream) Events() <-chan EventData {
	return s.events
}

func (s *AudioStream) readLoop() {
	defer close(s.audio)
	defer close(s.errs)
	defer close(s.events)

	for {
		_, data, err := s.ws.ReadMessage()
		if err != nil {
			select {
			case <-s.done:
			default:
				s.errs <- err
			}
			return
		}

		var resp WSResponse
		if err := json.Unmarshal(data, &resp); err != nil {
			continue
		}

		switch resp.Type {
		case TypeAudio:
			var chunk AudioData
			if err := json.Unmarshal(resp.Data, &chunk); err == nil {
				s.audio <- chunk
			}
		case TypeEvent:
			var event EventData
			if err := json.Unmarshal(resp.Data, &event); err == nil {
				s.events <- event
				if event.EventType == "final" {
					return
				}
			}
		case TypeError:
			s.errs <- fmt.Errorf("api error: %s", resp.Data)
		default:
			continue
		}
	}
}
