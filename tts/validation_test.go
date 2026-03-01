package tts

import (
	"testing"
)

func TestOptionValidation(t *testing.T) {
	t.Run("valid bulbul:v3 request", func(t *testing.T) {
		req := &ttsRequest{
			Text:               "Hello",
			TargetLanguageCode: "en-IN",
		}
		opts := []option{
			WithModel(BulbulV3),
			WithSpeakerVoice(SpeakerShubh),
		}
		for _, opt := range opts {
			opt(req)
		}
		if err := validateTTSRequest(req); err != nil {
			t.Errorf("validation failed: %v", err)
		}
	})

	t.Run("valid bulbul:v2 request", func(t *testing.T) {
		req := &ttsRequest{
			Text:               "Hello",
			TargetLanguageCode: "en-IN",
		}
		opts := []option{
			WithModel(BulbulV2),
			WithSpeakerVoice(SpeakerAnushka),
			WithPitch(0.0),
			WithLoudness(1.0),
			WithPace(1.0),
		}
		for _, opt := range opts {
			opt(req)
		}
		if err := validateTTSRequest(req); err != nil {
			t.Errorf("validation failed: %v", err)
		}
	})

	t.Run("default model is bulbul:v3", func(t *testing.T) {
		req := &ttsRequest{
			Text:               "Hello",
			TargetLanguageCode: "en-IN",
		}
		opts := []option{
			WithSpeakerVoice(SpeakerShubh),
		}
		for _, opt := range opts {
			opt(req)
		}
		if err := validateTTSRequest(req); err != nil {
			t.Errorf("validation failed: %v", err)
		}
	})
}

func TestOptionValidationBulbulV3Errors(t *testing.T) {
	tests := []struct {
		name    string
		opts    []option
		wantErr string
	}{
		{
			name: "bulbul:v3 with pitch should error",
			opts: []option{
				WithModel(BulbulV3),
				WithPitch(0.5),
			},
			wantErr: "pitch: pitch is not supported for bulbul:v3",
		},
		{
			name: "bulbul:v3 with loudness should error",
			opts: []option{
				WithModel(BulbulV3),
				WithLoudness(1.0),
			},
			wantErr: "loudness: loudness is not supported for bulbul:v3",
		},
		{
			name: "bulbul:v3 with preprocessing should error",
			opts: []option{
				WithModel(BulbulV3),
				WithEnablePreprocessing(true),
			},
			wantErr: "enable_preprocessing: enable_preprocessing is not supported for bulbul:v3",
		},
		{
			name: "bulbul:v2 speaker with bulbul:v3 should error",
			opts: []option{
				WithModel(BulbulV3),
				WithSpeakerVoice(SpeakerAnushka),
			},
			wantErr: "speaker_voice: speaker voice anushka is not supported for model bulbul:v3",
		},
		{
			name: "bulbul:v3 speaker with bulbul:v2 should error",
			opts: []option{
				WithModel(BulbulV2),
				WithSpeakerVoice(SpeakerShubh),
			},
			wantErr: "speaker_voice: speaker voice shubh is not supported for model bulbul:v2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &ttsRequest{
				Text:               "Hello",
				TargetLanguageCode: "en-IN",
			}
			for _, opt := range tt.opts {
				opt(req)
			}
			err := validateTTSRequest(req)
			if err == nil {
				t.Errorf("expected error but got nil")
				return
			}
			if err.Error() != tt.wantErr {
				t.Errorf("got %q, want %q", err.Error(), tt.wantErr)
			}
		})
	}
}

func TestOptionValidationBulbulV2Errors(t *testing.T) {
	tests := []struct {
		name    string
		opts    []option
		wantErr string
	}{
		{
			name: "bulbul:v2 with pitch out of range",
			opts: []option{
				WithModel(BulbulV2),
				WithPitch(1.0),
			},
			wantErr: "pitch: pitch must be between -0.75 and 0.75",
		},
		{
			name: "bulbul:v2 with loudness out of range",
			opts: []option{
				WithModel(BulbulV2),
				WithLoudness(4.0),
			},
			wantErr: "loudness: loudness must be between 0.3 and 3.0",
		},
		{
			name: "bulbul:v2 with temperature should error",
			opts: []option{
				WithModel(BulbulV2),
				WithTemperature(0.5),
			},
			wantErr: "temperature: temperature is not supported for bulbul:v2",
		},
		{
			name: "bulbul:v2 with sample rate > 24000 should error",
			opts: []option{
				WithModel(BulbulV2),
				WithSpeechSampleRate(SampleRate48000),
			},
			wantErr: "speech_sample_rate: sample rate must be 8000, 16000, 22050, or 24000 for bulbul:v2",
		},
		{
			name: "bulbul:v2 with pace out of range (too fast)",
			opts: []option{
				WithModel(BulbulV2),
				WithPace(3.5),
			},
			wantErr: "pace: pace must be between 0.3 and 3.0 for bulbul:v2",
		},
		{
			name: "bulbul:v2 with pace out of range (too slow)",
			opts: []option{
				WithModel(BulbulV2),
				WithPace(0.1),
			},
			wantErr: "pace: pace must be between 0.3 and 3.0 for bulbul:v2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &ttsRequest{
				Text:               "Hello",
				TargetLanguageCode: "en-IN",
			}
			for _, opt := range tt.opts {
				opt(req)
			}
			err := validateTTSRequest(req)
			if err == nil {
				t.Errorf("expected error but got nil")
				return
			}
			if err.Error() != tt.wantErr {
				t.Errorf("got %q, want %q", err.Error(), tt.wantErr)
			}
		})
	}
}

func TestOptionValidationBulbulV3(t *testing.T) {
	tests := []struct {
		name    string
		opts    []option
		wantErr string
	}{
		{
			name: "bulbul:v3 with valid temperature",
			opts: []option{
				WithModel(BulbulV3),
				WithTemperature(0.5),
			},
			wantErr: "",
		},
		{
			name: "bulbul:v3 with temperature out of range",
			opts: []option{
				WithModel(BulbulV3),
				WithTemperature(3.0), // range is 0.01 to 2.0
			},
			wantErr: "temperature: temperature must be between 0.01 and 2.0",
		},
		{
			name: "bulbul:v3 with sample rate 48000 should work",
			opts: []option{
				WithModel(BulbulV3),
				WithSpeechSampleRate(SampleRate48000),
			},
			wantErr: "",
		},
		{
			name: "bulbul:v3 with pace out of range (too fast)",
			opts: []option{
				WithModel(BulbulV3),
				WithPace(2.5), // range is 0.5 to 2.0
			},
			wantErr: "pace: pace must be between 0.5 and 2.0 for bulbul:v3",
		},
		{
			name: "bulbul:v3 with pace out of range (too slow)",
			opts: []option{
				WithModel(BulbulV3),
				WithPace(0.3), // range is 0.5 to 2.0
			},
			wantErr: "pace: pace must be between 0.5 and 2.0 for bulbul:v3",
		},
		{
			name: "bulbul:v3 with valid pace",
			opts: []option{
				WithModel(BulbulV3),
				WithPace(1.0),
			},
			wantErr: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &ttsRequest{
				Text:               "Hello",
				TargetLanguageCode: "en-IN",
			}
			for _, opt := range tt.opts {
				opt(req)
			}
			err := validateTTSRequest(req)
			if tt.wantErr == "" {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
				}
			} else {
				if err == nil {
					t.Errorf("expected error %q but got nil", tt.wantErr)
					return
				}
				if err.Error() != tt.wantErr {
					t.Errorf("got %q, want %q", err.Error(), tt.wantErr)
				}
			}
		})
	}
}

func TestStreamOptionValidation(t *testing.T) {
	t.Run("valid bulbul:v2 stream config", func(t *testing.T) {
		cfg := &ttsStreamRequest{
			TargetLanguageCode: "hi-IN",
			Speaker:            SpeakerAnushka,
		}
		opts := []streamOption{
			WithStreamModel(BulbulV2),
		}
		for _, opt := range opts {
			opt(cfg)
		}
		if err := validateTTSStreamRequest(cfg); err != nil {
			t.Errorf("validation failed: %v", err)
		}
	})

	t.Run("valid bulbul:v3-beta stream config", func(t *testing.T) {
		cfg := &ttsStreamRequest{
			TargetLanguageCode: "hi-IN",
			Speaker:            SpeakerShubh,
		}
		opts := []streamOption{
			WithStreamModel(BulbulV3Beta),
		}
		for _, opt := range opts {
			opt(cfg)
		}
		if err := validateTTSStreamRequest(cfg); err != nil {
			t.Errorf("validation failed: %v", err)
		}
	})

	t.Run("bulbul:v2 with pitch should succeed", func(t *testing.T) {
		cfg := &ttsStreamRequest{
			TargetLanguageCode: "hi-IN",
			Speaker:            SpeakerAnushka,
		}
		opts := []streamOption{
			WithStreamModel(BulbulV2),
			WithStreamPitch(0.5),
		}
		for _, opt := range opts {
			opt(cfg)
		}
		if err := validateTTSStreamRequest(cfg); err != nil {
			t.Errorf("validation failed: %v", err)
		}
	})

	t.Run("bulbul:v2 with loudness should succeed", func(t *testing.T) {
		cfg := &ttsStreamRequest{
			TargetLanguageCode: "hi-IN",
			Speaker:            SpeakerAnushka,
		}
		opts := []streamOption{
			WithStreamModel(BulbulV2),
			WithStreamLoudness(1.0),
		}
		for _, opt := range opts {
			opt(cfg)
		}
		if err := validateTTSStreamRequest(cfg); err != nil {
			t.Errorf("validation failed: %v", err)
		}
	})

	t.Run("bulbul:v3-beta with pitch should error", func(t *testing.T) {
		cfg := &ttsStreamRequest{
			TargetLanguageCode: "hi-IN",
			Speaker:            SpeakerShubh,
		}
		opts := []streamOption{
			WithStreamModel(BulbulV3Beta),
			WithStreamPitch(0.5),
		}
		for _, opt := range opts {
			opt(cfg)
		}
		err := validateTTSStreamRequest(cfg)
		if err == nil {
			t.Errorf("expected error but got nil")
		}
		if err.Error() != "pitch: pitch is not supported for bulbul:v3-beta" {
			t.Errorf("got %q", err.Error())
		}
	})

	t.Run("bulbul:v3-beta with loudness should error", func(t *testing.T) {
		cfg := &ttsStreamRequest{
			TargetLanguageCode: "hi-IN",
			Speaker:            SpeakerShubh,
		}
		opts := []streamOption{
			WithStreamModel(BulbulV3Beta),
			WithStreamLoudness(1.0),
		}
		for _, opt := range opts {
			opt(cfg)
		}
		err := validateTTSStreamRequest(cfg)
		if err == nil {
			t.Errorf("expected error but got nil")
		}
		if err.Error() != "loudness: loudness is not supported for bulbul:v3-beta" {
			t.Errorf("got %q", err.Error())
		}
	})

	t.Run("bulbul:v2 with temperature should error", func(t *testing.T) {
		cfg := &ttsStreamRequest{
			TargetLanguageCode: "hi-IN",
			Speaker:            SpeakerAnushka,
		}
		opts := []streamOption{
			WithStreamModel(BulbulV2),
			WithStreamTemperature(0.5),
		}
		for _, opt := range opts {
			opt(cfg)
		}
		err := validateTTSStreamRequest(cfg)
		if err == nil {
			t.Errorf("expected error but got nil")
		}
		if err.Error() != "temperature: temperature is not supported for bulbul:v2" {
			t.Errorf("got %q", err.Error())
		}
	})

	t.Run("bulbul:v3-beta with sample rate > 24000 should error", func(t *testing.T) {
		cfg := &ttsStreamRequest{
			TargetLanguageCode: "hi-IN",
			Speaker:            SpeakerShubh,
		}
		opts := []streamOption{
			WithStreamModel(BulbulV3Beta),
			WithStreamSampleRate(SampleRate48000),
		}
		for _, opt := range opts {
			opt(cfg)
		}
		err := validateTTSStreamRequest(cfg)
		if err == nil {
			t.Errorf("expected error but got nil")
		}
		if err.Error() != "speech_sample_rate: sample rate must be 8000, 16000, 22050, or 24000 for bulbul:v3-beta" {
			t.Errorf("got %q", err.Error())
		}
	})

	t.Run("bulbul:v2 with sample rate > 24000 should error", func(t *testing.T) {
		cfg := &ttsStreamRequest{
			TargetLanguageCode: "hi-IN",
			Speaker:            SpeakerAnushka,
		}
		opts := []streamOption{
			WithStreamModel(BulbulV2),
			WithStreamSampleRate(SampleRate48000),
		}
		for _, opt := range opts {
			opt(cfg)
		}
		err := validateTTSStreamRequest(cfg)
		if err == nil {
			t.Errorf("expected error but got nil")
		}
		if err.Error() != "speech_sample_rate: sample rate must be 8000, 16000, 22050, or 24000 for bulbul:v2" {
			t.Errorf("got %q", err.Error())
		}
	})
}

func TestTargetLanguageValidation(t *testing.T) {
	tests := []struct {
		name    string
		lang    string
		wantErr string
	}{
		{
			name:    "valid language hi-IN",
			lang:    "hi-IN",
			wantErr: "",
		},
		{
			name:    "valid language en-IN",
			lang:    "en-IN",
			wantErr: "",
		},
		{
			name:    "invalid language",
			lang:    "xx-XX",
			wantErr: "target_language_code: target_language_code is required and must be a supported language code",
		},
		{
			name:    "empty language",
			lang:    "",
			wantErr: "target_language_code: target_language_code is required and must be a supported language code",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateTargetLanguage(tt.lang)
			if tt.wantErr == "" {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
				}
			} else {
				if err == nil {
					t.Errorf("expected error %q but got nil", tt.wantErr)
				} else if err.Error() != tt.wantErr {
					t.Errorf("got %q, want %q", err.Error(), tt.wantErr)
				}
			}
		})
	}
}

func TestTextValidation(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		model   Model
		wantErr string
	}{
		{
			name:    "valid text for bulbul:v3",
			text:    "Hello world",
			model:   BulbulV3,
			wantErr: "",
		},
		{
			name:    "valid text for bulbul:v2",
			text:    "Hello world",
			model:   BulbulV2,
			wantErr: "",
		},
		{
			name:    "empty text",
			text:    "",
			model:   BulbulV3,
			wantErr: "text: text must be between 1 and 2500 characters",
		},
		{
			name:    "text too long for bulbul:v3",
			text:    string(make([]byte, 2501)),
			model:   BulbulV3,
			wantErr: "text: text must be between 1 and 2500 characters",
		},
		{
			name:    "text too long for bulbul:v2",
			text:    string(make([]byte, 1501)),
			model:   BulbulV2,
			wantErr: "text: text must be between 1 and 1500 characters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateText(tt.text, tt.model)
			if tt.wantErr == "" {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
				}
			} else {
				if err == nil {
					t.Errorf("expected error %q but got nil", tt.wantErr)
				} else if err.Error() != tt.wantErr {
					t.Errorf("got %q, want %q", err.Error(), tt.wantErr)
				}
			}
		})
	}
}
