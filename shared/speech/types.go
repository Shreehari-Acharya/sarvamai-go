// Contains types used across multiple packages in the speech domain.
// stt, sttjob, translate, translatejob.
package speech

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/transport"
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
	ModelSaarika   Model = "saarika:v2.5"
	ModelSaaras    Model = "saaras:v3"
	ModelSaarasV25 Model = "saaras:v2.5"
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

// Stream is an iterator for streaming speech recognition responses.
// Use Next() to iterate through responses, Current() to get the current response,
// and Err() to check for errors.
type Stream struct {
	mu         sync.Mutex
	ws         *transport.WSConnection
	done       bool
	flushSent  bool
	err        error
	current    StreamResponse
	transcript strings.Builder
	sampleRate StreamSampleRate
	doneChan   chan struct{}
}

// StreamData contains the transcribed text and metadata from streaming.
type StreamData struct {
	RequestID           string              `json:"request_id"`
	Transcript          string              `json:"transcript"`
	Timestamps          *Timestamps         `json:"timestamps"`
	DiarizedTranscript  *DiarizedTranscript `json:"diarized_transcript"`
	LanguageCode        *string             `json:"language_code"`
	LanguageProbability *float64            `json:"language_probability"`
	Metrics             Metrics             `json:"metrics"`
}

// Metrics contains performance metrics for the transcription.
type Metrics struct {
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

// NewStream creates a new STTStream from a WebSocket connection and config.
func NewStream(ws *transport.WSConnection, sampleRate StreamSampleRate) *Stream {
	return &Stream{
		ws:         ws,
		sampleRate: sampleRate,
		doneChan:   make(chan struct{}),
	}
}

// SendAudio sends audio data to the streaming transcription.
// Audio should be PCM format at the sample rate specified in StreamConfig.
func (s *Stream) SendAudio(pcm []byte) error {
	payload := map[string]any{
		"audio": map[string]any{
			"data":        base64.StdEncoding.EncodeToString(pcm),
			"sample_rate": strconv.Itoa(int(s.sampleRate)),
			"encoding":    "audio/wav",
		},
	}

	return s.ws.WriteJSON(payload)
}

// Flush signals the end of the current audio segment.
// Call this after sending audio to get the final transcription for that segment.
func (s *Stream) Flush() error {
	if s.flushSent {
		return nil // Prevent multiple flush signals
	}
	s.flushSent = true
	return s.ws.WriteJSON(map[string]string{
		"type": "flush",
	})
}

// Next advances the iterator to the next response.
// Returns true if there is a response available, false if the stream is done or errored.
func (s *Stream) Next() bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.done || s.err != nil {
		return false
	}

	_, data, err := s.ws.ReadMessage()
	if err != nil {
		select {
		case <-s.doneChan:
			s.done = true
		default:
			s.err = err
		}
		return false
	}

	var resp StreamResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		s.err = err
		return false
	}

	s.current = resp

	switch resp.Type {
	case TypeError:
		var errData ErrorData
		if err := resp.UnmarshalData(&errData); err == nil {
			s.err = fmt.Errorf("stream error: %s (code: %s)", errData.Error, errData.Code)
		} else {
			s.err = fmt.Errorf("stream error with invalid error data: %w", err)
		}
		return false
	case TypeData:
		var data StreamData
		if err := resp.UnmarshalData(&data); err == nil {
			s.transcript.WriteString(data.Transcript)
		}
		if s.flushSent {
			s.done = true
		}
		return true
	case TypeEvents:
		// Handle events if needed (e.g., VAD signals)
		return true
	}

	return true
}

// Current returns the current streaming response.
// Valid only after Next returns true.
func (s *Stream) Current() StreamResponse {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.current
}

// Text returns all accumulated transcript as a string.
func (s *Stream) Text() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.transcript.String()
}

// Err returns the error encountered during streaming, if any.
func (s *Stream) Err() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.err
}

// Close closes the stream and releases resources.
func (s *Stream) Close() error {
	close(s.doneChan)
	return s.ws.Close()
}

type Callback struct {
	URL       string  `json:"url"`
	AuthToken *string `json:"auth_token,omitempty"`
}

type ContainerType string

const (
	ContainerTypeAzure   ContainerType = "Azure"
	ContainerTypeLocal   ContainerType = "Local"
	ContainerTypeGCS     ContainerType = "Google"
	ContainerTypeAzureV1 ContainerType = "Azure_V1"
)

type JobState = string

const (
	JobStateAccepted  JobState = "Accepted"
	JobStatePending   JobState = "Pending"
	JobStateRunning   JobState = "Running"
	JobStateCompleted JobState = "Completed"
	JobStateFailed    JobState = "Failed"
)

type FileState string

const (
	FileStateSuccess             FileState = "Success"
	FileStateAPIError            FileState = "API Error"
	FileStateInternalServerError FileState = "Internal Server Error"
)

type FileDetail struct {
	FileName string `json:"file_name"`
	FileId   string `json:"file_id"`
}

type JobDetail struct {
	Inputs        []FileDetail `json:"inputs,omitempty"`
	Outputs       []FileDetail `json:"outputs,omitempty"`
	State         *FileState   `json:"state,omitempty"`
	ErrorMessage  *string      `json:"error_message,omitempty"`
	ExceptionName *string      `json:"exception_name,omitempty"`
}
