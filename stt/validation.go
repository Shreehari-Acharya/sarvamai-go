package stt

import (
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

	model := req.Model
	mode := req.Mode
	language := req.Language

	if err := speech.ValidateMode(model, mode); err != nil {
		return err
	}

	if language == nil {
		return nil
	}

	if err := speech.ValidateLanguageWithSpec(model, *language, true); err != nil {
		return err
	}

	return nil
}

func validateStreamConfig(cfg *streamTranscribeRequest) error {

	if err := speech.ValidateMode(cfg.Model, cfg.Mode); err != nil {
		return err
	}

	if err := speech.ValidateLanguageWithSpec(cfg.Model, cfg.Language, true); err != nil {
		return err
	}

	return nil
}
