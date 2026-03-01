package speech

import (
	"fmt"

	"github.com/Shreehari-Acharya/sarvamai-go/internal/sarvamaierrors"
	"github.com/Shreehari-Acharya/sarvamai-go/languages"
)

// GetModelSpec is a helper function to get the model spec for a given model, with an option to return a default spec if the model is nil.
// This is used in validation functions to check for language and mode support based on the model.
func GetModelSpec(model *Model, defaultIfNil bool) (*modelSpec, error) {
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

func ValidateMode(model *Model, mode *Mode) error {
	if mode == nil {
		return nil
	}

	spec, err := GetModelSpec(model, true)
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

func ValidateLanguageWithSpec(
	model *Model,
	language languages.Code,
	defaultIfNil bool,
) error {
	if language == "" {
		return nil
	}

	spec, err := GetModelSpec(model, defaultIfNil)
	if err != nil {
		return err
	}
	if spec == nil || spec.supportedLanguages == nil {
		return nil
	}

	if !spec.supportedLanguages[language] {
		return &sarvamaierrors.ValidationError{
			Field:   "language_code",
			Message: fmt.Sprintf("%s is not supported by %s model", language, spec.name),
		}
	}

	return nil
}

func ValidateCodecValue[T comparable](
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
