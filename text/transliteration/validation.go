package transliteration

import (
	"fmt"

	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/sarvamaierrors"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/languages"
)

func validateTransliterateSourceLanguage(lang languages.Code) error {
	if !transliterateLanguages[lang] {
		return &sarvamaierrors.ValidationError{
			Field:   "source_language_code",
			Message: fmt.Sprintf("%s is not supported for transliteration", lang),
		}
	}
	return nil
}

func validateTransliterateTargetLanguage(lang languages.Code) error {
	if !transliterateLanguages[lang] {
		return &sarvamaierrors.ValidationError{
			Field:   "target_language_code",
			Message: fmt.Sprintf("%s is not supported for transliteration", lang),
		}
	}
	return nil
}
