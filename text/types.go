package text

import (
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
