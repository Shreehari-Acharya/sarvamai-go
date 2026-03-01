package text

import (
	"github.com/Shreehari-Acharya/sarvamai-go/languages"
)

// Aliases for common types and constants to improve developer experience.

type LanguageCode = languages.Code

const (
	LanguageBnIN = languages.CodeBnIN
	LanguageEnIN = languages.CodeEnIN
	LanguageGuIN = languages.CodeGuIN
	LanguageHiIN = languages.CodeHiIN
	LanguageKnIN = languages.CodeKnIN
	LanguageMlIN = languages.CodeMlIN
	LanguageMrIN = languages.CodeMrIN
	LanguageOrIN = languages.CodeOrIN
	LanguagePaIN = languages.CodePaIN
	LanguageTaIN = languages.CodeTaIN
	LanguageTeIN = languages.CodeTeIN
	LanguageAuto = languages.CodeAuto
)

// detect types

// ScriptCode represents the script of the detected language.
//
//   - ScriptLatn: Latin (Romanized script)
//   - ScriptDeva: Devanagari (Hindi, Marathi)
//   - ScriptBeng: Bengali
//   - ScriptGujr: Gujarati
//   - ScriptKnda: Kannada
//   - ScriptMlym: Malayalam
//   - ScriptOrya: Odia
//   - ScriptGuru: Gurmukhi (Punjabi)
//   - ScriptTaml: Tamil
//   - ScriptTelu: Telugu
type ScriptCode string

const (
	ScriptLatn ScriptCode = "Latn"
	ScriptDeva ScriptCode = "Deva"
	ScriptBeng ScriptCode = "Beng"
	ScriptGujr ScriptCode = "Gujr"
	ScriptKnda ScriptCode = "Knda"
	ScriptMlym ScriptCode = "Mlym"
	ScriptOrya ScriptCode = "Orya"
	ScriptGuru ScriptCode = "Guru"
	ScriptTaml ScriptCode = "Taml"
	ScriptTelu ScriptCode = "Telu"
)

// Response represents a language detection response.
//
// # Fields
//
//   - RequestID: Unique identifier for the request
//   - LanguageCode: Detected language code (e.g., "hi-IN" for Hindi)
//   - ScriptCode: Detected script code (e.g., "Deva" for Devanagari)
type DetectLanguageResponse struct {
	RequestID    *string     `json:"request_id"`
	LanguageCode *string     `json:"language_code"`
	ScriptCode   *ScriptCode `json:"script_code"`
}

// translate types

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

// TranslateResponse represents a translation response.
//
// # Fields
//
//   - RequestID: Unique identifier for the request
//   - TranslatedText: The translated text
//   - SourceLanguageCode: Detected or specified source language
type TranslateResponse struct {
	RequestID          *string        `json:"request_id"`
	TranslatedText     string         `json:"translated_text"`
	SourceLanguageCode languages.Code `json:"source_language_code"`
}

// transliteration types

// SpokenFormNumeralsLanguage specifies the language for spoken form numerals.
// Only applicable when SpokenForm is enabled.
//
//   - SpokenFormNumeralsEnglish: Use English words for numbers
//   - SpokenFormNumeralsNative: Use native language words for numbers
type SpokenFormNumeralsLanguage string

const (
	SpokenFormNumeralsEnglish SpokenFormNumeralsLanguage = "english"
	SpokenFormNumeralsNative  SpokenFormNumeralsLanguage = "native"
)

// TransliterateResponse represents a transliteration response.
//
// # Fields
//
//   - RequestID: Unique identifier for the request
//   - TransliteratedText: The transliterated text
//   - SourceLanguageCode: Source language code
type TransliterateResponse struct {
	RequestID          *string        `json:"request_id"`
	TransliteratedText string         `json:"transliterated_text"`
	SourceLanguageCode languages.Code `json:"source_language_code"`
}

// transliteration and translate common types

// NumeralsFormat specifies the format for numerals in the translation.
//
//   - NumeralsInternational: International numerals (1, 2, 3)
//   - NumeralsNative: Native script numerals (१, २, ३)
type NumeralsFormat string

const (
	NumeralsInternational NumeralsFormat = "international"
	NumeralsNative        NumeralsFormat = "native"
)
