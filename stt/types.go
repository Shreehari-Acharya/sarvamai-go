// Package stt provides types for Speech-to-Text API requests and responses.
package stt

import (
	"encoding/json"
	"io"
	"sync"

	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/transport"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/languages"
)

// Model specifies the speech recognition model to use.
//
//   - ModelSaarika: A multilingual speech recognition model (v2.5)
//   - ModelSaaras: A language-specific model (v3)
//
// # Model Differences
//
//   - ModelSaarika: Best for multilingual audio, supports multiple modes
//   - ModelSaaras: Optimized for specific languages, supports specific modes per language
type Model string

const (
	ModelSaarika Model = "saarika:v2.5"
	ModelSaaras  Model = "saaras:v3"
)

// Mode specifies the processing mode for speech recognition.
//
//   - ModeTranscribe: Standard transcription
//   - ModeTranslate: Translate the transcribed text to English
//   - ModeVerbatim: Preserve exact words including filler words
//   - ModeTranslit: Transliterate to Roman script
//   - ModeCodemix: Handle code-mixed audio (multiple languages)
type Mode string

const (
	ModeTranscribe Mode = "transcribe"
	ModeTranslate  Mode = "translate"
	ModeVerbatim   Mode = "verbatim"
	ModeTranslit   Mode = "translit"
	ModeCodemix    Mode = "codemix"
)

// InputAudioCodec specifies the audio codec of the input file.
// Supported codecs include: wav, mp3, aac, ogg, opus, flac, mp4, amr, webm, pcm formats.
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

// Common audio codec constants for use in requests.
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

// StreamSampleRate specifies the sample rate for streaming audio.
// Common sample rates for speech: 8000, 16000, 22050, 24000 Hz.
type StreamSampleRate int

const (
	SampleRate8000  StreamSampleRate = 8000
	SampleRate16000 StreamSampleRate = 16000
	SampleRate22050 StreamSampleRate = 22050
	SampleRate24000 StreamSampleRate = 24000
)

// StreamAudioEncoding specifies the audio encoding format for streaming.
type StreamAudioEncoding string

const (
	EncodingWAV StreamAudioEncoding = "audio/wav"
)

// VADSensitivity specifies the Voice Activity Detection sensitivity.
// Higher sensitivity detects quieter speech but may have more false positives.
//
//   - VADSensitivityHigh: Most sensitive, detects quiet speech
//   - VADSensitivityMedium: Balanced sensitivity
//   - VADSensitivityLow: Least sensitive, only clear speech
type VADSensitivity string

const (
	VADSensitivityHigh   VADSensitivity = "high"
	VADSensitivityMedium VADSensitivity = "medium"
	VADSensitivityLow    VADSensitivity = "low"
)

// TranscribeRequest represents a speech-to-text transcription request.
//
// # Required Fields
//
//   - File: Audio file reader
//   - FileName: Name of the audio file (e.g., "audio.mp3")
//
// # Optional Fields
//
//   - Model: Speech recognition model (defaults to saarika:v2.5)
//   - Mode: Processing mode (transcribe, translate, verbatim, translit, codemix)
//   - Language: Language code for the audio
//   - AudioCodec: Audio codec of the input file
//
// # Example
//
//	req := stt.TranscribeRequest{
//	    File:     audioFile,
//	    FileName: "recording.wav",
//	    Model:    stt.ModelPtr(stt.ModelSaarika),
//	    Mode:    stt.ModePtr(stt.ModeTranscribe),
//	}
type TranscribeRequest struct {
	File     io.Reader
	FileName string

	Model      *Model
	Mode       *Mode
	Language   *languages.Code
	AudioCodec *InputAudioCodec
}

// Timestamps contains word-level timing information for the transcription.
type Timestamps struct {
	Words            []string  `json:"words"`
	StartTimeSeconds []float64 `json:"start_time_seconds"`
	EndTimeSeconds   []float64 `json:"end_time_seconds"`
}

// DiarizedTranscript represents speaker-separated transcription.
// Each entry contains the transcript for a specific speaker.
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

// TranscribeResponse represents a speech-to-text transcription response.
//
// # Fields
//
//   - RequestID: Unique identifier for the request
//   - Transcript: The transcribed text
//   - Timestamps: Word-level timing information (if requested)
//   - DiarizedTranscript: Speaker-separated transcription (if speaker diarization enabled)
//   - LanguageCode: Detected or specified language
//   - LanguageProbability: Confidence of language detection
type TranscribeResponse struct {
	RequestID           *string             `json:"request_id"`
	Transcript          string              `json:"transcript"`
	Timestamps          *Timestamps         `json:"timestamps"`
	DiarizedTranscript  *DiarizedTranscript `json:"diarized_transcript"`
	LanguageCode        *string             `json:"language_code"`
	LanguageProbability *float64            `json:"language_probability"`
}

// StreamConfig holds configuration for streaming speech recognition.
//
// # Fields
//
//   - Language: Language code for recognition (use "auto" for auto-detection)
//   - Model: Speech recognition model (saarika or saaras)
//   - Mode: Processing mode
//   - SampleRate: Audio sample rate (8000, 16000, 22050, or 24000)
//   - HighVADSensitivity: Enable high VAD sensitivity
//   - VADSignals: Receive voice activity detection signals
//   - FlushSignal: Enable flush signals
//   - InputAudioCodec: Audio codec of the input
//
// # Sample Rate Notes
//
// For streaming, use 16000 Hz for best compatibility with the API.
// Other supported rates: 8000, 22050, 24000 Hz.
type StreamConfig struct {
	Language           *languages.Code
	Model              *Model
	Mode               *Mode
	SampleRate         StreamSampleRate
	HighVADSensitivity bool
	VADSignals         bool
	FlushSignal        bool
	InputAudioCodec    *InputAudioCodec
}

// ResponseType indicates the type of streaming response.
type ResponseType string

const (
	TypeData   ResponseType = "data"
	TypeError  ResponseType = "error"
	TypeEvents ResponseType = "events"
)

// StreamResponse represents a response from the streaming API.
type StreamResponse struct {
	Type ResponseType    `json:"type"`
	Data json.RawMessage `json:"data"`
}

// UnmarshalData deserializes the data field into the given destination.
func (r *StreamResponse) UnmarshalData(dest any) error {
	if r.Data == nil {
		return nil
	}
	return json.Unmarshal(r.Data, dest)
}

// TranscriptionData contains the transcribed text and metadata from streaming.
type TranscriptionData struct {
	RequestID           string               `json:"request_id"`
	Transcript          string               `json:"transcript"`
	Timestamps          *Timestamps          `json:"timestamps"`
	DiarizedTranscript  *DiarizedTranscript  `json:"diarized_transcript"`
	LanguageCode        *string              `json:"language_code"`
	LanguageProbability *float64             `json:"language_probability"`
	Metrics             TranscriptionMetrics `json:"metrics"`
}

// TranscriptionMetrics contains performance metrics for the transcription.
type TranscriptionMetrics struct {
	AudioDuration     float64 `json:"audio_duration"`
	ProcessingLatency float64 `json:"processing_latency"`
}

// ErrorData contains error information from the streaming API.
type ErrorData struct {
	Error string `json:"error"`
	Code  string `json:"code"`
}

// EventType indicates the type of voice activity detection event.
type EventType string

const (
	EventStartSpeech EventType = "START_SPEECH"
	EventEndSpeech   EventType = "END_SPEECH"
)

// EventData contains voice activity detection event information.
type EventData struct {
	EventType  EventType `json:"event_type"`
	Timestamp  string    `json:"timestamp"`
	SignalType string    `json:"signal_type"`
	OccuredAt  float64   `json:"occured_at"`
}

// Stream represents a WebSocket connection for streaming speech recognition.
// Use SendAudio to send audio data, Flush to finalize the current segment,
// and Messages/Errors to receive responses.
type Stream struct {
	ws         *transport.WSConnection
	messages   chan StreamResponse
	errs       chan error
	done       chan struct{}
	mu         sync.Mutex
	transcript string
	sampleRate StreamSampleRate
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

func (s StreamConfig) Validate() error {
	if err := validateStreamCodec(s); err != nil {
		return err
	}

	if err := validateStreamMode(s); err != nil {
		return err
	}

	if err := validateStreamLanguage(s); err != nil {
		return err
	}

	if err := validateStreamSampleRate(s); err != nil {
		return err
	}

	return nil
}
