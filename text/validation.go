package text

import (
	"slices"

	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/sarvamaierrors"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/languages"
)

func validateInputTextForTranslation(input string, model TranslateModel) error {
	if input == "" {
		return &sarvamaierrors.ValidationError{
			Field:   "input",
			Message: "input text cannot be empty",
		}
	}

	switch model {
	case ModelMayura: // max of 1000 characters for mayura
		if len(input) > 1000 {
			return &sarvamaierrors.ValidationError{
				Field:   "input",
				Message: "input text cannot exceed 1000 characters for mayura model",
			}
		}
	case ModelSarvamTranslate: // max of 2000 characters for sarvam-translate:v1
		if len(input) > 2000 {
			return &sarvamaierrors.ValidationError{
				Field:   "input",
				Message: "input text cannot exceed 2000 characters for sarvam-translate:v1 model",
			}
		}
	default:
		return &sarvamaierrors.ValidationError{
			Field:   "model",
			Message: "invalid model",
		}
	}

	return nil
}

func validateInputTextForDetectionAndTransliteration(input string) error {
	if input == "" {
		return &sarvamaierrors.ValidationError{
			Field:   "input",
			Message: "input text cannot be empty",
		}
	}

	if len(input) > 1000 {
		return &sarvamaierrors.ValidationError{
			Field:   "input",
			Message: "input text cannot exceed 1000 characters",
		}
	}

	return nil
}

func validateDetectLanguageRequest(req detectLanguageRequest) error {
	return validateInputTextForDetectionAndTransliteration(req.Input)
}

func validateTranslateRequest(req translateRequest) error {

	model := ModelSarvamTranslate
	if req.Model != nil {
		model = *req.Model
	}

	err := validateInputTextForTranslation(req.Input, model)
	if err != nil {
		return err
	}

	switch model {
	case ModelMayura:
		if languages.MayuraLanguages[req.SourceLanguageCode] != true {
			return &sarvamaierrors.ValidationError{
				Field:   "source_language_code",
				Message: "invalid source language code for mayura model",
			}
		}

		if languages.MayuraLanguages[req.TargetLanguageCode] != true {
			return &sarvamaierrors.ValidationError{
				Field:   "target_language_code",
				Message: "invalid target language code for mayura model",
			}
		}

		if req.Mode != nil && !slices.Contains(allowedModesForMayura, *req.Mode) {
			return &sarvamaierrors.ValidationError{
				Field:   "mode",
				Message: "invalid mode for mayura model",
			}
		}

		if req.OutputScript != nil && !slices.Contains(allowedOutputScriptsForMayura, *req.OutputScript) {
			return &sarvamaierrors.ValidationError{
				Field:   "output_script",
				Message: "invalid output script for mayura model",
			}
		}

	case ModelSarvamTranslate:
		if languages.SarvamTranslateLanguages[req.SourceLanguageCode] != true {
			return &sarvamaierrors.ValidationError{
				Field:   "source_language_code",
				Message: "invalid source language code for sarvam-translate model",
			}
		}

		if languages.SarvamTranslateLanguages[req.TargetLanguageCode] != true {
			return &sarvamaierrors.ValidationError{
				Field:   "target_language_code",
				Message: "invalid target language code for sarvam-translate model",
			}
		}

		if req.Mode != nil && *req.Mode != ModeFormal {
			return &sarvamaierrors.ValidationError{
				Field:   "mode",
				Message: "only 'formal' mode is supported for sarvam-translate model",
			}
		}

		if req.OutputScript != nil {
			return &sarvamaierrors.ValidationError{
				Field:   "output_script",
				Message: "transliteration is not supported for sarvam-translate model",
			}
		}

	default:
		return &sarvamaierrors.ValidationError{
			Field:   "model",
			Message: "invalid model",
		}
	}

	return nil

}

func validateTransliterateRequest(req transliterateRequest) error {

	err := validateInputTextForDetectionAndTransliteration(req.Input)
	if err != nil {
		return err
	}

	if languages.TransliterateLanguages[req.SourceLanguageCode] != true {
		return &sarvamaierrors.ValidationError{
			Field:   "source_language_code",
			Message: "invalid source language code for transliterate",
		}
	}

	if languages.TransliterateLanguages[req.TargetLanguageCode] != true {
		return &sarvamaierrors.ValidationError{
			Field:   "target_language_code",
			Message: "invalid target language code for transliterate",
		}
	}

	return nil
}
