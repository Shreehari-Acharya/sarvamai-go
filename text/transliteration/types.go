package transliteration

import (
	"github.com/Shreehari-Acharya/sarvam-go-sdk/languages"
)

var (
	transliterateLanguages = map[languages.Code]bool{
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
)

type NumeralsFormat string

const (
	NumeralsInternational NumeralsFormat = "international"
	NumeralsNative        NumeralsFormat = "native"
)

type SpokenFormNumeralsLanguage string

const (
	SpokenFormNumeralsEnglish SpokenFormNumeralsLanguage = "english"
	SpokenFormNumeralsNative  SpokenFormNumeralsLanguage = "native"
)

type Request struct {
	Input                      string                      `json:"input"`
	SourceLanguageCode         languages.Code              `json:"source_language_code"`
	TargetLanguageCode         languages.Code              `json:"target_language_code"`
	NumeralsFormat             *NumeralsFormat             `json:"numerals_format,omitempty"`
	SpokenFormNumeralsLanguage *SpokenFormNumeralsLanguage `json:"spoken_form_numerals_language,omitempty"`
	SpokenForm                 *bool                       `json:"spoken_form,omitempty"`
}

type Response struct {
	RequestID          *string        `json:"request_id"`
	TransliteratedText string         `json:"transliterated_text"`
	SourceLanguageCode languages.Code `json:"source_language_code"`
}

func (r Request) Validate() error {
	if err := validateTransliterateSourceLanguage(r.SourceLanguageCode); err != nil {
		return err
	}
	if err := validateTransliterateTargetLanguage(r.TargetLanguageCode); err != nil {
		return err
	}
	return nil
}
