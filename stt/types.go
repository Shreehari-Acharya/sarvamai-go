package stt

import (
	"github.com/Shreehari-Acharya/sarvam-go-sdk/shared/speech"
)

// TranscribeResponse represents the response from a speech-to-text transcription.
//
// This struct contains the transcribed text along with optional metadata
// such as timestamps, diarized transcript, and language detection results.
type TranscribeResponse struct {
	// RequestID is the unique identifier for this transcription request.
	// This can be used for debugging or tracing purposes.
	RequestID *string `json:"request_id,omitempty"`

	// Transcript is the transcribed text from the audio file.
	// This is the primary output of the transcription service.
	Transcript string `json:"transcript"`

	// Timestamps contains word-level timing information for the transcription.
	// This field is only populated when WithTimeStamps(true) option is used.
	// It includes the start and end times for each word in seconds.
	Timestamps *speech.Timestamps `json:"timestamps,omitempty"`

	// DiarizedTranscript contains speaker-separated transcription.
	// This field is only populated when WithDiarization(true) option is used
	// in batch processing (not available in REST API).
	DiarizedTranscript *speech.DiarizedTranscript `json:"diarized_transcript,omitempty"`

	// LanguageCode is the BCP-47 code of the detected language in the audio.
	// This is returned when language auto-detection is enabled (default).
	// If a specific language is provided in the request, this will match that language.
	LanguageCode *string `json:"language_code,omitempty"`

	// LanguageProbability is a float value (0.0 to 1.0) indicating the
	// confidence of the detected language.
	// Higher values indicate higher confidence.
	// This field is only populated when language is auto-detected.
	LanguageProbability *float64 `json:"language_probability,omitempty"`
}
