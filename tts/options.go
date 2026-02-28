package tts

import (
	"fmt"
	"slices"

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
// Note: Only supported for Bulbul:v2.
func WithPitch(pitch float64) option {
	return func(r *ttsRequest) error {
		r.Pitch = &pitch
		return nil
	}
}

// WithPace sets the speech speed.
// Note: Model-specific range (0.3-3.0 for v2, 0.5-2.0 for v3)
func WithPace(pace float64) option {
	return func(r *ttsRequest) error {
		r.Pace = &pace
		return nil
	}
}

// WithLoudness sets the speech volume.
// Note: Only supported for Bulbul:v2.
func WithLoudness(loudness float64) option {
	return func(r *ttsRequest) error {
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
// Note: Only supported for Bulbul:v3.
func WithTemperature(temp float64) option {
	return func(r *ttsRequest) error {
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
// Note: Only supported for Bulbul:v2.
func WithStreamPitch(pitch float64) streamOption {
	return func(c *ttsStreamRequest) error {
		c.Pitch = &pitch
		return nil
	}
}

// WithStreamPace sets the speech speed for streaming.
// Note: Model-specific range (0.3-3.0 for v2, 0.5-2.0 for v3-beta)
func WithStreamPace(pace float64) streamOption {
	return func(c *ttsStreamRequest) error {
		c.Pace = &pace
		return nil
	}
}

// WithStreamLoudness sets the speech volume for streaming.
// Note: Only supported for Bulbul:v2.
func WithStreamLoudness(loudness float64) streamOption {
	return func(c *ttsStreamRequest) error {
		c.Loudness = &loudness
		return nil
	}
}

// WithStreamTemperature sets the temperature for streaming.
// Note: Only supported for Bulbul:v3-beta.
func WithStreamTemperature(temp float64) streamOption {
	return func(c *ttsStreamRequest) error {
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

// WithStreamEnablePreprocessing enables text preprocessing for streaming.
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
