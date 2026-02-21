package translate

import (
	"fmt"

	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/sarvamaierrors"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/languages"
)

func validateMayuraSourceLanguage(lang languages.Code) error {
	if !mayuraLanguages[lang] {
		return &sarvamaierrors.ValidationError{
			Field:   "source_language_code",
			Message: fmt.Sprintf("%s is not supported by mayura:v1 model", lang),
		}
	}
	return nil
}

func validateMayuraTargetLanguage(lang languages.Code) error {
	if lang == "auto" {
		return &sarvamaierrors.ValidationError{
			Field:   "target_language_code",
			Message: "auto is not supported as target language for mayura:v1 model",
		}
	}
	if !mayuraLanguages[lang] {
		return &sarvamaierrors.ValidationError{
			Field:   "target_language_code",
			Message: fmt.Sprintf("%s is not supported by mayura:v1 model", lang),
		}
	}
	return nil
}

func validateSarvamTranslateSourceLanguage(lang languages.Code) error {
	if !sarvamTranslateLanguages[lang] {
		return &sarvamaierrors.ValidationError{
			Field:   "source_language_code",
			Message: fmt.Sprintf("%s is not supported by sarvam-translate:v1 model", lang),
		}
	}
	return nil
}

func validateSarvamTranslateTargetLanguage(lang languages.Code) error {
	if lang == "auto" {
		return &sarvamaierrors.ValidationError{
			Field:   "target_language_code",
			Message: "auto is not supported as target language for sarvam-translate:v1 model",
		}
	}
	if !sarvamTranslateLanguages[lang] {
		return &sarvamaierrors.ValidationError{
			Field:   "target_language_code",
			Message: fmt.Sprintf("%s is not supported by sarvam-translate:v1 model", lang),
		}
	}
	return nil
}
