package text

// options for translate API

type translateOption func(*translateRequest) error

// WithSpeakerGender sets the desired speaker gender.
// This influences the translation style based on the specified gender.
// Options: GenderMale = "Male", GenderFemale = "Female".
func WithSpeakerGender(gender SpeakerGender) translateOption {
	return func(req *translateRequest) error {
		req.SpeakerGender = &gender
		return nil
	}
}

// WithMode sets the desired translation mode/style.
// Options:
//   - ModeFormal: Formal translation suitable for professional contexts
//   - ModeModernColloquial: Modern colloquial translation with contemporary language
//   - ModeClassicColloquial: Classic colloquial translation with traditional phrasing
//   - ModeCodeMixed: Code-mixed translation mixing multiple languages
//
// # Usage Notes
//
//   - ModeFormal is the only mode supported by sarvam-translate:v1 model
//   - Mayura model supports all modes
//
// Default - ModeFormal
func WithMode(mode TranslateMode) translateOption {
	return func(req *translateRequest) error {
		req.Mode = &mode
		return nil
	}
}

// WithModel sets the translation model to use.
// Options:
//   - Mayura: The latest translation model with support for all modes and features
//   - sarvam-translate:v1: An older model that only supports formal translation mode
func WithModel(model TranslateModel) translateOption {
	return func(req *translateRequest) error {
		req.Model = &model
		return nil
	}
}

// WithOutputScript sets the desired script for the translated output.
// For mayura:v1 - We support transliteration with four options:
//
// Options:
// OutputScriptNull: No transliteration applied.
// OutputScriptRoman: Transliteration in Romanized script.
// OutputScriptFullyNative: Transliteration in the native script with formal style.
// OutputScriptSpokenFormInNative: Transliteration in the native script with spoken style.
// For sarvam-translate:v1 - Transliteration is not supported.
//
// Default - OutputScriptNull (no transliteration)
func WithOutputScript(script OutputScript) translateOption {
	return func(req *translateRequest) error {
		req.OutputScript = &script
		return nil
	}
}

// WithNumeralsFormat sets the desired format for numerals in the translated output.
// supported for both mayura:v1 and sarvam-translate:v1
// Options:
// NumeralsInternational: Uses regular numerals (0-9).
// NumeralsNative: Uses language-specific native numerals.
// Default - international
func WithNumeralsFormat(format NumeralsFormat) translateOption {
	return func(req *translateRequest) error {
		req.NumeralsFormat = &format
		return nil
	}
}

// options for transliteration API

type transliterationOption func(*transliterateRequest) error

// WithNumeralsFormatTransliteration sets the desired format for numerals in the transliterated output.
// Options:
// NumeralsInternational: Uses regular numerals (0-9).
// NumeralsNative: Uses language-specific native numerals.
// Default - international
func WithNumeralsFormatTransliteration(format NumeralsFormat) transliterationOption {
	return func(req *transliterateRequest) error {
		req.NumeralsFormat = &format
		return nil
	}
}

// WithSpokenFormNumeralsLanguage sets the language for spoken form numerals in the transliterated output.
// This option is relevant only when the spoken_form field is set to true.
// It specifies the language to use for converting numerals into their spoken form in the output.
// For example, if the input text contains "2024" and the target language is Hindi with spoken form enabled,
// setting this option to Hindi will convert "2024" to "दो हजार चौबीस" in the output.
// Options:
// SpokenFormNumeralsEnglish: Use English language rules for spoken form numerals.
// SpokenFormNumeralsNative: Use the target language's native rules for spoken form numerals.
// Default - SpokenFormNumeralsNative (use the target language's native rules)
func WithSpokenFormNumeralsLanguage(lang SpokenFormNumeralsLanguage) transliterationOption {
	return func(req *transliterateRequest) error {
		req.SpokenFormNumeralsLanguage = &lang
		return nil
	}
}

// WithSpokenForm sets whether to convert numerals into their spoken form in the transliterated output.
// When enabled, numerals in the input text will be converted to their spoken form in the output based on the specified language rules.
// For example, if the input text contains "2024" and the target language is Hindi with spoken form enabled,
// the output will convert "2024" to "दो हजार चौबीस" instead of keeping it as "2024".
func WithSpokenForm(spokenForm bool) transliterationOption {
	return func(req *transliterateRequest) error {
		req.SpokenForm = &spokenForm
		return nil
	}
}
