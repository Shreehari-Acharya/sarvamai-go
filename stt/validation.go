package stt

import (
	"fmt"

	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/sarvamaierrors"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/languages"
)

func validateForSaarasMode(r *TranscribeRequest) error {
	if r.Mode != nil {
		if r.Model == nil || *r.Model != ModelSaaras {
			return &sarvamaierrors.ValidationError{
				Field:   "mode",
				Message: "mode is only supported with saaras:v3 model",
			}
		}
	}
	return nil
}

func validateLanguage(r *TranscribeRequest) error {
	if r.Language == nil {
		return nil
	}

	supported, modelName := getSupportedLanguagesForModel(r.Model)
	if supported == nil {
		return nil
	}

	if !supported[*r.Language] {
		return &sarvamaierrors.ValidationError{
			Field:   "language_code",
			Message: fmt.Sprintf("%s is not supported by %s model", *r.Language, modelName),
		}
	}
	return nil
}

func getSupportedLanguagesForModel(model *Model) (map[languages.Code]bool, string) {
	if model == nil {
		return nil, ""
	}
	switch *model {
	case ModelSaarika:
		return saarikaLanguages, "saarika:v2.5"
	case ModelSaaras:
		return saarasLanguages, "saaras:v3"
	}
	return nil, ""
}

func validateFile(r *TranscribeRequest) error {
	if r.File == nil {
		return &sarvamaierrors.ValidationError{Field: "file", Message: "file is required"}
	}
	return nil
}

func validateCodec(r *TranscribeRequest) error {
	if r.AudioCodec != nil {
		if !allowedAudioCodecs[*r.AudioCodec] {
			return &sarvamaierrors.ValidationError{
				Field:   "input_audio_codec",
				Message: "unsupported audio codec",
			}
		}
	}
	return nil
}
