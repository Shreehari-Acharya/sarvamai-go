package text

type Language string

// Supported languages for translation. "auto" can be used to auto-detect the source language.
const (
	LangAuto Language = "auto"

	// languages supported my mayura:v1 model
	LangBN Language = "bn-IN" // bengali
	LangEN Language = "en-IN" // english
	LangGU Language = "gu-IN" // gujarati
	LangHI Language = "hi-IN" // hindi
	LangKN Language = "kn-IN" // kannada
	LangML Language = "ml-IN" // malayalam
	LangMR Language = "mr-IN" // marathi
	LangOD Language = "od-IN" // odia
	LangPA Language = "pa-IN" // punjabi
	LangTA Language = "ta-IN" // tamil
	LangTE Language = "te-IN" // telugu

	// additional languages supported by sarvam-translate:v1 model
	LangAS  Language = "as-IN"  // assamese
	LangBRX Language = "brx-IN" // bodo
	LangDOI Language = "doi-IN" // dogri
	LangKOK Language = "kok-IN" // konkani
	LangKS  Language = "ks-IN"  // kashmiri
	LangMAI Language = "mai-IN" // maithili
	LangMNI Language = "mni-IN" // manipuri
	LangNE  Language = "ne-IN"  // nepali
	LangSA  Language = "sa-IN"  // sanskrit
	LangSAT Language = "sat-IN" // santali
	LangSD  Language = "sd-IN"  // sindhi
	LangUR  Language = "ur-IN"  // urdu
)

// speakerGender
type SpeakerGender string

const (
	GenderMale   SpeakerGender = "Male"   // male speaker
	GenderFemale SpeakerGender = "Female" // female speaker
)

// translation modes (formal, modern-colloquial, classic-colloquial, code-mixed)
type TranslateMode string

const (
	ModeFormal            TranslateMode = "formal"             // formal translation
	ModeModernColloquial  TranslateMode = "modern-colloquial"  // modern colloquial translation (only supported for mayura:v1 model)
	ModeClassicColloquial TranslateMode = "classic-colloquial" // classic colloquial translation (only supported for mayura:v1 model)
	ModeCodeMixed         TranslateMode = "code-mixed"         // code-mixed translation (only supported for mayura:v1 model)
)

// models for translation
type TranslateModel string

const (
	ModelMayura          TranslateModel = "mayura:v1"           // Supports formal, classic-colloquial, and modern-colloquial modes
	ModelSarvamTranslate TranslateModel = "sarvam-translate:v1" // Only formal mode is supported
)

// optional parameter which controls the transliteration style applied to the output text.
// Transliteration: Converting text from one script to another while preserving pronunciation.
// supported my mayura:v1 model only
type OutputString string

const (
	OutputStringNull               OutputString = "null"                  // (default) No transliteration applied to the output text.
	OutputStringRoman              OutputString = "roman"                 // Transliteration in Romanized script.
	OutputStringFullyNative        OutputString = "fully-native"          // Transliteration in the native script with formal style.
	OutputStringSpokenFormInNative OutputString = "spoken-form-in-native" // Transliteration in the native script with spoken style.
)

// how to translate numerals in the output text
type NumeralsFormat string

const (
	NumeralsInternational NumeralsFormat = "international" // (default): Uses regular numerals (0-9).
	NumeralsNative        NumeralsFormat = "native"        // Uses language-specific native numerals.
)

// TranslateRequest represents the request payload for the translation API.
// optional fields are pointers to distinguish between "not set" and "set to zero value".
type TranslateRequest struct {
	Input              string   `json:"input"`
	SourceLanguageCode Language `json:"source_language_code"`
	TargetLanguageCode Language `json:"target_language_code"`

	SpeakerGender  *SpeakerGender  `json:"speaker_gender,omitempty"`
	Mode           *TranslateMode  `json:"mode,omitempty"`
	Model          *TranslateModel `json:"model,omitempty"`
	OutputScript   *OutputString   `json:"output_script,omitempty"`
	NumeralsFormat *NumeralsFormat `json:"numerals_format,omitempty"`
}

// TranslateResponse represents the response from the translation API.
type TranslateResponse struct {
	RequestID          *string  `json:"request_id"`
	TranslatedText     string   `json:"translated_text"`
	SourceLanguageCode Language `json:"source_language_code"`
}
