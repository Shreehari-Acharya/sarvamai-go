package translate

import (
	"slices"

	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/sarvamaierrors"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/shared/speech"
)

func validateTranslateRequest(r *translateRequest) error {
	// File is required
	if r.File == nil {
		return &sarvamaierrors.ValidationError{
			Field:   "file",
			Message: "file is required",
		}
	}

	return nil
}

func validateTranslateStreamRequest(cfg *streamTranslateRequest) error {
	if cfg.Model != nil {
		model := *cfg.Model
		if model != speech.ModelSaarasV25 && model != speech.ModelSaaras {
			return &sarvamaierrors.ValidationError{
				Field:   "model",
				Message: "invalid model for translation. Supported models: saaras:v2.5, saaras:v3",
			}
		}
	}

	if cfg.Mode != nil {
		validModes := []speech.Mode{
			speech.ModeTranscribe,
			speech.ModeTranslate,
			speech.ModeVerbatim,
			speech.ModeTranslit,
			speech.ModeCodemix,
		}
		if !slices.Contains(validModes, *cfg.Mode) {
			return &sarvamaierrors.ValidationError{
				Field:   "mode",
				Message: "invalid mode for streaming translation. Supported modes: transcribe, translate, verbatim, translit, codemix",
			}
		}
	}

	if cfg.InputAudioCodec != nil {
		if !slices.Contains(speech.AllowedInputAudioCodecsForStream, *cfg.InputAudioCodec) {
			return &sarvamaierrors.ValidationError{
				Field:   "input_audio_codec",
				Message: "unsupported audio codec for streaming translation. Supported codecs: wav, pcm_s16le, pcm_l16, pcm_raw",
			}
		}
	}

	if cfg.SampleRate != nil {
		if !slices.Contains(speech.AllowedSampleRatesForStream, *cfg.SampleRate) {
			return &sarvamaierrors.ValidationError{
				Field:   "sample_rate",
				Message: "unsupported sample rate for streaming translation. Supported rates: 8000, 16000 Hz",
			}
		}
	}

	return nil
}
