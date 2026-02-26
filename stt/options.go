package stt

import (
	"github.com/Shreehari-Acharya/sarvam-go-sdk/languages"
)

// Option is a functional option for configuring a request.
type Option func(*transcribeRequest) error

// WithModel sets the speech recognition model.
// Options: ModelSaarika (saarika:v2.5), ModelSaaras (saaras:v3)
// Default: ModelSaarika
func WithModel(model Model) Option {
	return func(req *transcribeRequest) error {
		req.Model = &model
		return nil
	}
}

// WithMode sets the processing mode for speech recognition.
// Options: ModeTranscribe, ModeTranslate, ModeVerbatim, ModeTranslit, ModeCodemix
// Note: Mode is only supported with saaras:v3 model
func WithMode(mode Mode) Option {
	return func(req *transcribeRequest) error {
		req.Mode = &mode
		return nil
	}
}

// WithLanguage sets the language code for the audio.
// Use languages.Code values (e.g., languages.CodeEnIN)
func WithLanguage(language languages.Code) Option {
	return func(req *transcribeRequest) error {
		req.Language = &language
		return nil
	}
}

// WithAudioCodec sets the audio codec of the input file.
// This helps the API process the audio correctly.
func WithAudioCodec(codec InputAudioCodec) Option {
	return func(req *transcribeRequest) error {
		req.AudioCodec = &codec
		return nil
	}
}

// StreamOption is a functional option for configuring a streaming request.
type StreamOption func(*streamTranscribeRequest) error

// WithStreamSampleRate sets the audio sample rate for streaming.
// Options: SampleRate8000, SampleRate16000 (recommended), SampleRate22050, SampleRate24000
// Default: 16000
func WithStreamSampleRate(rate StreamSampleRate) StreamOption {
	return func(cfg *streamTranscribeRequest) error {
		cfg.SampleRate = &rate
		return nil
	}
}

// WithStreamHighVADSensitivity enables high Voice Activity Detection sensitivity.
// This helps detect quieter speech but may result in more false positives.
func WithStreamHighVADSensitivity(enabled bool) StreamOption {
	return func(cfg *streamTranscribeRequest) error {
		cfg.HighVADSensitivity = &enabled
		return nil
	}
}

// WithStreamVADSignals enables receiving voice activity detection signals.
// When enabled, you'll receive events for speech start/end detection.
func WithStreamVADSignals(enabled bool) StreamOption {
	return func(cfg *streamTranscribeRequest) error {
		cfg.VADSignals = &enabled
		return nil
	}
}

// WithStreamFlushSignal enables flush signals.
// When enabled, the server will send signals when it finishes processing a segment.
func WithStreamFlushSignal(enabled bool) StreamOption {
	return func(cfg *streamTranscribeRequest) error {
		cfg.FlushSignal = &enabled
		return nil
	}
}

// WithStreamInputAudioCodec sets the audio codec for streaming input.
// Options: wav, pcm_s16le, pcm_l16, pcm_raw
func WithStreamInputAudioCodec(codec InputAudioCodec) StreamOption {
	return func(cfg *streamTranscribeRequest) error {
		cfg.InputAudioCodec = &codec
		return nil
	}
}

// WithStreamModel sets the speech recognition model for streaming.
// Options: ModelSaarika (saarika:v2.5), ModelSaaras (saaras:v3)
func WithStreamModel(model Model) StreamOption {
	return func(cfg *streamTranscribeRequest) error {
		cfg.Model = &model
		return nil
	}
}

// WithStreamMode sets the processing mode for streaming speech recognition.
// Options: ModeTranscribe, ModeTranslate, ModeVerbatim, ModeTranslit, ModeCodemix
// Note: Mode is only supported with saaras:v3 model
func WithStreamMode(mode Mode) StreamOption {
	return func(cfg *streamTranscribeRequest) error {
		cfg.Mode = &mode
		return nil
	}
}

// WithStreamLanguage sets the language code for streaming recognition.
func WithStreamLanguage(language languages.Code) StreamOption {
	return func(cfg *streamTranscribeRequest) error {
		cfg.Language = language
		return nil
	}
}
