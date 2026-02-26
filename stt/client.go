package stt

import (
	"context"
	"encoding/base64"
	"io"
	"net/url"
	"strconv"

	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/transport"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/languages"
)

// STTClient provides speech-to-text transcription services.
type STTClient struct {
	transport *transport.Transport
}

// NewSTTClient creates a new STT client.
func NewSTTClient(t *transport.Transport) *STTClient {
	return &STTClient{
		transport: t,
	}
}

// transcribeRequest represents a speech-to-text transcription request.
//
// # Required Fields
//
//   - File: Audio file reader (e.g., os.File, bytes.Buffer)
//
// # Optional Fields
//
//   - Model: Speech recognition model (defaults to saarika:v2.5)
//   - Mode: Processing mode (transcribe, translate, verbatim, translit, codemix)
//   - Language: Language code for the audio
//   - AudioCodec: Audio codec of the input file
type transcribeRequest struct {
	File       io.Reader        `form:"-"`                 // Exclude from JSON serialization
	Model      *Model           `form:"model"`             // Optional model field
	Mode       *Mode            `form:"mode"`              // Optional mode field
	Language   *languages.Code  `form:"language_code"`     // Optional language code
	AudioCodec *InputAudioCodec `form:"input_audio_codec"` // Optional audio codec
}

// Transcribe performs speech-to-text transcription on the provided audio file.
//
// # Parameters
//
//	ctx: Context for the request
//	file: Audio file reader
//	opts: Optional functional options to customize the request
//
// # Returns
//
//	TranscribeResponse containing the transcribed text, or an error
//
// # Functional Options
//
//	WithModel(Model)           - Speech recognition model (saarika:v2.5, saaras:v3)
//	WithMode(Mode)             - Processing mode (transcribe, translate, verbatim, translit, codemix)
//	WithLanguage(languages.Code) - Language code for the audio
//	WithAudioCodec(InputAudioCodec) - Audio codec of the input file
//
// # Example
//
//	audio, err := os.Open("recording.wav")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer audio.Close()
//
//	resp, err := client.Transcribe(
//	    context.Background(),
//	    audio,
//	    stt.WithModel(stt.ModelSaarika),
//	)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(resp.Transcript)
func (c *STTClient) Transcribe(
	ctx context.Context,
	file interface{ Read([]byte) (int, error) },
	opts ...Option,
) (*TranscribeResponse, error) {

	var resp TranscribeResponse

	req := &transcribeRequest{
		File: file,
	}

	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, err
		}
	}

	if err := validateTranscribeRequest(req); err != nil {
		return nil, err
	}

	err := c.transport.DoMultipartRequest(
		ctx,
		"/speech-to-text",
		"file",
		"audio.wav",
		req.File,
		req,
		&resp,
	)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// streamTranscribeRequest holds configuration for streaming speech recognition.
//
// # Fields
//
//   - Language: Language code for recognition (use "unknown" for auto-detection)
//   - Model: Speech recognition model (saarika or saaras)
//   - Mode: Processing mode
//   - SampleRate: Audio sample rate (8000, 16000, 22050, or 24000)
//   - HighVADSensitivity: Enable high VAD sensitivity
//   - VADSignals: Receive voice activity detection signals
//   - FlushSignal: Enable flush signals
//   - InputAudioCodec: Audio codec of the input
//
// # Sample Rate Notes
//
// For streaming, use 16000 Hz for best compatibility with the API.
// Other supported rates: 8000, 22050, 24000 Hz.
type streamTranscribeRequest struct {
	Language           languages.Code    `json:"language_code,omitempty"`
	Model              *Model            `json:"model,omitempty"`
	Mode               *Mode             `json:"mode,omitempty"`
	SampleRate         *StreamSampleRate `json:"sample_rate,omitempty"`
	HighVADSensitivity *bool             `json:"high_vad_sensitivity,omitempty"`
	VADSignals         *bool             `json:"vad_signals,omitempty"`
	FlushSignal        *bool             `json:"flush_signal,omitempty"`
	InputAudioCodec    *InputAudioCodec  `json:"input_audio_codec,omitempty"`
}

// TranscribeStream starts a streaming speech-to-text transcription session.
//
// # Parameters
//
//	ctx: Context for the WebSocket connection
//	opts: Optional functional options to configure the stream
//
// # Returns
//
//	*STTStream iterator for receiving transcriptions, or an error
//
// # Functional Options
//
//	WithStreamLanguage(languages.Code)    - Language code for recognition
//	WithStreamModel(Model)                - Speech recognition model
//	WithStreamMode(Mode)                  - Processing mode
//	WithStreamSampleRate(StreamSampleRate) - Audio sample rate
//	WithStreamHighVADSensitivity(bool)    - Enable high VAD sensitivity
//	WithStreamVADSignals(bool)            - Receive VAD signals
//	WithStreamFlushSignal(bool)           - Enable flush signals
//	WithStreamInputAudioCodec(InputAudioCodec) - Audio codec
//
// # Streaming Workflow
//
//  1. Call TranscribeStream to get an STTStream instance
//  2. Send audio using stream.SendAudio() in a loop
//  3. Call stream.Flush() to finalize current segment and get results
//  4. Iterate through responses using stream.Next() and stream.Current()
//  5. Check for errors using stream.Err()
//  6. Call stream.Close() when done
//
// # Example
//
//	stream, err := client.TranscribeStream(
//	    context.Background(),
//	    stt.WithStreamLanguage(languages.CodeEnIN),
//	    stt.WithStreamModel(stt.ModelSaarika),
//	    stt.WithStreamSampleRate(stt.SampleRate16000),
//	)
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
//	// Receive transcriptions using iterator
//	for stream.Next() {
//	    resp := stream.Current()
//	    if resp.Type == stt.TypeData {
//	        var data stt.TranscriptionData
//	        resp.UnmarshalData(&data)
//	        fmt.Println(data.Transcript)
//	    }
//	}
//
//	if err := stream.Err(); err != nil {
//	    log.Fatal(err)
//	}
//
//	// Or get all accumulated text at once
//	fmt.Println(stream.Text())
func (c *STTClient) TranscribeStream(
	ctx context.Context,
	language languages.Code,
	opts ...StreamOption,
) (*STTStream, error) {

	cfg := &streamTranscribeRequest{
		Language: language,
	}

	for _, opt := range opts {
		if err := opt(cfg); err != nil {
			return nil, err
		}
	}

	if err := validateStreamConfig(cfg); err != nil {
		return nil, err
	}

	query := url.Values{}

	if cfg.Language == "" {
		lang := languages.Code("unknown")
		cfg.Language = lang
	}
	query.Set("language-code", string(cfg.Language))

	if cfg.Model != nil {
		query.Set("model", string(*cfg.Model))
	}
	if cfg.Mode != nil {
		query.Set("mode", string(*cfg.Mode))
	}
	if cfg.SampleRate != nil {
		query.Set("sample_rate", strconv.Itoa(int(*cfg.SampleRate)))
	}
	if cfg.InputAudioCodec != nil {
		query.Set("input_audio_codec", string(*cfg.InputAudioCodec))
	}
	if cfg.HighVADSensitivity != nil {
		query.Set("high_vad_sensitivity", strconv.FormatBool(*cfg.HighVADSensitivity))
	}
	if cfg.VADSignals != nil {
		query.Set("vad_signals", strconv.FormatBool(*cfg.VADSignals))
	}
	if cfg.FlushSignal != nil {
		query.Set("flush_signal", strconv.FormatBool(*cfg.FlushSignal))
	}

	wsConn, err := c.transport.DialWebSocket(ctx, "/speech-to-text/ws", query)
	if err != nil {
		return nil, err
	}

	stream := NewSTTStream(wsConn, *cfg.SampleRate)

	return stream, nil
}

// SendAudio sends audio data to the streaming transcription.
// Audio should be PCM format at the sample rate specified in StreamConfig.
func (s *STTStream) SendAudio(pcm []byte) error {
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
func (s *STTStream) Flush() error {
	if s.flushSent {
		return nil // Prevent multiple flush signals
	}
	s.flushSent = true
	return s.ws.WriteJSON(map[string]string{
		"type": "flush",
	})
}
