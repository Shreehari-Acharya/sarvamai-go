package translate

import (
	"fmt"

	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/sarvamaierrors"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/languages"
)

//
// Model Registry
//

type modelSpec struct {
	supportedLanguages   map[languages.Code]bool
	supportsMode         bool
	supportsOutputScript bool
	name                 string
}

var modelRegistry = map[TranslateModel]modelSpec{
	ModelMayura: {
		supportedLanguages:   languages.MayuraLanguages,
		supportsMode:         false,
		supportsOutputScript: false,
		name:                 "mayura:v1",
	},
	ModelSarvamTranslate: {
		supportedLanguages:   languages.SarvamTranslateLanguages,
		supportsMode:         false,
		supportsOutputScript: false,
		name:                 "sarvam-translate:v1",
	},
}

func getModelSpec(model *TranslateModel, defaultIfNil bool) (*modelSpec, error) {
	if model == nil {
		if defaultIfNil {
			spec := modelRegistry[ModelMayura]
			return &spec, nil
		}
		return nil, nil
	}

	spec, ok := modelRegistry[*model]
	if !ok {
		return nil, fmt.Errorf("unknown model: %s", *model)
	}

	return &spec, nil
}

func validateSourceLanguage(model *TranslateModel, lang languages.Code) error {
	spec, err := getModelSpec(model, true)
	if err != nil {
		return err
	}
	if spec == nil || spec.supportedLanguages == nil {
		return nil
	}

	if !spec.supportedLanguages[lang] {
		return &sarvamaierrors.ValidationError{
			Field:   "source_language_code",
			Message: fmt.Sprintf("%s is not supported by %s model", lang, spec.name),
		}
	}
	return nil
}

func validateTargetLanguage(model *TranslateModel, lang languages.Code) error {
	// auto is not allowed as target language for translation models
	if lang == "auto" {
		spec, err := getModelSpec(model, true)
		if err != nil {
			return err
		}
		return &sarvamaierrors.ValidationError{
			Field:   "target_language_code",
			Message: fmt.Sprintf("auto is not supported as target language for %s model", spec.name),
		}
	}

	spec, err := getModelSpec(model, true)
	if err != nil {
		return err
	}
	if spec == nil || spec.supportedLanguages == nil {
		return nil
	}

	if !spec.supportedLanguages[lang] {
		return &sarvamaierrors.ValidationError{
			Field:   "target_language_code",
			Message: fmt.Sprintf("%s is not supported by %s model", lang, spec.name),
		}
	}
	return nil
}
