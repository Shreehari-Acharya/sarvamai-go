package detect

import (
	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/sarvamaierrors"
)

func validateDetectLanguageInput(input string) error {
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