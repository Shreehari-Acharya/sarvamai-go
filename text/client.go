package text

import (
	"context"

	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/transport"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/languages"
)

type TextClient struct {
	transport *transport.Transport
}

func NewTextClient(t *transport.Transport) *TextClient {
	return &TextClient{
		transport: t,
	}
}

type translateRequest struct {
	Input              string          `json:"input"`
	SourceLanguageCode languages.Code  `json:"source_language_code"`
	TargetLanguageCode languages.Code  `json:"target_language_code"`
	SpeakerGender      *SpeakerGender  `json:"speaker_gender,omitempty"`
	Mode               *TranslateMode  `json:"mode,omitempty"`
	Model              *TranslateModel `json:"model,omitempty"`
	OutputScript       *OutputScript   `json:"output_script,omitempty"`
	NumeralsFormat     *NumeralsFormat `json:"numerals_format,omitempty"`
}

// Translate converts text from one language to another while preserving its meaning.
//
// For example, 'मैं ऑफिस जा रहा हूँ' translates to 'I am going to the office' in English,
// where the script and language change, but the original meaning remains the same.
//
// # Parameters
//
//	ctx: Context for the request
//	input: The text to translate (max 1000 chars for mayura:v1, 2000 chars for sarvam-translate:v1)
//	sourceLang: Source language code (use "auto" for mayura:v1 to auto-detect)
//	targetLang: Target language code
//	options: Optional functional options (WithModel, WithMode, WithOutputScript, etc.)
//
// # Model-specific notes
//
// mayura:v1:
//   - Supports 12 languages with auto-detection
//   - Supports modes: formal, modern-colloquial, classic-colloquial, code-mixed
//   - Supports output scripts: null, roman, fully-native, spoken-form-in-native
//
// sarvam-translate:v1:
//   - Supports 22 Indian languages
//   - Only supports formal mode
//   - Does not support output_script
//
// # Example
//
//	resp, err := client.Text.Translate(ctx, "Hello", "en-IN", "hi-IN")
//	// resp.TranslatedText = "नमस्ते"
//
// # API Reference
//
// https://docs.sarvam.ai/api-reference-docs/text/translate-text
func (c *TextClient) Translate(
	ctx context.Context,
	input string,
	sourceLang languages.Code,
	targetLang languages.Code,
	options ...translateOption,
) (*TranslateResponse, error) {
	var resp TranslateResponse

	req := &translateRequest{
		Input:              input,
		SourceLanguageCode: sourceLang,
		TargetLanguageCode: targetLang,
	}

	for _, opt := range options {
		if err := opt(req); err != nil {
			return nil, err
		}
	}

	// Validate the request before sending to catch any client-side errors early.
	if err := validateTranslateRequest(*req); err != nil {
		return nil, err
	}

	err := c.transport.DoRequest(
		ctx,
		"POST",
		"/translate",
		req,
		&resp,
		"application/json",
	)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

type transliterateRequest struct {
	Input                      string                      `json:"input"`
	SourceLanguageCode         languages.Code              `json:"source_language_code"`
	TargetLanguageCode         languages.Code              `json:"target_language_code"`
	NumeralsFormat             *NumeralsFormat             `json:"numerals_format,omitempty"`
	SpokenFormNumeralsLanguage *SpokenFormNumeralsLanguage `json:"spoken_form_numerals_language,omitempty"`
	SpokenForm                 *bool                       `json:"spoken_form,omitempty"`
}

// Transliterate converts text from one script to another while preserving the original pronunciation.
//
// For example, 'नमस्ते' becomes 'namaste' in English, and 'how are you' can be written as 'हाउ आर यू'
// in Devanagari. This process ensures that the sound of the original text remains intact,
// even when written in a different script.
//
// # Parameters
//
//	ctx: Context for the request
//	input: The text to transliterate (max 1000 chars)
//	sourceLang: Source language code (use "auto" for auto-detection)
//	targetLang: Target language code
//	options: Optional functional options (WithSpokenForm, WithNumeralsFormatTransliteration, etc.)
//
// # Supported Languages
//
// # English, Hindi, Bengali, Gujarati, Kannada, Malayalam, Marathi, Odia, Punjabi, Tamil, Telugu
//
// # Example
//
//	resp, err := client.Text.Transliterate(ctx, "Hello", "en-IN", "hi-IN")
//	// resp.TransliteratedText = "हैलो"
//
// # API Reference
//
// https://docs.sarvam.ai/api-reference-docs/text/transliterate-text
func (c *TextClient) Transliterate(ctx context.Context,
	input string,
	sourceLang languages.Code,
	targetLang languages.Code,
	options ...transliterationOption,
) (*TransliterateResponse, error) {
	var resp TransliterateResponse

	req := &transliterateRequest{
		Input:              input,
		SourceLanguageCode: sourceLang,
		TargetLanguageCode: targetLang,
	}

	for _, opt := range options {
		if err := opt(req); err != nil {
			return nil, err
		}
	}
	// Validate the request before sending to catch any client-side errors early.
	if err := validateTransliterateRequest(*req); err != nil {
		return nil, err
	}

	err := c.transport.DoRequest(
		ctx,
		"POST",
		"/transliterate",
		req,
		&resp,
		"application/json",
	)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

type detectLanguageRequest struct {
	Input string `json:"input"`
}

// DetectLanguage identifies the language and script of the given text.
//
// It returns the detected language code (e.g., "hi-IN", "en-IN") and script code
// (e.g., "Deva" for Devanagari, "Latn" for Latin).
//
// # Parameters
//
//	ctx: Context for the request
//	input: The text to analyze (max 1000 characters)
//
// # Example
//
//	resp, err := client.Text.DetectLanguage(ctx, "नमस्ते")
//	// resp.LanguageCode = "hi-IN"
//	// resp.ScriptCode = "Deva"
//
// # API Reference
//
// https://docs.sarvam.ai/api-reference-docs/text/identify-language
func (c *TextClient) DetectLanguage(ctx context.Context,
	input string,
) (*DetectLanguageResponse, error) {
	var resp DetectLanguageResponse

	req := detectLanguageRequest{
		Input: input,
	}

	// Validate the request before sending to catch any client-side errors early.
	if err := validateDetectLanguageRequest(req); err != nil {
		return nil, err
	}
	err := c.transport.DoRequest(
		ctx,
		"POST",
		"/text-lid",
		req,
		&resp,
		"application/json",
	)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
