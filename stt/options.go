package stt

import (
	"github.com/Shreehari-Acharya/sarvamai-go/languages"
	"github.com/Shreehari-Acharya/sarvamai-go/shared/speech"
)

// Option is a functional option for configuring a request.
type Option func(*transcribeRequest)

// WithModel sets the speech recognition model.
// Options: ModelSaarika (saarika:v2.5), ModelSaaras (saaras:v3)
// Default: ModelSaarika
func WithModel(model speech.Model) Option {
	return func(req *transcribeRequest) {
		req.Model = &model
	}
}

// WithMode sets the processing mode for speech recognition.
// Options: ModeTranscribe, ModeTranslate, ModeVerbatim, ModeTranslit, ModeCodemix
// Note: Mode is only supported with saaras:v3 model
func WithMode(mode speech.Mode) Option {
	return func(req *transcribeRequest) {
		req.Mode = &mode
	}
}

// WithLanguage sets the language code for the audio.
// Use languages.Code values (e.g., languages.CodeEnIN)
func WithLanguage(language languages.Code) Option {
	return func(req *transcribeRequest) {
		req.Language = &language
	}
}

// WithAudioCodec sets the audio codec of the input file.
// This helps the API process the audio correctly.
func WithAudioCodec(codec speech.InputAudioCodec) Option {
	return func(req *transcribeRequest) {
		req.AudioCodec = &codec
	}
}

// StreamOption is a functional option for configuring a streaming request.
type StreamOption func(*streamTranscribeRequest)

// WithStreamSampleRate sets the audio sample rate for streaming.
// Options: SampleRate8000, SampleRate16000 (recommended)
// Default: 16000
func WithStreamSampleRate(rate speech.StreamSampleRate) StreamOption {
	return func(cfg *streamTranscribeRequest) {
		cfg.SampleRate = &rate
	}
}

// WithStreamHighVADSensitivity enables high Voice Activity Detection sensitivity.
// This helps detect quieter speech but may result in more false positives.
func WithStreamHighVADSensitivity(enabled bool) StreamOption {
	return func(cfg *streamTranscribeRequest) {
		cfg.HighVADSensitivity = &enabled
	}
}

// WithStreamVADSignals enables receiving voice activity detection signals.
// When enabled, you'll receive signals indicating when speech starts and ends in the stream.
func WithStreamVADSignals(enabled bool) StreamOption {
	return func(cfg *streamTranscribeRequest) {
		cfg.VADSignals = &enabled
	}
}

// WithStreamFlushSignal enables flush signals.
// When enabled, the server will send signals when it finishes processing a segment.
func WithStreamFlushSignal(enabled bool) StreamOption {
	return func(cfg *streamTranscribeRequest) {
		cfg.FlushSignal = &enabled
	}
}

// WithStreamInputAudioCodec sets the audio codec for streaming input.
// Options: wav, pcm_s16le, pcm_l16, pcm_raw
func WithStreamInputAudioCodec(codec speech.InputAudioCodec) StreamOption {
	return func(cfg *streamTranscribeRequest) {
		cfg.InputAudioCodec = &codec
	}
}

// WithStreamModel sets the speech recognition model for streaming.
// Options: ModelSaarika (saarika:v2.5), ModelSaaras (saaras:v3)
func WithStreamModel(model speech.Model) StreamOption {
	return func(cfg *streamTranscribeRequest) {
		cfg.Model = &model
	}
}

// WithStreamMode sets the processing mode for streaming speech recognition.
// Options: ModeTranscribe, ModeTranslate, ModeVerbatim, ModeTranslit, ModeCodemix
// Note: Mode is only supported with saaras:v3 model
func WithStreamMode(mode speech.Mode) StreamOption {
	return func(cfg *streamTranscribeRequest) {
		cfg.Mode = &mode
	}
}

// WithStreamLanguage sets the language code for streaming recognition.
func WithStreamLanguage(language languages.Code) StreamOption {
	return func(cfg *streamTranscribeRequest) {
		cfg.Language = language
	}
}
