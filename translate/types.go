// Package translate provides types for the Speech-to-Text Translate API.
//
// The Translate API automatically detects the input language, transcribes the speech,
// and translates the text to English.
//
// Two modes are supported:
//   - REST API: For quick responses under 30 seconds with immediate results
//   - Streaming: For real-time speech-to-text translation via WebSocket
package translate

import (
	"github.com/Shreehari-Acharya/sarvamai-go/languages"
	"github.com/Shreehari-Acharya/sarvamai-go/shared/speech"
)

// Aliases for common types and constants to improve developer experience.

type Model = speech.Model

const (
	ModelSaaras    = speech.ModelSaaras
	ModelSaarasV25 = speech.ModelSaarasV25
)

type Mode = speech.Mode

const (
	ModeTranscribe = speech.ModeTranscribe
	ModeTranslate  = speech.ModeTranslate
	ModeVerbatim   = speech.ModeVerbatim
	ModeTranslit   = speech.ModeTranslit
	ModeCodemix    = speech.ModeCodemix
)

type InputAudioCodec = speech.InputAudioCodec

const (
	CodecWAV      = speech.CodecWAV
	CodecMP3      = speech.CodecMP3
	CodecFLAC     = speech.CodecFLAC
	CodecPCMS16LE = speech.CodecPCMS16LE
)

type StreamSampleRate = speech.StreamSampleRate

const (
	SampleRate8000  = speech.SampleRate8000
	SampleRate16000 = speech.SampleRate16000
)

type LanguageCode = languages.Code

const (
	LanguageBnIN = languages.CodeBnIN
	LanguageEnIN = languages.CodeEnIN
	LanguageGuIN = languages.CodeGuIN
	LanguageHiIN = languages.CodeHiIN
	LanguageKnIN = languages.CodeKnIN
	LanguageMlIN = languages.CodeMlIN
	LanguageMrIN = languages.CodeMrIN
	LanguageOrIN = languages.CodeOrIN
	LanguagePaIN = languages.CodePaIN
	LanguageTaIN = languages.CodeTaIN
	LanguageTeIN = languages.CodeTeIN
)

// TranslateResponse represents the response from a speech-to-text translation request.
type TranslateResponse struct {
	// RequestID is the unique identifier for the request.
	RequestID *string `json:"request_id,omitempty"`

	// Transcript is the English translation of the provided speech.
	Transcript string `json:"transcript"`

	// LanguageCode is the BCP-47 code of detected source language.
	// Returns null when language detection is skipped (when specific language is provided).
	LanguageCode *languages.Code `json:"language_code,omitempty"`

	// DiarizedTranscript contains speaker-separated transcription.
	// Only populated when diarization is enabled (for Batch API, not REST).
	DiarizedTranscript *speech.DiarizedTranscript `json:"diarized_transcript,omitempty"`

	// LanguageProbability is the probability (0.0 to 1.0) of the detected language being correct.
	// Returns null when a specific language_code is provided in the request.
	LanguageProbability *float64 `json:"language_probability,omitempty"`
}
