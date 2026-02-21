package translate

import (
	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/sarvamaierrors"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/languages"
)

var (
	mayuraLanguages = map[languages.Code]bool{
		"auto":  true,
		"bn-IN": true,
		"en-IN": true,
		"gu-IN": true,
		"hi-IN": true,
		"kn-IN": true,
		"ml-IN": true,
		"mr-IN": true,
		"od-IN": true,
		"pa-IN": true,
		"ta-IN": true,
		"te-IN": true,
	}

	sarvamTranslateLanguages = map[languages.Code]bool{
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

type SpeakerGender string

const (
	GenderMale   SpeakerGender = "Male"
	GenderFemale SpeakerGender = "Female"
)

type TranslateMode string

const (
	ModeFormal            TranslateMode = "formal"
	ModeModernColloquial  TranslateMode = "modern-colloquial"
	ModeClassicColloquial TranslateMode = "classic-colloquial"
	ModeCodeMixed         TranslateMode = "code-mixed"
)

type TranslateModel string

const (
	ModelMayura          TranslateModel = "mayura:v1"
	ModelSarvamTranslate TranslateModel = "sarvam-translate:v1"
)

type OutputScript string

const (
	OutputScriptNull               OutputScript = "null"
	OutputScriptRoman              OutputScript = "roman"
	OutputScriptFullyNative        OutputScript = "fully-native"
	OutputScriptSpokenFormInNative OutputScript = "spoken-form-in-native"
)

type NumeralsFormat string

const (
	NumeralsInternational NumeralsFormat = "international"
	NumeralsNative        NumeralsFormat = "native"
)

type Request struct {
	Input              string         `json:"input"`
	SourceLanguageCode languages.Code `json:"source_language_code"`
	TargetLanguageCode languages.Code `json:"target_language_code"`

	SpeakerGender  *SpeakerGender  `json:"speaker_gender,omitempty"`
	Mode           *TranslateMode  `json:"mode,omitempty"`
	Model          *TranslateModel `json:"model,omitempty"`
	OutputScript   *OutputScript   `json:"output_script,omitempty"`
	NumeralsFormat *NumeralsFormat `json:"numerals_format,omitempty"`
}

type Response struct {
	RequestID          *string        `json:"request_id"`
	TranslatedText     string         `json:"translated_text"`
	SourceLanguageCode languages.Code `json:"source_language_code"`
}

func (r Request) Validate() error {
	model := r.Model
	if model == nil || *model == ModelMayura {
		if err := validateMayuraSourceLanguage(r.SourceLanguageCode); err != nil {
			return err
		}
		if err := validateMayuraTargetLanguage(r.TargetLanguageCode); err != nil {
			return err
		}
	} else if *model == ModelSarvamTranslate {
		if err := validateSarvamTranslateSourceLanguage(r.SourceLanguageCode); err != nil {
			return err
		}
		if err := validateSarvamTranslateTargetLanguage(r.TargetLanguageCode); err != nil {
			return err
		}
	}

	if model != nil && *model == ModelSarvamTranslate {
		if r.Mode != nil && *r.Mode != ModeFormal {
			return &sarvamaierrors.ValidationError{
				Field:   "mode",
				Message: "sarvam-translate:v1 only supports formal mode",
			}
		}
		if r.OutputScript != nil && *r.OutputScript != OutputScriptNull {
			return &sarvamaierrors.ValidationError{
				Field:   "output_script",
				Message: "sarvam-translate:v1 does not support output_script",
			}
		}
	}

	return nil
}
