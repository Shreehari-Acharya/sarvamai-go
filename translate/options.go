package translate

import (
	"slices"

	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/sarvamaierrors"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/shared/speech"
)

type TranslateOption func(*translateRequest) error

// WithPrompt sets a custom prompt to guide the translation output.
// Conversation context can be passed as a prompt to boost model accuracy.
// However, the current system is at an experimentation stage and doesn't
// match the prompt performance of large language models.
func WithPrompt(prompt string) TranslateOption {
	return func(req *translateRequest) error {
		req.Prompt = &prompt
		return nil
	}
}

// WithModel sets the speech recognition model for translation.
// currently only saaras:v2.5 is supported for translation mode.
func WithModel(model speech.Model) TranslateOption {

	// Note: overrinding the model to saaras:v2.5 since translation is only supported with that model as of now.
	// In the future, if more models support translation, we can remove this override and allow users to specify the model.
	model = speech.ModelSaarasV25

	return func(req *translateRequest) error {
		req.Model = &model
		return nil
	}
}

// WithAudioCodec sets the audio codec of the input file for translation.
// Audio codec/format of the input file. Our API automatically detects all codec formats,
// but for PCM files specifically (pcm_s16le, pcm_l16, pcm_raw), you must pass this parameter.
// PCM files are supported only at 16kHz sample rate.
func WithAudioCodec(codec speech.InputAudioCodec) TranslateOption {
	return func(req *translateRequest) error {
		req.AudioCodec = &codec
		return nil
	}
}

type TranslateStreamOption func(*streamTranslateRequest) error

// WithModelForTranslateStream sets the speech recognition model for streaming translation.
//
// saaras:v3 (recommended) supports mode parameter and provides best accuracy.
// saaras:v2.5 (legacy) is kept for backward compatibility.
func WithModelForTranslateStream(model speech.Model) TranslateStreamOption {

	return func(cfg *streamTranslateRequest) error {

		// model must be saaras:v2.5 or saaras:v3.
		if model != speech.ModelSaarasV25 && model != speech.ModelSaaras {
			return &sarvamaierrors.ValidationError{
				Field:   "model",
				Message: "invalid model for translation. Supported models: saaras:v2.5, saaras:v3",
			}
		}
		cfg.Model = &model
		return nil
	}
}

// WithModeForTranslateStream sets the mode for streaming translation.
//
// Only applicable when using saaras:v3 model.
//
// Available modes:
//   - translate (default): Translates speech from any supported Indic language to English
//   - transcribe: Standard transcription in the original language
//   - verbatim: Exact word-for-word transcription without normalization
//   - translit: Romanization - Transliterates speech to Latin/Roman script only
//   - codemix: Code-mixed text with English words in English and Indic words in native script
func WithModeForTranslateStream(mode speech.Mode) TranslateStreamOption {
	return func(cfg *streamTranslateRequest) error {
		validModes := []speech.Mode{
			speech.ModeTranscribe,
			speech.ModeTranslate,
			speech.ModeVerbatim,
			speech.ModeTranslit,
			speech.ModeCodemix,
		}
		if !slices.Contains(validModes, mode) {
			return &sarvamaierrors.ValidationError{
				Field:   "mode",
				Message: "invalid mode for streaming translation. Supported modes: transcribe, translate, verbatim, translit, codemix",
			}
		}
		cfg.Mode = &mode
		return nil
	}
}

// WithAudioCodecForTranslateStream sets the audio codec of the input stream for translation.
// Supported codecs for streaming translation: wav, pcm_s16le, pcm_l16, pcm_raw
func WithAudioCodecForTranslateStream(codec speech.InputAudioCodec) TranslateStreamOption {
	return func(cfg *streamTranslateRequest) error {
		if !slices.Contains(speech.AllowedInputAudioCodecsForStream, codec) {
			return &sarvamaierrors.ValidationError{
				Field:   "input_audio_codec",
				Message: "unsupported audio codec for streaming translation. Supported codecs: wav, pcm_s16le, pcm_l16, pcm_raw",
			}
		}
		cfg.InputAudioCodec = &codec
		return nil
	}
}

// WithSampleRateForTranslateStream sets the sample rate for streaming translation.
// Audio sample rate for the WebSocket connection. When specified as a connection parameter,
// only 16kHz and 8kHz are supported. 8kHz is only available via this connection parameter.
// If not specified, defaults to 16kHz.
func WithSampleRateForTranslateStream(rate speech.StreamSampleRate) TranslateStreamOption {
	return func(cfg *streamTranslateRequest) error {
		if !slices.Contains(speech.AllowedSampleRatesForStream, rate) {
			return &sarvamaierrors.ValidationError{
				Field:   "sample_rate",
				Message: "unsupported sample rate for streaming translation. Supported rates: 8000, 16000 Hz",
			}
		}

		cfg.SampleRate = &rate
		return nil
	}
}

// WithHighVADSensitivityForTranslateStream enables or disables high VAD (Voice Activity Detection) sensitivity for streaming translation.
func WithHighVADSensitivityForTranslateStream(enabled bool) TranslateStreamOption {
	return func(cfg *streamTranslateRequest) error {
		cfg.HighVadSensitivity = &enabled
		return nil
	}
}

// WithVADSignalsForTranslateStream enables or disables  VAD (Voice Activity Detection) in response for streaming translation.
func WithVADSignalsForTranslateStream(enabled bool) TranslateStreamOption {
	return func(cfg *streamTranslateRequest) error {
		cfg.VADSignals = &enabled
		return nil
	}
}

// WithFlushSignalForTranslateStream enables or disables flush signals for streaming translation.
// Signal to flush the audio buffer and finalize transcription and translation
func WithFlushSignalForTranslateStream(enabled bool) TranslateStreamOption {
	return func(cfg *streamTranslateRequest) error {
		cfg.FlushSignal = &enabled
		return nil
	}
}
