package text

import (
	"context"

	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/transport"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/text/detect"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/text/translate"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/text/transliteration"
)

type (
	TranslateRequest        = translate.Request
	TranslateResponse       = translate.Response
	TranslateMode           = translate.TranslateMode
	TranslateModel          = translate.TranslateModel
	SpeakerGender           = translate.SpeakerGender
	OutputScript            = translate.OutputScript
	TranslateNumeralsFormat = translate.NumeralsFormat

	TransliterateRequest        = transliteration.Request
	TransliterateResponse       = transliteration.Response
	TransliterateNumeralsFormat = transliteration.NumeralsFormat
	SpokenFormNumeralsLanguage  = transliteration.SpokenFormNumeralsLanguage

	DetectRequest  = detect.Request
	DetectResponse = detect.Response
)

const (
	ModeFormal            = translate.ModeFormal
	ModeModernColloquial  = translate.ModeModernColloquial
	ModeClassicColloquial = translate.ModeClassicColloquial
	ModeCodeMixed         = translate.ModeCodeMixed

	ModelMayura          = translate.ModelMayura
	ModelSarvamTranslate = translate.ModelSarvamTranslate

	GenderMale   = translate.GenderMale
	GenderFemale = translate.GenderFemale

	OutputScriptNull               = translate.OutputScriptNull
	OutputScriptRoman              = translate.OutputScriptRoman
	OutputScriptFullyNative        = translate.OutputScriptFullyNative
	OutputScriptSpokenFormInNative = translate.OutputScriptSpokenFormInNative

	NumeralsInternational = translate.NumeralsInternational
	NumeralsNative        = translate.NumeralsNative

	TransliterateNumeralsInternational = transliteration.NumeralsInternational
	TransliterateNumeralsNative        = transliteration.NumeralsNative

	SpokenFormNumeralsEnglish = transliteration.SpokenFormNumeralsEnglish
	SpokenFormNumeralsNative  = transliteration.SpokenFormNumeralsNative
)

type Client struct {
	transport *transport.Transport
}

func NewClient(t *transport.Transport) *Client {
	return &Client{
		transport: t,
	}
}

// Translate converts text from one language to another while preserving its meaning.
//
// For example, 'मैं ऑफिस जा रहा हूँ' translates to 'I am going to the office' in English,
// where the script and language change, but the original meaning remains the same.
//
// The request can be validated before sending by calling req.Validate().
// If validation fails, the error will contain details about what fields are invalid.
//
// Example:
//
//	resp, err := client.Translate(ctx, translate.Request{
//	    Input:              "मैं ऑफिस जा रहा हूँ",
//	    SourceLanguageCode: "hi-IN",
//	    TargetLanguageCode: "en-IN",
//	})
func (c *Client) Translate(ctx context.Context, req translate.Request) (*translate.Response, error) {
	var resp translate.Response

	// Validate the request before sending to catch any client-side errors early.
	if err := req.Validate(); err != nil {
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

// Transliterate converts text from one script to another while preserving the original pronunciation.
//
// For example, 'नमस्ते' becomes 'namaste' in English, and 'how are you' can be written as 'हाउ आर यू'
// in Devanagari. This process ensures that the sound of the original text remains intact,
// even when written in a different script.
//
// The request can be validated before sending by calling req.Validate().
// If validation fails, the error will contain details about what fields are invalid.
//
// Example:
//
//	resp, err := client.Transliterate(ctx, transliteration.Request{
//	    Input:              "Hello",
//	    SourceLanguageCode: "en-IN",
//	    TargetLanguageCode: "hi-IN",
//	})
func (c *Client) Transliterate(ctx context.Context, req transliteration.Request) (*transliteration.Response, error) {
	var resp transliteration.Response

	// Validate the request before sending to catch any client-side errors early.
	if err := req.Validate(); err != nil {
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

// DetectLanguage identifies the language and script of the given text.
//
// It returns the detected language code (e.g., "hi-IN", "en-IN") and script code
// (e.g., "Deva" for Devanagari, "Latn" for Latin).
//
// Example:
//
//	resp, err := client.DetectLanguage(ctx, detect.Request{
//	    Input: "नमस्ते",
//	})
//	// resp.LanguageCode = "hi-IN"
//	// resp.ScriptCode = "Deva"
func (c *Client) DetectLanguage(ctx context.Context, req detect.Request) (*detect.Response, error) {
	var resp detect.Response

	// Validate the request before sending to catch any client-side errors early.
	if err := req.Validate(); err != nil {
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
