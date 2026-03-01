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
//   - or-IN: Odia (India)
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

const (
	CodeUnknown Code = "unknown"
	CodeAuto    Code = "auto"
	CodeBnIN    Code = "bn-IN"
	CodeEnIN    Code = "en-IN"
	CodeGuIN    Code = "gu-IN"
	CodeHiIN    Code = "hi-IN"
	CodeKnIN    Code = "kn-IN"
	CodeMlIN    Code = "ml-IN"
	CodeMrIN    Code = "mr-IN"
	CodeOrIN    Code = "or-IN"
	CodePaIN    Code = "pa-IN"
	CodeTaIN    Code = "ta-IN"
	CodeTeIN    Code = "te-IN"
	CodeAsIN    Code = "as-IN"
	CodeBrxIN   Code = "brx-IN"
	CodeDoiIN   Code = "doi-IN"
	CodeKokIN   Code = "kok-IN"
	CodeKsIN    Code = "ks-IN"
	CodeMaiIN   Code = "mai-IN"
	CodeMniIN   Code = "mni-IN"
	CodeNeIN    Code = "ne-IN"
	CodeSaIN    Code = "sa-IN"
	CodeSatIN   Code = "sat-IN"
	CodeSdIN    Code = "sd-IN"
	CodeUrIN    Code = "ur-IN"
)

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
	CodeUnknown: true,
	CodeAuto:    true,
	CodeBnIN:    true,
	CodeEnIN:    true,
	CodeGuIN:    true,
	CodeHiIN:    true,
	CodeKnIN:    true,
	CodeMlIN:    true,
	CodeMrIN:    true,
	CodeOrIN:    true,
	CodePaIN:    true,
	CodeTaIN:    true,
	CodeTeIN:    true,
	CodeAsIN:    true,
	CodeBrxIN:   true,
	CodeDoiIN:   true,
	CodeKokIN:   true,
	CodeKsIN:    true,
	CodeMaiIN:   true,
	CodeMniIN:   true,
	CodeNeIN:    true,
	CodeSaIN:    true,
	CodeSatIN:   true,
	CodeSdIN:    true,
	CodeUrIN:    true,
}
