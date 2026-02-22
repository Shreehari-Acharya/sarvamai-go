package sarvamai

import (
	"github.com/Shreehari-Acharya/sarvam-go-sdk/stt"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/text"
)

// Re-export text package types
type (
	TranslateRequest        = text.TranslateRequest
	TranslateResponse       = text.TranslateResponse
	TranslateMode           = text.TranslateMode
	TranslateModel          = text.TranslateModel
	SpeakerGender          = text.SpeakerGender
	OutputScript           = text.OutputScript
	TranslateNumeralsFormat = text.TranslateNumeralsFormat

	TransliterateRequest        = text.TransliterateRequest
	TransliterateResponse       = text.TransliterateResponse
	TransliterateNumeralsFormat = text.TransliterateNumeralsFormat
	SpokenFormNumeralsLanguage  = text.SpokenFormNumeralsLanguage

	DetectRequest  = text.DetectRequest
	DetectResponse = text.DetectResponse
)

const (
	ModeFormal            = text.ModeFormal
	ModeModernColloquial  = text.ModeModernColloquial
	ModeClassicColloquial = text.ModeClassicColloquial
	ModeCodeMixed         = text.ModeCodeMixed

	ModelMayura          = text.ModelMayura
	ModelSarvamTranslate = text.ModelSarvamTranslate

	GenderMale   = text.GenderMale
	GenderFemale = text.GenderFemale

	OutputScriptNull               = text.OutputScriptNull
	OutputScriptRoman              = text.OutputScriptRoman
	OutputScriptFullyNative        = text.OutputScriptFullyNative
	OutputScriptSpokenFormInNative = text.OutputScriptSpokenFormInNative

	NumeralsInternational = text.NumeralsInternational
	NumeralsNative        = text.NumeralsNative

	TransliterateNumeralsInternational = text.TransliterateNumeralsInternational
	TransliterateNumeralsNative        = text.TransliterateNumeralsNative

	SpokenFormNumeralsEnglish = text.SpokenFormNumeralsEnglish
	SpokenFormNumeralsNative  = text.SpokenFormNumeralsNative
)

// Re-export STT package types
type (
	STTTranscribeRequest  = stt.TranscribeRequest
	STTTranscribeResponse = stt.TranscribeResponse
	STTStreamConfig       = stt.StreamConfig
	STTStream             = stt.Stream
	STTStreamResponse     = stt.StreamResponse
	STTTranscriptionData  = stt.TranscriptionData
	STTTimestamps         = stt.Timestamps
	STTDiarizedTranscript = stt.DiarizedTranscript
	STTDiarizedEntry      = stt.DiarizedEntry
	STTTranscriptionMetrics = stt.TranscriptionMetrics
	STTErrorData          = stt.ErrorData
	STTEventData          = stt.EventData
)

const (
	STTModelSaarika = stt.ModelSaarika
	STTModelSaaras  = stt.ModelSaaras

	STTModeTranscribe = stt.ModeTranscribe
	STTModeTranslate  = stt.ModeTranslate
	STTModeVerbatim   = stt.ModeVerbatim
	STTModeTranslit   = stt.ModeTranslit
	STTModeCodemix    = stt.ModeCodemix

	STTCodecWAV      = stt.CodecWAV
	STTCodecXWAV     = stt.CodecXWAV
	STTCodecWAVE     = stt.CodecWAVE
	STTCodecMP3      = stt.CodecMP3
	STTCodecMPEG     = stt.CodecMPEG
	STTCodecMPEG3    = stt.CodecMPEG3
	STTCodecXMPEG3   = stt.CodecXMPEG3
	STTCodecXMP3     = stt.CodecXMP3
	STTCodecXAAC     = stt.CodecXAAC
	STTCodecAAC      = stt.CodecAAC
	STTCodecAIFF     = stt.CodecAIFF
	STTCodecXAIF     = stt.CodecXAIF
	STTCodecOGG      = stt.CodecOGG
	STTCodecOPUS     = stt.CodecOPUS
	STTCodecFLAC     = stt.CodecFLAC
	STTCodecXFLAC    = stt.CodecXFLAC
	STTCodecMP4      = stt.CodecMP4
	STTCodecXM4A     = stt.CodecXM4A
	STTCodecAMR      = stt.CodecAMR
	STTCodecXMSWMA   = stt.CodecXMSWMA
	STTCodecWEBM     = stt.CodecWEBM
	STTCodecPCMS16LE = stt.CodecPCMS16LE
	STTCodecPCML16   = stt.CodecPCML16
	STTCodecPCMRAW   = stt.CodecPCMRAW

	STTSampleRate8000  = stt.SampleRate8000
	STTSampleRate16000 = stt.SampleRate16000
	STTSampleRate22050 = stt.SampleRate22050
	STTSampleRate24000 = stt.SampleRate24000

	STTEncodingWAV = stt.EncodingWAV

	STTVADSensitivityHigh   = stt.VADSensitivityHigh
	STTVADSensitivityMedium = stt.VADSensitivityMedium
	STTVADSensitivityLow    = stt.VADSensitivityLow

	STTTypeData   = stt.TypeData
	STTTypeError  = stt.TypeError
	STTTypeEvents = stt.TypeEvents

	STTEventStartSpeech = stt.EventStartSpeech
	STTEventEndSpeech   = stt.EventEndSpeech
)
