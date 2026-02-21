package languages

import (
	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/sarvamaierrors"
)

type Code string

func (c Code) String() string {
	return string(c)
}

func (c Code) IsValid() bool {
	_, ok := Languages[c]
	return ok
}

var (
	Languages = map[Code]bool{
		"auto":   true,
		"bn-IN":  true,
		"en-IN":  true,
		"gu-IN":  true,
		"hi-IN":  true,
		"kn-IN":  true,
		"ml-IN":  true,
		"mr-IN":  true,
		"od-IN":  true,
		"pa-IN":  true,
		"ta-IN":  true,
		"te-IN":  true,
		"as-IN":  true,
		"brx-IN": true,
		"doi-IN": true,
		"kok-IN": true,
		"ks-IN":  true,
		"mai-IN": true,
		"mni-IN": true,
		"ne-IN":  true,
		"sa-IN":  true,
		"sat-IN": true,
		"sd-IN":  true,
		"ur-IN":  true,
	}
)

func ValidateDetectLanguageInput(input string) error {
	if input == "" {
		return &sarvamaierrors.ValidationError{
			Field:   "input",
			Message: "input cannot be empty for language detection",
		}
	}
	if len(input) > 1000 {
		return &sarvamaierrors.ValidationError{
			Field:   "input",
			Message: "input exceeds maximum length of 1000 characters for language detection",
		}
	}
	return nil
}
