package stt

import (
	"fmt"

	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/sarvamaierrors"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/languages"
)

//
// Model Registry
//

type modelSpec struct {
	supportedLanguages map[languages.Code]bool
	supportsMode       bool
	name               string
}

var modelRegistry = map[Model]modelSpec{
	ModelSaarika: {
		supportedLanguages: saarikaLanguages,
		supportsMode:       false,
		name:               "saarika:v2.5",
	},
	ModelSaaras: {
		supportedLanguages: saarasLanguages,
		supportsMode:       true,
		name:               "saaras:v3",
	},
}

//
// Shared Helpers
//

func getModelSpec(model *Model, defaultIfNil bool) (*modelSpec, error) {
	if model == nil {
		if defaultIfNil {
			spec := modelRegistry[ModelSaaras]
			return &spec, nil
		}
		return nil, nil
	}

	spec, ok := modelRegistry[*model]
	if !ok {
		return nil, fmt.Errorf("unknown model")
	}

	return &spec, nil
}

func validateMode(model *Model, mode *Mode) error {
	if mode == nil {
		return nil
	}

	spec, err := getModelSpec(model, false)
	if err != nil {
		return err
	}
	if spec == nil || !spec.supportsMode {
		return &sarvamaierrors.ValidationError{
			Field:   "mode",
			Message: "mode is only supported with saaras:v3 model",
		}
	}

	return nil
}

func validateLanguageWithSpec(
	model *Model,
	language *languages.Code,
	defaultIfNil bool,
) error {
	if language == nil {
		return nil
	}

	spec, err := getModelSpec(model, defaultIfNil)
	if err != nil {
		return err
	}
	if spec == nil || spec.supportedLanguages == nil {
		return nil
	}

	if !spec.supportedLanguages[*language] {
		return &sarvamaierrors.ValidationError{
			Field:   "language_code",
			Message: fmt.Sprintf("%s is not supported by %s model", *language, spec.name),
		}
	}

	return nil
}

func validateCodecValue[T comparable](
	codec *T,
	allowed map[T]bool,
	field string,
	message string,
) error {
	if codec != nil && !allowed[*codec] {
		return &sarvamaierrors.ValidationError{
			Field:   field,
			Message: message,
		}
	}
	return nil
}

//
// Non-Streaming Validations
//

func validateForSaarasMode(r *TranscribeRequest) error {
	return validateMode(r.Model, r.Mode)
}

func validateLanguage(r *TranscribeRequest) error {
	return validateLanguageWithSpec(r.Model, r.Language, false)
}

func validateFile(r *TranscribeRequest) error {
	if r.File == nil {
		return &sarvamaierrors.ValidationError{
			Field:   "file",
			Message: "file is required",
		}
	}
	return nil
}

func validateCodec(r *TranscribeRequest) error {
	return validateCodecValue(
		r.AudioCodec,
		allowedAudioCodecs,
		"input_audio_codec",
		"unsupported audio codec",
	)
}

//
// Streaming Validations
//

func validateStreamMode(s StreamConfig) error {
	return validateMode(s.Model, s.Mode)
}

func validateStreamLanguage(s StreamConfig) error {
	// Streaming defaults to Saaras if model is nil
	return validateLanguageWithSpec(s.Model, s.Language, true)
}

func validateStreamCodec(s StreamConfig) error {
	return validateCodecValue(
		s.InputAudioCodec,
		allowedStreamCodecs,
		"input_audio_codec",
		"unsupported audio codec for streaming",
	)
}

func validateStreamSampleRate(s StreamConfig) error {
	allowedSampleRates := map[StreamSampleRate]bool{
		SampleRate8000:  true,
		SampleRate16000: true,
	}

	if s.SampleRate != 0 && !allowedSampleRates[s.SampleRate] {
		return &sarvamaierrors.ValidationError{
			Field:   "sample_rate",
			Message: "invalid sample rate for streaming",
		}
	}
	return nil
}

//
// Streaming Codec Map
//

var allowedStreamCodecs = map[InputAudioCodec]bool{
	"wav":       true,
	"pcm_s16le": true,
	"pcm_l16":   true,
	"pcm_raw":   true,
}
