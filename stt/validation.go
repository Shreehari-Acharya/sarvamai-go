package stt

import (
	"slices"

	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/sarvamaierrors"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/shared/speech"
)

func validateTranscribeRequest(req *transcribeRequest) error {
	if req.File == nil {
		return &sarvamaierrors.ValidationError{
			Field:   "file",
			Message: "file is required",
		}
	}

	if req.Model != nil {
		model := *req.Model
		if model != speech.ModelSaarika && model != speech.ModelSaaras {
			return &sarvamaierrors.ValidationError{
				Field:   "model",
				Message: "invalid model, must be saarika:v2.5 or saaras:v3",
			}
		}
	}

	if err := speech.ValidateMode(req.Model, req.Mode); err != nil {
		return err
	}

	if req.Language != nil {
		if err := speech.ValidateLanguageWithSpec(req.Model, *req.Language, true); err != nil {
			return err
		}
	}

	return nil
}

func validateStreamConfig(cfg *streamTranscribeRequest) error {
	if cfg.Model != nil {
		model := *cfg.Model
		if model != speech.ModelSaarika && model != speech.ModelSaaras {
			return &sarvamaierrors.ValidationError{
				Field:   "model",
				Message: "invalid model, must be saarika:v2.5 or saaras:v3",
			}
		}
	}

	if cfg.SampleRate != nil {
		if !slices.Contains(speech.AllowedSampleRatesForStream, *cfg.SampleRate) {
			return &sarvamaierrors.ValidationError{
				Field:   "sample_rate",
				Message: "invalid sample rate for streaming, must be 8000 or 16000",
			}
		}
	}

	if cfg.InputAudioCodec != nil {
		if !slices.Contains(speech.AllowedInputAudioCodecsForStream, *cfg.InputAudioCodec) {
			return &sarvamaierrors.ValidationError{
				Field:   "input_audio_codec",
				Message: "unsupported audio codec for streaming, must be wav, pcm_s16le, pcm_l16, or pcm_raw",
			}
		}
	}

	if err := speech.ValidateMode(cfg.Model, cfg.Mode); err != nil {
		return err
	}

	if err := speech.ValidateLanguageWithSpec(cfg.Model, cfg.Language, true); err != nil {
		return err
	}

	return nil
}
