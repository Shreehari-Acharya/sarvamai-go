// Package languages provides language codes supported by the Sarvam AI API.
package languages

// Code represents a language code conforming to the BCP 47 standard.
// Use "auto" for automatic language detection.
//
// # Supported Languages
//
// The following language codes are supported:
//
//   - unknown: Unknown language
//   - auto: Auto-detect language
//   - bn-IN: Bengali (India)
//   - en-IN: English (India)
//   - gu-IN: Gujarati (India)
//   - hi-IN: Hindi (India)
//   - kn-IN: Kannada (India)
//   - ml-IN: Malayalam (India)
//   - mr-IN: Marathi (India)
//   - od-IN: Odia (India)
//   - pa-IN: Punjabi (India)
//   - ta-IN: Tamil (India)
//   - te-IN: Telugu (India)
//   - as-IN: Assamese (India)
//   - brx-IN: Bodo (India)
//   - doi-IN: Dogri (India)
//   - kok-IN: Konkani (India)
//   - ks-IN: Kashmiri (India)
//   - mai-IN: Maithili (India)
//   - mni-IN: Manipuri (India)
//   - ne-IN: Nepali (India)
//   - sa-IN: Sanskrit (India)
//   - sat-IN: Santali (India)
//   - sd-IN: Sindhi (India)
//   - ur-IN: Urdu (India)
//
// # Usage Notes
//
//   - Use "auto" when you want the API to detect the language automatically
//   - Specify a specific language code when you know the source language
//   - Not all APIs support all languages - check each API's documentation
type Code string

// String returns the string representation of the language code.
func (c Code) String() string {
	return string(c)
}

// IsValid returns true if the language code is a supported language.
func (c Code) IsValid() bool {
	_, ok := Languages[c]
	return ok
}

// Languages is a map of all supported language codes.
// The value indicates whether the language code is valid.
var Languages = map[Code]bool{
	"unknown": true, // unknown
	"auto":    true, // auto
	"bn-IN":   true, // Bengali
	"en-IN":   true, // English
	"gu-IN":   true, // Gujarati
	"hi-IN":   true, // Hindi
	"kn-IN":   true, // Kannada
	"ml-IN":   true, // Malayalam
	"mr-IN":   true, // Marathi
	"od-IN":   true, // Odia
	"pa-IN":   true, // Punjabi
	"ta-IN":   true, // Tamil
	"te-IN":   true, // Telugu
	"as-IN":   true, // Assamese
	"brx-IN":  true, // Bodo
	"doi-IN":  true, // Dogri
	"kok-IN":  true, // Konkani
	"ks-IN":   true, // Kashmiri
	"mai-IN":  true, // Maithili
	"mni-IN":  true, // Manipuri
	"ne-IN":   true, // Nepali
	"sa-IN":   true, // Sanskrit
	"sat-IN":  true, // Santali
	"sd-IN":   true, // Sindhi
	"ur-IN":   true, // Urdu
}
