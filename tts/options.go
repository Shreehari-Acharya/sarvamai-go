package tts

import (
	"slices"
	"fmt"

	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/sarvamaierrors"
)

//
// Option Type and Functions
//

type option func(*ttsRequest) error

// WithModel sets the TTS model to use.
func WithModel(model Model) option {
	return func(r *ttsRequest) error {
		r.Model = &model
		return nil
	}
}

// WithSpeakerVoice sets the voice for synthesis.
func WithSpeakerVoice(voice SpeakerVoice) option {
	return func(r *ttsRequest) error {
		model := BulbulV3
		if r.Model != nil {
			model = *r.Model
		}

		var allowedVoices []SpeakerVoice
		if model == BulbulV2 {
			allowedVoices = allowedSpeakerVoicesForBulbulV2
		} else {
			allowedVoices = allowedSpeakerVoicesForBulbulV3
		}

		if slices.Contains(allowedVoices, voice) {
				r.SpeakerVoice = &voice
				return nil
			}
		return &sarvamaierrors.ValidationError{
			Field:   "speaker_voice",
			Message: fmt.Sprintf("speaker voice %s is not supported for model %s", voice, model),
		}
	}
}

// WithPitch sets the speech pitch adjustment.
// Range: -0.75 to 0.75
func WithPitch(pitch float64) option {
	return func(r *ttsRequest) error {
		if pitch < -0.75 || pitch > 0.75 {
			return &sarvamaierrors.ValidationError{
				Field:   "pitch",
				Message: "pitch must be between -0.75 and 0.75",
			}
		}
		r.Pitch = &pitch
		return nil
	}
}

// WithPace sets the speech speed.
func WithPace(pace float64) option {
	return func(r *ttsRequest) error {
		if pace < 0.3 || pace > 3.0 {
			return &sarvamaierrors.ValidationError{
				Field:   "pace",
				Message: "pace must be between 0.3 and 3.0",
			}
		}
		r.Pace = &pace
		return nil
	}
}

// WithLoudness sets the speech volume.
// Range: 0.3 to 3.0
func WithLoudness(loudness float64) option {
	return func(r *ttsRequest) error {
		if loudness < 0.3 || loudness > 3.0 {
			return &sarvamaierrors.ValidationError{
				Field:   "loudness",
				Message: "loudness must be between 0.3 and 3.0",
			}
		}
		r.Loudness = &loudness
		return nil
	}
}

// WithSpeechSampleRate sets the audio sample rate.
func WithSpeechSampleRate(rate SpeechSampleRate) option {
	return func(r *ttsRequest) error {
		r.SpeechSampleRate = &rate
		return nil
	}
}

// WithOutputAudioCodec sets the output audio format.
func WithOutputAudioCodec(codec AudioCodec) option {
	return func(r *ttsRequest) error {
		r.OutputAudioCodec = &codec
		return nil
	}
}

// WithEnablePreprocessing enables text preprocessing.
func WithEnablePreprocessing(enable bool) option {
	return func(r *ttsRequest) error {
		r.EnablePreprocessing = &enable
		return nil
	}
}

// WithTemperature controls how much randomness and expressiveness the TTS model uses.
// Range: 0.01 to 2.0
func WithTemperature(temp float64) option {
	return func(r *ttsRequest) error {
		if temp < 0.01 || temp > 2.0 {
			return &sarvamaierrors.ValidationError{
				Field:   "temperature",
				Message: "temperature must be between 0.01 and 2.0",
			}
		}
		r.Temperature = &temp
		return nil
	}
}

//
// Streaming Options
//

type streamOption func(*ttsStreamRequest) error

// WithStreamModel sets the TTS model for streaming.
// BulbulV2 or BulbulV3Beta
func WithStreamModel(model Model) streamOption {
	return func(c *ttsStreamRequest) error {
		if model == BulbulV3 {
			return &sarvamaierrors.ValidationError{
				Field:   "model",
				Message: "Bulbul:v3 is not supported for streaming. Use Bulbul:v2 or Bulbul:v3-beta.",
			}
		}
		c.Model = &model
		return nil
	}
}

// WithStreamSpeaker sets the voice for streaming synthesis.
func WithStreamSpeaker(voice SpeakerVoice) streamOption {
	return func(c *ttsStreamRequest) error {
		model := BulbulV3Beta
		if c.Model != nil {
			model = *c.Model
		}

		var allowedVoices []SpeakerVoice
		if model == BulbulV2 {
			allowedVoices = allowedSpeakerVoicesForBulbulV2
		} else {
			allowedVoices = allowedSpeakerVoicesForBulbulV3Beta
		}

		if slices.Contains(allowedVoices, voice) {
				c.Speaker = voice
				return nil
			}
		return &sarvamaierrors.ValidationError{
			Field:   "speaker_voice",
			Message: fmt.Sprintf("speaker voice %s is not supported for model %s", voice, model),
		}
	}
}

// WithStreamSendCompletionEvent enables completion events in streaming.
func WithStreamSendCompletionEvent(enable bool) streamOption {
	return func(c *ttsStreamRequest) error {
		c.SendCompletionEvent = &enable
		return nil
	}
}

// WithStreamPitch sets the speech pitch for streaming.
// Range: -0.75 to 0.75
func WithStreamPitch(pitch float64) streamOption {
	return func(c *ttsStreamRequest) error {
		if pitch < -0.75 || pitch > 0.75 {
			return &sarvamaierrors.ValidationError{
				Field:   "pitch",
				Message: "pitch must be between -0.75 and 0.75",
			}
		}
		c.Pitch = &pitch
		return nil
	}
}

// WithStreamPace sets the speech speed for streaming.
func WithStreamPace(pace float64) streamOption {
	return func(c *ttsStreamRequest) error {
		if pace < 0.3 || pace > 3.0 {
			return &sarvamaierrors.ValidationError{
				Field:   "pace",
				Message: "pace must be between 0.3 and 3.0",
			}
		}
		c.Pace = &pace
		return nil
	}
}

// WithStreamLoudness sets the speech volume for streaming.
// Range: 0.3 to 3.0
func WithStreamLoudness(loudness float64) streamOption {
	return func(c *ttsStreamRequest) error {
		if loudness < 0.3 || loudness > 3.0 {
			return &sarvamaierrors.ValidationError{
				Field:   "loudness",
				Message: "loudness must be between 0.3 and 3.0",
			}
		}
		c.Loudness = &loudness
		return nil
	}
}

// WithStreamTemperature sets the temperature for streaming.
// Range: 0.01 to 1.0
func WithStreamTemperature(temp float64) streamOption {
	return func(c *ttsStreamRequest) error {
		if temp < 0.01 || temp > 1.0 {
			return &sarvamaierrors.ValidationError{
				Field:   "temperature",
				Message: "temperature must be between 0.01 and 1.0",
			}
		}
		c.Temperature = &temp
		return nil
	}
}

// WithStreamSampleRate sets the sample rate for streaming.
func WithStreamSampleRate(rate SpeechSampleRate) streamOption {
	return func(c *ttsStreamRequest) error {
		c.SpeechSampleRate = &rate
		return nil
	}
}

func WithStreamEnablePreprocessing(enable bool) streamOption {
	return func(c *ttsStreamRequest) error {
		c.EnablePreprocessing = &enable
		return nil
	}
}

// WithStreamAudioCodec sets the output codec for streaming.
func WithStreamAudioCodec(codec AudioCodec) streamOption {
	return func(c *ttsStreamRequest) error {
		c.OutputAudioCodec = &codec
		return nil
	}
}

// WithStreamBitrate sets the output bitrate for streaming.
func WithStreamBitrate(bitrate Bitrate) streamOption {
	return func(c *ttsStreamRequest) error {
		c.OutputAudioBitrate = &bitrate
		return nil
	}
}

// WithMinBufferSize sets the minimum buffer size for streaming audio chunks.
// range between 30 - 200
func WithMinBufferSize(size int) streamOption {
	return func(c *ttsStreamRequest) error {
		if size < 30 || size > 200 {
			return &sarvamaierrors.ValidationError{
				Field:   "min_buffer_size",
				Message: "min_buffer_size must be between 30 and 200",
			}
		}
		c.MinBufferSize = &size
		return nil
	}
}

// WithMaxChunkSize sets the maximum chunk size for streaming audio chunks.
// Range is between 50 and 500.
func WithMaxChunkSize(size int) streamOption {
	return func(c *ttsStreamRequest) error {
		if size < 50 || size > 500 {
			return &sarvamaierrors.ValidationError{
				Field:   "max_chunk_length",
				Message: "max_chunk_length must be between 50 and 500",
			}
		}
		c.MaxChunkLength = &size
		return nil
	}
}
