// Package transliteration provides types for transliteration API requests and responses.
package transliteration

import (
	"github.com/Shreehari-Acharya/sarvam-go-sdk/languages"
)

// NumeralsFormat specifies the format for numerals in transliteration.
//
//   - NumeralsInternational: International numerals (1, 2, 3)
//   - NumeralsNative: Native script numerals
type NumeralsFormat string

const (
	NumeralsInternational NumeralsFormat = "international"
	NumeralsNative        NumeralsFormat = "native"
)

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

// Request represents a transliteration request.
// Transliteration converts text from one script to another (e.g., Hindi in Roman script to Devanagari).
//
// # Required Fields
//
//   - Input: The text to transliterate
//   - SourceLanguageCode: Source language code
//   - TargetLanguageCode: Target language code
//
// # Optional Fields
//
//   - NumeralsFormat: Format for numerals (international or native)
//   - SpokenFormNumeralsLanguage: Language for spoken form numerals
//   - SpokenForm: Enable spoken form output (numbers as words)
//
// # Example
//
//	req := transliteration.Request{
//	    Input:              "Namaste",
//	    SourceLanguageCode: "en-IN",
//	    TargetLanguageCode: "hi-IN",
//	}
type Request struct {
	Input                      string                      `json:"input"`
	SourceLanguageCode         languages.Code              `json:"source_language_code"`
	TargetLanguageCode         languages.Code              `json:"target_language_code"`
	NumeralsFormat             *NumeralsFormat             `json:"numerals_format,omitempty"`
	SpokenFormNumeralsLanguage *SpokenFormNumeralsLanguage `json:"spoken_form_numerals_language,omitempty"`
	SpokenForm                 *bool                       `json:"spoken_form,omitempty"`
}

// Response represents a transliteration response.
//
// # Fields
//
//   - RequestID: Unique identifier for the request
//   - TransliteratedText: The transliterated text
//   - SourceLanguageCode: Source language code
type Response struct {
	RequestID          *string        `json:"request_id"`
	TransliteratedText string         `json:"transliterated_text"`
	SourceLanguageCode languages.Code `json:"source_language_code"`
}

// Validate validates the transliteration request.
func (r Request) Validate() error {
	if err := validateTransliterateSourceLanguage(r.SourceLanguageCode); err != nil {
		return err
	}
	if err := validateTransliterateTargetLanguage(r.TargetLanguageCode); err != nil {
		return err
	}
	return nil
}
