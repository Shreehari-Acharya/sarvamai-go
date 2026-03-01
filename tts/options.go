package tts

//
// Option Type and Functions
//

type option func(*ttsRequest)

// WithModel sets the TTS model to use.
func WithModel(model Model) option {
	return func(r *ttsRequest) {
		r.Model = &model
	}
}

// WithSpeakerVoice sets the voice for synthesis.
func WithSpeakerVoice(voice SpeakerVoice) option {
	return func(r *ttsRequest) {
		r.SpeakerVoice = &voice
	}
}

// WithPitch sets the speech pitch adjustment.
// Note: Only supported for Bulbul:v2.
func WithPitch(pitch float64) option {
	return func(r *ttsRequest) {
		r.Pitch = &pitch
	}
}

// WithPace sets the speech speed.
// Note: Model-specific range (0.3-3.0 for v2, 0.5-2.0 for v3)
func WithPace(pace float64) option {
	return func(r *ttsRequest) {
		r.Pace = &pace
	}
}

// WithLoudness sets the speech volume.
// Note: Only supported for Bulbul:v2.
func WithLoudness(loudness float64) option {
	return func(r *ttsRequest) {
		r.Loudness = &loudness
	}
}

// WithSpeechSampleRate sets the audio sample rate.
func WithSpeechSampleRate(rate SpeechSampleRate) option {
	return func(r *ttsRequest) {
		r.SpeechSampleRate = &rate
	}
}

// WithOutputAudioCodec sets the output audio format.
func WithOutputAudioCodec(codec AudioCodec) option {
	return func(r *ttsRequest) {
		r.OutputAudioCodec = &codec
	}
}

// WithEnablePreprocessing enables text preprocessing.
func WithEnablePreprocessing(enable bool) option {
	return func(r *ttsRequest) {
		r.EnablePreprocessing = &enable
	}
}

// WithTemperature controls how much randomness and expressiveness the TTS model uses.
// Note: Only supported for Bulbul:v3.
func WithTemperature(temp float64) option {
	return func(r *ttsRequest) {
		r.Temperature = &temp
	}
}

//
// Streaming Options
//

type streamOption func(*ttsStreamRequest)

// WithStreamModel sets the TTS model for streaming.
// BulbulV2 or BulbulV3Beta
func WithStreamModel(model Model) streamOption {
	return func(c *ttsStreamRequest) {
		c.Model = &model
	}
}

// WithStreamSpeaker sets the voice for streaming synthesis.
func WithStreamSpeaker(voice SpeakerVoice) streamOption {
	return func(c *ttsStreamRequest) {
		c.Speaker = voice
	}
}

// WithStreamSendCompletionEvent enables completion events in streaming.
func WithStreamSendCompletionEvent(enable bool) streamOption {
	return func(c *ttsStreamRequest) {
		c.SendCompletionEvent = &enable
	}
}

// WithStreamPitch sets the speech pitch for streaming.
// Note: Only supported for Bulbul:v2.
func WithStreamPitch(pitch float64) streamOption {
	return func(c *ttsStreamRequest) {
		c.Pitch = &pitch
	}
}

// WithStreamPace sets the speech speed for streaming.
// Note: Model-specific range (0.3-3.0 for v2, 0.5-2.0 for v3-beta)
func WithStreamPace(pace float64) streamOption {
	return func(c *ttsStreamRequest) {
		c.Pace = &pace
	}
}

// WithStreamLoudness sets the speech volume for streaming.
// Note: Only supported for Bulbul:v2.
func WithStreamLoudness(loudness float64) streamOption {
	return func(c *ttsStreamRequest) {
		c.Loudness = &loudness
	}
}

// WithStreamTemperature sets the temperature for streaming.
// Note: Only supported for Bulbul:v3-beta.
func WithStreamTemperature(temp float64) streamOption {
	return func(c *ttsStreamRequest) {
		c.Temperature = &temp
	}
}

// WithStreamSampleRate sets the sample rate for streaming.
func WithStreamSampleRate(rate SpeechSampleRate) streamOption {
	return func(c *ttsStreamRequest) {
		c.SpeechSampleRate = &rate
	}
}

// WithStreamEnablePreprocessing enables text preprocessing for streaming.
func WithStreamEnablePreprocessing(enable bool) streamOption {
	return func(c *ttsStreamRequest) {
		c.EnablePreprocessing = &enable
	}
}

// WithStreamAudioCodec sets the output codec for streaming.
func WithStreamAudioCodec(codec AudioCodec) streamOption {
	return func(c *ttsStreamRequest) {
		c.OutputAudioCodec = &codec
	}
}

// WithStreamBitrate sets the output bitrate for streaming.
func WithStreamBitrate(bitrate Bitrate) streamOption {
	return func(c *ttsStreamRequest) {
		c.OutputAudioBitrate = &bitrate
	}
}

// WithMinBufferSize sets the minimum buffer size for streaming audio chunks.
// range between 30 - 200
func WithMinBufferSize(size int) streamOption {
	return func(c *ttsStreamRequest) {
		c.MinBufferSize = &size
	}
}

// WithMaxChunkSize sets the maximum chunk size for streaming audio chunks.
// Range is between 50 and 500.
func WithMaxChunkSize(size int) streamOption {
	return func(c *ttsStreamRequest) {
		c.MaxChunkLength = &size
	}
}
