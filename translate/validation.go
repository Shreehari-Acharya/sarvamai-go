package translate

import (
	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/sarvamaierrors"
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
