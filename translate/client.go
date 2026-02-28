package translate

import (
	"context"
	"io"
	"net/url"
	"strconv"

	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/transport"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/shared/speech"
)

// TranslateClient provides access to the Speech-to-Text Translate API (REST and Streaming).
//
// This client supports:
//   - REST API: For quick responses under 30 seconds with immediate results
//   - Streaming: For real-time speech-to-text translation via WebSocket
type TranslateClient struct {
	transport *transport.Transport
}

// NewTranslateClient creates a new Translate client.
//
// # Parameters
//
//	t: Transport instance configured with API key and base URL
//
// # Returns
//
//	A new TranslateClient instance
func NewTranslateClient(transport *transport.Transport) *TranslateClient {
	return &TranslateClient{
		transport: transport,
	}
}

// translateRequest represents a speech-to-text translation request.
type translateRequest struct {
	File       io.Reader               `form:"-"`                           // Exclude from JSON serialization
	Prompt     *string                 `json:"prompt,omitempty"`            // Optional prompt for translation context
	Model      *speech.Model           `json:"model,omitempty"`             // Optional model field
	AudioCodec *speech.InputAudioCodec `json:"input_audio_codec,omitempty"` // Optional audio codec
}

// Translate performs speech-to-text translation on an audio file.
//
// This REST API automatically detects the input language, transcribes the speech,
// and translates the text to English. Best for audio files under 30 seconds.
//
// # Parameters
//
//	ctx: Context for the request (used for cancellation and timeouts)
//	file: Audio file reader (supports any audio format: WAV, MP3, AAC, OGG, FLAC, etc.)
//	opts: Optional functional options to customize the request
//
// # Returns
//
//	[TranslateResponse] containing the translated transcript and metadata, or an error
//
// # Functional Options
//
//	[WithPrompt] - Set a prompt to assist transcription
//	[WithModel] - Set the model (only saaras:v2.5 supported for REST)
//	[WithAudioCodec] - Set input audio codec (required for PCM files)
//
// # Example
//
//	file, err := os.Open("audio.mp3")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer file.Close()
//
//	resp, err := client.SpeechToTextTranslate.Translate(
//	    context.Background(),
//	    file,
//	    translate.WithPrompt("This is a medical transcription"),
//	)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Transcript: %s\n", resp.Transcript)
//	fmt.Printf("Language: %s\n", *resp.LanguageCode)
//
// # API Reference
//
// https://docs.sarvam.ai/api-reference-docs/speech-to-text-translate/translate
func (c *TranslateClient) Translate(
	ctx context.Context,
	file interface{ Read([]byte) (int, error) },
	opts ...TranslateOption,
) (*TranslateResponse, error) {

	var resp TranslateResponse

	req := &translateRequest{
		File: file,
	}

	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, err
		}
	}

	if err := validateTranslateRequest(req); err != nil {
		return nil, err
	}

	err := c.transport.DoMultipartRequest(
		ctx,
		"/speech-to-text-translate",
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

type streamTranslateRequest struct {
	Model              *speech.Model            `json:"model,omitempty"`
	Mode               *speech.Mode             `json:"mode,omitempty"`
	SampleRate         *speech.StreamSampleRate `json:"sample_rate,omitempty"`
	HighVadSensitivity *bool                    `json:"high_vad_sensitivity,omitempty"`
	VADSignals         *bool                    `json:"vad_signals,omitempty"`
	FlushSignal        *bool                    `json:"flush_signal,omitempty"`
	InputAudioCodec    *speech.InputAudioCodec  `json:"input_audio_codec,omitempty"`
}

// TranslateStream establishes a WebSocket connection for real-time speech-to-text translation.
//
// Use this method for streaming audio input and receiving real-time transcription
// and translation results. The WebSocket remains open until closed by the client.
//
// # Parameters
//
//	ctx: Context for the request
//	opts: Optional functional options to configure the stream
//
// # Returns
//
//	[*speech.Stream] that can be iterated to receive transcription chunks, or an error
//
// # Functional Options
//
//	[WithModelForTranslateStream] - Set the model (saaras:v3 recommended, saaras:v2.5 legacy)
//	[WithModeForTranslateStream] - Set mode (translate, transcribe, verbatim, translit, codemix)
//	[WithSampleRateForTranslateStream] - Set sample rate (8000 or 16000 Hz)
//	[WithAudioCodecForTranslateStream] - Set audio codec (wav, pcm_s16le, pcm_l16, pcm_raw)
//	[WithHighVADSensitivityForTranslateStream] - Enable high VAD sensitivity
//	[WithVADSignalsForTranslateStream] - Enable VAD signals in response
//	[WithFlushSignalForTranslateStream] - Enable flush signals
//
// # Example
//
//	stream, err := client.SpeechToTextTranslate.TranslateStream(
//	    context.Background(),
//	    translate.WithModelForTranslateStream(speech.ModelSaaras),
//	    translate.WithModeForTranslateStream(speech.ModeTranslate),
//	    translate.WithSampleRateForTranslateStream(speech.SampleRate16000),
//	)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer stream.Close()
//
//	// Send audio data
//	audioData := readAudioFile()
//	if err := stream.SendAudio(audioData); err != nil {
//	    log.Fatal(err)
//	}
//	stream.Flush()
//
//	// Read responses
//	for stream.Next() {
//	    resp := stream.Current()
//	    if resp.Type == speech.TypeData {
//	        var data speech.StreamData
//	        resp.UnmarshalData(&data)
//	        fmt.Print(data.Transcript)
//	    }
//	}
//	if err := stream.Err(); err != nil {
//	    log.Fatal(err)
//	}
//
// # API Reference
//
// https://docs.sarvam.ai/api-reference-docs/speech-to-text-translate/translate/ws
func (c *TranslateClient) TranslateStream(
	ctx context.Context,
	opts ...TranslateStreamOption,
) (*speech.Stream, error) {

	cfg := &streamTranslateRequest{}

	for _, opt := range opts {
		if err := opt(cfg); err != nil {
			return nil, err
		}
	}

	query := url.Values{}

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
	if cfg.HighVadSensitivity != nil {
		query.Set("high_vad_sensitivity", strconv.FormatBool(*cfg.HighVadSensitivity))
	}
	if cfg.VADSignals != nil {
		query.Set("vad_signals", strconv.FormatBool(*cfg.VADSignals))
	}
	if cfg.FlushSignal != nil {
		query.Set("flush_signal", strconv.FormatBool(*cfg.FlushSignal))
	}

	wsConn, err := c.transport.DialWebSocket(ctx, "/speech-to-text-translate/ws", query)
	if err != nil {
		return nil, err
	}

	sampleRate := speech.SampleRate16000
	if cfg.SampleRate != nil {
		sampleRate = *cfg.SampleRate
	}

	stream := speech.NewStream(wsConn, sampleRate)

	return stream, nil
}
