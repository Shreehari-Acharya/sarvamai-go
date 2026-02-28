package stt

import (
	"slices"

	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/sarvamaierrors"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/languages"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/shared/speech"
)

// Option is a functional option for configuring a request.
type Option func(*transcribeRequest) error

// WithModel sets the speech recognition model.
// Options: ModelSaarika (saarika:v2.5), ModelSaaras (saaras:v3)
// Default: ModelSaarika
func WithModel(model speech.Model) Option {
	return func(req *transcribeRequest) error {
		if model != speech.ModelSaarika && model != speech.ModelSaaras {
			return &sarvamaierrors.ValidationError{
				Field:   "model",
				Message: "invalid model, must be saarika:v2.5 or saaras:v3",
			}
		}
		req.Model = &model
		return nil
	}
}

// WithMode sets the processing mode for speech recognition.
// Options: ModeTranscribe, ModeTranslate, ModeVerbatim, ModeTranslit, ModeCodemix
// Note: Mode is only supported with saaras:v3 model
func WithMode(mode speech.Mode) Option {
	return func(req *transcribeRequest) error {
		req.Mode = &mode
		return nil
	}
}

// WithLanguage sets the language code for the audio.
// Use languages.Code values (e.g., languages.CodeEnIN)
func WithLanguage(language languages.Code) Option {
	return func(req *transcribeRequest) error {
		if language == "" {
			return nil
		}

		// we are checking against SaarasLanguages because Saaras is the superset of languages for both Saarika and Saaras models.
		// model specific language validation will be done in the validation step where we will check if the selected model supports the selected language.
		if !languages.SaarasLanguages[language] {
			return &sarvamaierrors.ValidationError{
				Field:   "language",
				Message: "invalid language code.",
			}
		}

		req.Language = &language
		return nil
	}
}

// WithAudioCodec sets the audio codec of the input file.
// This helps the API process the audio correctly.
func WithAudioCodec(codec speech.InputAudioCodec) Option {
	return func(req *transcribeRequest) error {
		req.AudioCodec = &codec
		return nil
	}
}

// StreamOption is a functional option for configuring a streaming request.
type StreamOption func(*streamTranscribeRequest) error

// WithStreamSampleRate sets the audio sample rate for streaming.
// Options: SampleRate8000, SampleRate16000 (recommended)
// Default: 16000
func WithStreamSampleRate(rate speech.StreamSampleRate) StreamOption {
	return func(cfg *streamTranscribeRequest) error {
		if !slices.Contains(speech.AllowedSampleRatesForStream, rate) {
			return &sarvamaierrors.ValidationError{
				Field:   "sample_rate",
				Message: "invalid sample rate for streaming, must be 8000 or 16000",
			}
		}
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
// When enabled, you'll receive signals indicating when speech starts and ends in the stream.
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
func WithStreamInputAudioCodec(codec speech.InputAudioCodec) StreamOption {
	return func(cfg *streamTranscribeRequest) error {
		if !slices.Contains(speech.AllowedInputAudioCodecsForStream, codec) {
			return &sarvamaierrors.ValidationError{
				Field:   "input_audio_codec",
				Message: "unsupported audio codec for streaming, must be wav, pcm_s16le, pcm_l16, or pcm_raw",
			}
		}
		cfg.InputAudioCodec = &codec
		return nil
	}
}

// WithStreamModel sets the speech recognition model for streaming.
// Options: ModelSaarika (saarika:v2.5), ModelSaaras (saaras:v3)
func WithStreamModel(model speech.Model) StreamOption {
	return func(cfg *streamTranscribeRequest) error {
		if model != speech.ModelSaarika && model != speech.ModelSaaras {
			return &sarvamaierrors.ValidationError{
				Field:   "model",
				Message: "invalid model, must be saarika:v2.5 or saaras:v3",
			}
		}
		cfg.Model = &model
		return nil
	}
}

// WithStreamMode sets the processing mode for streaming speech recognition.
// Options: ModeTranscribe, ModeTranslate, ModeVerbatim, ModeTranslit, ModeCodemix
// Note: Mode is only supported with saaras:v3 model
func WithStreamMode(mode speech.Mode) StreamOption {
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
