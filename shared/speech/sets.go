package speech

import "github.com/Shreehari-Acharya/sarvam-go-sdk/languages"

// Model Registry

type modelSpec struct {
	supportedLanguages map[languages.Code]bool
	supportsMode       bool
	name               string
}

var modelRegistry = map[Model]modelSpec{
	ModelSaarika: {
		supportedLanguages: languages.SaarikaLanguages,
		supportsMode:       false,
		name:               "saarika:v2.5",
	},
	ModelSaaras: {
		supportedLanguages: languages.SaarasLanguages,
		supportsMode:       true,
		name:               "saaras:v3",
	},
}

var AllowedInputAudioCodecsForStream = []InputAudioCodec{
	CodecWAV,
	CodecPCMS16LE,
	CodecPCML16,
	CodecPCMRAW,
}

var AllowedSampleRatesForStream = []StreamSampleRate{
	SampleRate8000,
	SampleRate16000,
}
