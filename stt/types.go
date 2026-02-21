package stt

import (
	"github.com/Shreehari-Acharya/sarvam-go-sdk/languages"
	"io"
)

var (
	saarikaLanguages = map[languages.Code]bool{
		"hi-IN": true, // Hindi
		"bn-IN": true, // Bengali
		"kn-IN": true, // Kannada
		"ml-IN": true, // Malayalam
		"mr-IN": true, // Marathi
		"od-IN": true, // Odia
		"pa-IN": true, // Punjabi
		"ta-IN": true, // Tamil
		"te-IN": true, // Telugu
		"en-IN": true, // English
		"gu-IN": true, // Gujarati
	}

	saarasLanguages = map[languages.Code]bool{
		"hi-IN": true, // Hindi
		"bn-IN": true, // Bengali
		"kn-IN": true, // Kannada
		"ml-IN": true, // Malayalam
		"mr-IN": true, // Marathi
		"od-IN": true, // Odia
		"pa-IN": true, // Punjabi
		"ta-IN": true, // Tamil
		"te-IN": true, // Telugu
		"en-IN": true, // English
		"gu-IN": true, // Gujarati

		"as-IN":  true, // Assamese
		"ur-IN":  true, // Urdu
		"ne-IN":  true, // Nepali
		"kok-IN": true, // Konkani
		"ks-IN":  true, // Kashmiri
		"sd-IN":  true, // Sindhi
		"sa-IN":  true, // Sanskrit
		"sat-IN": true, // Santali
		"mni-IN": true, // Manipuri
		"brx-IN": true, // Bodo
		"mai-IN": true, // Maithili
		"doi-IN": true, // Dogri
	}
)

type Model string

const (
	ModelSaarika Model = "saarika:v2.5"
	ModelSaaras  Model = "saaras:v3"
)

type Mode string

const (
	ModeTranscribe Mode = "transcribe"
	ModeTranslate  Mode = "translate"
	ModeVerbatim   Mode = "verbatim"
	ModeTranslit   Mode = "translit"
	ModeCodemix    Mode = "codemix"
)

type InputAudioCodec string

var (
	allowedAudioCodecs = map[InputAudioCodec]bool{
		"wav":       true,
		"x-wav":     true,
		"wave":      true,
		"mp3":       true,
		"mpeg":      true,
		"mpeg3":     true,
		"x-mpeg-3":  true,
		"x-mp3":     true,
		"x-aac":     true,
		"aac":       true,
		"aiff":      true,
		"x-aiff":    true,
		"ogg":       true,
		"opus":      true,
		"flac":      true,
		"x-flac":    true,
		"mp4":       true,
		"x-m4a":     true,
		"amr":       true,
		"x-ms-wma":  true,
		"webm":      true,
		"pcm_s16le": true,
		"pcm_l16":   true,
		"pcm_raw":   true,
	}
)

const (
	CodecWAV      InputAudioCodec = "wav"
	CodecXWAV     InputAudioCodec = "x-wav"
	CodecWAVE     InputAudioCodec = "wave"
	CodecMP3      InputAudioCodec = "mp3"
	CodecMPEG     InputAudioCodec = "mpeg"
	CodecMPEG3    InputAudioCodec = "mpeg3"
	CodecXMPEG3   InputAudioCodec = "x-mpeg-3"
	CodecXMP3     InputAudioCodec = "x-mp3"
	CodecXAAC     InputAudioCodec = "x-aac"
	CodecAAC      InputAudioCodec = "aac"
	CodecAIFF     InputAudioCodec = "aiff"
	CodecXAIF     InputAudioCodec = "x-aiff"
	CodecOGG      InputAudioCodec = "ogg"
	CodecOPUS     InputAudioCodec = "opus"
	CodecFLAC     InputAudioCodec = "flac"
	CodecXFLAC    InputAudioCodec = "x-flac"
	CodecMP4      InputAudioCodec = "mp4"
	CodecXM4A     InputAudioCodec = "x-m4a"
	CodecAMR      InputAudioCodec = "amr"
	CodecXMSWMA   InputAudioCodec = "x-ms-wma"
	CodecWEBM     InputAudioCodec = "webm"
	CodecPCMS16LE InputAudioCodec = "pcm_s16le"
	CodecPCML16   InputAudioCodec = "pcm_l16"
	CodecPCMRAW   InputAudioCodec = "pcm_raw"
)

type TranscribeRequest struct {
	File     io.Reader
	FileName string

	Model      *Model
	Mode       *Mode
	Language   *languages.Code
	AudioCodec *InputAudioCodec
}

type Timestamps struct {
	Words            []string  `json:"words"`
	StartTimeSeconds []float64 `json:"start_time_seconds"`
	EndTimeSeconds   []float64 `json:"end_time_seconds"`
}

// DiarizedTranscript represents speaker-separated transcription.
type DiarizedTranscript struct {
	Entries []DiarizedEntry `json:"entries"`
}

// DiarizedEntry represents one speaker segment.
type DiarizedEntry struct {
	Transcript       string  `json:"transcript"`
	StartTimeSeconds float64 `json:"start_time_seconds"`
	EndTimeSeconds   float64 `json:"end_time_seconds"`
	SpeakerID        string  `json:"speaker_id"`
}

type TranscribeResponse struct {
	RequestID           *string             `json:"request_id"`
	Transcript          string              `json:"transcript"`
	Timestamps          *Timestamps         `json:"timestamps"`
	DiarizedTranscript  *DiarizedTranscript `json:"diarized_transcript"`
	LanguageCode        *string             `json:"language_code"`
	LanguageProbability *float64            `json:"language_probability"`
}

func (r *TranscribeRequest) Validate() error {

	if err := validateFile(r); err != nil {
		return err
	}

	if err := validateCodec(r); err != nil {
		return err
	}

	if err := validateForSaarasMode(r); err != nil {
		return err
	}

	if err := validateLanguage(r); err != nil {
		return err
	}

	return nil
}
