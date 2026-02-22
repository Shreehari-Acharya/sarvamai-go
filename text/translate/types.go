package translate

import (
	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/sarvamaierrors"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/languages"
)

// SpeakerGender specifies the gender of the speaker for translation.
// This is used for appropriate formality and language style.
//
//   - GenderMale: Male speaker
//   - GenderFemale: Female speaker
type SpeakerGender string

const (
	GenderMale   SpeakerGender = "Male"
	GenderFemale SpeakerGender = "Female"
)

// TranslateMode specifies the mode/style of translation.
//
//   - ModeFormal: Formal translation suitable for professional contexts
//   - ModeModernColloquial: Modern colloquial translation with contemporary language
//   - ModeClassicColloquial: Classic colloquial translation with traditional phrasing
//   - ModeCodeMixed: Code-mixed translation mixing multiple languages
//
// # Usage Notes
//
//   - ModeFormal is the only mode supported by sarvam-translate:v1 model
//   - Mayura model supports all modes
type TranslateMode string

const (
	ModeFormal            TranslateMode = "formal"
	ModeModernColloquial  TranslateMode = "modern-colloquial"
	ModeClassicColloquial TranslateMode = "classic-colloquial"
	ModeCodeMixed         TranslateMode = "code-mixed"
)

// TranslateModel specifies the translation model to use.
//
//   - ModelMayura: Mayura translation model - supports multiple modes and output scripts
//   - ModelSarvamTranslate: Sarvam Translate model - simpler model with formal mode only
//
// # Model Differences
//
// ModelMayura (mayura:v1):
//   - Supports all translation modes (formal, colloquial, code-mixed)
//   - Supports output_script options (roman, fully-native, spoken-form-in-native)
//   - Supports speaker_gender for appropriate translation style
//
// ModelSarvamTranslate (sarvam-translate:v1):
//   - Only supports formal mode
//   - Does not support output_script
//   - Simpler, faster model for basic translation needs
type TranslateModel string

const (
	ModelMayura          TranslateModel = "mayura:v1"
	ModelSarvamTranslate TranslateModel = "sarvam-translate:v1"
)

// OutputScript specifies the script for the translated output.
//
//   - OutputScriptNull: No script conversion (use default)
//   - OutputScriptRoman: Roman script (Latin alphabet)
//   - OutputScriptFullyNative: Fully native script for the target language
//   - OutputScriptSpokenFormInNative: Spoken form in native script
//
// # Usage Notes
//
//   - Only supported by ModelMayura
//   - Not supported by ModelSarvamTranslate
type OutputScript string

const (
	OutputScriptNull               OutputScript = "null"
	OutputScriptRoman              OutputScript = "roman"
	OutputScriptFullyNative        OutputScript = "fully-native"
	OutputScriptSpokenFormInNative OutputScript = "spoken-form-in-native"
)

// NumeralsFormat specifies the format for numerals in the translation.
//
//   - NumeralsInternational: International numerals (1, 2, 3)
//   - NumeralsNative: Native script numerals (१, २, ३)
type NumeralsFormat string

const (
	NumeralsInternational NumeralsFormat = "international"
	NumeralsNative        NumeralsFormat = "native"
)

// Request represents a translation request.
//
// # Fields
//
//   - Input: The text to translate (required)
//   - SourceLanguageCode: Source language code (required). Use "auto" for auto-detection.
//   - TargetLanguageCode: Target language code (required)
//   - SpeakerGender: Gender of the speaker for appropriate translation style (optional, Mayura only)
//   - Mode: Translation mode/style (optional, Mayura only)
//   - Model: Translation model to use (optional, defaults to mayura:v1)
//   - OutputScript: Output script format (optional, Mayura only)
//   - NumeralsFormat: Format for numerals in output (optional)
//
// # Example
//
//	req := translate.Request{
//	    Input:              "Hello, how are you?",
//	    SourceLanguageCode: languages.Code("en-IN"),
//	    TargetLanguageCode: languages.Code("hi-IN"),
//	    Mode:               translate.Ptr(translate.ModeFormal),
//	}
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

// Response represents a translation response.
//
// # Fields
//
//   - RequestID: Unique identifier for the request
//   - TranslatedText: The translated text
//   - SourceLanguageCode: Detected or specified source language
type Response struct {
	RequestID          *string        `json:"request_id"`
	TranslatedText     string         `json:"translated_text"`
	SourceLanguageCode languages.Code `json:"source_language_code"`
}

func (r Request) Validate() error {
	model := r.Model

	// Validate source and target languages based on model
	if err := validateSourceLanguage(model, r.SourceLanguageCode); err != nil {
		return err
	}
	if err := validateTargetLanguage(model, r.TargetLanguageCode); err != nil {
		return err
	}

	// Validate mode and output_script for sarvam-translate model
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
