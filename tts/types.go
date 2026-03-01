// Package tts provides types for Text-to-Speech API requests and responses.
package tts

import (
	"encoding/json"
	"sync"

	"github.com/Shreehari-Acharya/sarvamai-go/internal/transport"
	"github.com/Shreehari-Acharya/sarvamai-go/languages"
)

// LanguageCode is an alias for languages.Code to improve developer experience.
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

// SpeakerVoice specifies the voice to use for text-to-speech synthesis.
//
// Bulbul:v3 voices provide the latest high-quality synthesis across many languages.
// Bulbul:v2 voices are available for backward compatibility.
//
// # Bulbul:v3 Voices
//
// Indian English: shubh, aditya, ritu
// Hindi: priya, neha, rohan, roopa
// Bengali: simran, kavya
// Tamil: amit, dev
// Telugu: ishita, shreya
// Marathi: ratan, varun
// Gujarati: manan, sumit
// Kannada: kabir, aayan
// Malayalam: ashutosh, advait
// English: amelia, sophia
// Other: anand, tanya, tarun, sunny, mani, gokul, vijay, shruti, suhani, mohit, kavitha, rehan, soham, rupali
//
// # Bulbul:v2 Voices
//
// anushka, manisha, vidya, arya, abhilash, karun, hitesh
type SpeakerVoice string

const (
	// Bulbul:v3 voices
	SpeakerShubh    SpeakerVoice = "shubh"
	SpeakerAditya   SpeakerVoice = "aditya"
	SpeakerRitu     SpeakerVoice = "ritu"
	SpeakerPriya    SpeakerVoice = "priya"
	SpeakerNeha     SpeakerVoice = "neha"
	SpeakerRahul    SpeakerVoice = "rahul"
	SpeakerPooja    SpeakerVoice = "pooja"
	SpeakerRohan    SpeakerVoice = "rohan"
	SpeakerSimran   SpeakerVoice = "simran"
	SpeakerKavya    SpeakerVoice = "kavya"
	SpeakerAmit     SpeakerVoice = "amit"
	SpeakerDev      SpeakerVoice = "dev"
	SpeakerIshita   SpeakerVoice = "ishita"
	SpeakerShreya   SpeakerVoice = "shreya"
	SpeakerRatan    SpeakerVoice = "ratan"
	SpeakerVarun    SpeakerVoice = "varun"
	SpeakerManan    SpeakerVoice = "manan"
	SpeakerSumit    SpeakerVoice = "sumit"
	SpeakerRoopa    SpeakerVoice = "roopa"
	SpeakerKabir    SpeakerVoice = "kabir"
	SpeakerAayan    SpeakerVoice = "aayan"
	SpeakerAshutosh SpeakerVoice = "ashutosh"
	SpeakerAdvait   SpeakerVoice = "advait"
	SpeakerAmelia   SpeakerVoice = "amelia"
	SpeakerSophia   SpeakerVoice = "sophia"
	SpeakerAnand    SpeakerVoice = "anand"
	SpeakerTanya    SpeakerVoice = "tanya"
	SpeakerTarun    SpeakerVoice = "tarun"
	SpeakerSunny    SpeakerVoice = "sunny"
	SpeakerMani     SpeakerVoice = "mani"
	SpeakerGokul    SpeakerVoice = "gokul"
	SpeakerVijay    SpeakerVoice = "vijay"
	SpeakerShruti   SpeakerVoice = "shruti"
	SpeakerSuhani   SpeakerVoice = "suhani"
	SpeakerMohit    SpeakerVoice = "mohit"
	SpeakerKavitha  SpeakerVoice = "kavitha"
	SpeakerRehan    SpeakerVoice = "rehan"
	SpeakerSoham    SpeakerVoice = "soham"
	SpeakerRupali   SpeakerVoice = "rupali"

	// Bulbul:v2 voices
	SpeakerAnushka  SpeakerVoice = "anushka"
	SpeakerManisha  SpeakerVoice = "manisha"
	SpeakerVidya    SpeakerVoice = "vidya"
	SpeakerArya     SpeakerVoice = "arya"
	SpeakerAbhilash SpeakerVoice = "abhilash"
	SpeakerKarun    SpeakerVoice = "karun"
	SpeakerHitesh   SpeakerVoice = "hitesh"
)

// SpeechSampleRate specifies the audio sample rate for the output speech.
//
// Common sample rates:
//   - 8000:  Low quality, minimal bandwidth
//   - 16000: Standard quality (default for streaming)
//   - 22050: Good quality
//   - 24000: High quality
//   - 32000: Very high quality (Bulbul:v3 only)
//   - 44100: CD quality (Bulbul:v3 only)
//   - 48000: Studio quality (Bulbul:v3 only)
type SpeechSampleRate int

const (
	SampleRate8000  SpeechSampleRate = 8000
	SampleRate16000 SpeechSampleRate = 16000
	SampleRate22050 SpeechSampleRate = 22050
	SampleRate24000 SpeechSampleRate = 24000
	SampleRate32000 SpeechSampleRate = 32000
	SampleRate44100 SpeechSampleRate = 44100
	SampleRate48000 SpeechSampleRate = 48000
)

// Model specifies the TTS model to use.
//
//   - BulbulV2: Legacy model (v2) - supports pitch, pace, loudness adjustments
//   - BulbulV3: Current model (v3) - supports temperature control, higher sample rates
//   - BulbulV3Beta: Beta version of v3 for streaming
//
// # Model Differences
//
// | Feature          | Bulbul:v2 | Bulbul:v3 | Bulbul:v3-beta |
// |-----------------|-----------|-----------|----------------|
// | Max text length | 1500      | 2500      | -              |
// | Pitch control   | Yes       | No        | No             |
// | Pace range      | 0.3-3.0   | 0.5-2.0   | 0.5-2.0        |
// | Loudness        | Yes       | No        | No             |
// | Temperature     | No        | Yes       | Yes            |
// | Sample rates    | <=24000   | <=48000   | <=24000        |
// | Streaming       | Yes       | No        | Yes            |
type Model string

const (
	BulbulV2     Model = "bulbul:v2"
	BulbulV3     Model = "bulbul:v3"
	BulbulV3Beta Model = "bulbul:v3-beta"
)

// AudioCodec specifies the audio encoding format for the output.
//
//   - MP3:      Most compatible, good compression
//   - Linear16: PCM 16-bit, uncompressed
//   - Mulaw:    G.711 μ-law, narrowband
//   - Alaw:     G.711 A-law, narrowband
//   - Opus:     High quality, low latency
//   - FLAC:     Lossless compression
//   - AAC:      High quality, good compression
//   - WAV:      Uncompressed WAV format
type AudioCodec string

const (
	AudioCodecMP3      AudioCodec = "mp3"
	AudioCodecLinear16 AudioCodec = "linear16"
	AudioCodecMulaw    AudioCodec = "mulaw"
	AudioCodecAlaw     AudioCodec = "alaw"
	AudioCodecOpus     AudioCodec = "opus"
	AudioCodecFlac     AudioCodec = "flac"
	AudioCodecAAC      AudioCodec = "aac"
	AudioCodecWAV      AudioCodec = "wav"
)

// Bitrate specifies the audio bitrate for compressed formats.
type Bitrate string

const (
	Bitrate32K  Bitrate = "32k"
	Bitrate64K  Bitrate = "64k"
	Bitrate96K  Bitrate = "96k"
	Bitrate128K Bitrate = "128k"
	Bitrate192K Bitrate = "192k"
)

// ConvertResponse represents a text-to-speech conversion response.
type ConvertResponse struct {
	RequestId string   `json:"request_id,omitempty"`
	Audios    []string `json:"audios"`
}

// AudioData contains synthesized audio data from the streaming TTS API.
type AudioData struct {
	Audio       string `json:"audio"` // Base64 encoded audio
	ContentType string `json:"content_type"`
	RequestID   string `json:"request_id"`
}

// EventData contains event notifications from the streaming TTS API.
type EventData struct {
	EventType string `json:"event_type"` // e.g., "final" for completion
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

// MessageType for WebSocket communication
type MessageType string

const (
	TypeConfig MessageType = "config"
	TypeText   MessageType = "text"
	TypeFlush  MessageType = "flush"
	TypePing   MessageType = "ping"
	TypeAudio  MessageType = "audio"
	TypeEvent  MessageType = "event"
	TypeError  MessageType = "error"
)

// WSMessage is the envelope for WebSocket communication
type WSMessage struct {
	Type MessageType `json:"type"`
	Data any         `json:"data,omitempty"`
}

// WSResponse is the envelope for WebSocket responses
type WSResponse struct {
	Type MessageType     `json:"type"`
	Data json.RawMessage `json:"data"`
}

// TTSStream represents a WebSocket connection for streaming text-to-speech synthesis.
type TTSStream struct {
	ws     *transport.WSConnection
	audio  chan AudioData
	errs   chan error
	events chan EventData
	done   chan struct{}
}

// AudioStream is an iterator for streaming audio responses.
type AudioStream struct {
	mu       sync.Mutex
	ws       *transport.WSConnection
	audio    chan AudioData
	errs     chan error
	events   chan EventData
	done     chan struct{}
	current  AudioData
	doneFlag bool
	err      error
}
