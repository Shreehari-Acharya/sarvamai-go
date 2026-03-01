package stt

import (
	"strings"
	"testing"

	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/sarvamaierrors"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/languages"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/shared/speech"
)

func ptrModel(m speech.Model) *speech.Model {
	return &m
}

func ptrMode(m speech.Mode) *speech.Mode {
	return &m
}

func ptrLanguageCode(c languages.Code) *languages.Code {
	return &c
}

func ptrCodec(c speech.InputAudioCodec) *speech.InputAudioCodec {
	return &c
}

func ptrSampleRate(r speech.StreamSampleRate) *speech.StreamSampleRate {
	return &r
}

func ptrBool(b bool) *bool {
	return &b
}

type testFileReader struct{}

func (t *testFileReader) Read(p []byte) (int, error) {
	return 0, nil
}

func TestValidateTranscribeRequest(t *testing.T) {
	tests := []struct {
		name    string
		req     *transcribeRequest
		wantErr bool
		errMsg  string
	}{
		// File validation tests
		{
			name:    "nil file should error",
			req:     &transcribeRequest{File: nil},
			wantErr: true,
			errMsg:  "file is required",
		},
		{
			name:    "valid file reader should pass",
			req:     &transcribeRequest{File: &testFileReader{}},
			wantErr: false,
		},

		// Model tests
		{
			name: "nil model should pass",
			req: &transcribeRequest{
				File:  &testFileReader{},
				Model: nil,
			},
			wantErr: false,
		},
		{
			name: "saarika model should pass",
			req: &transcribeRequest{
				File:  &testFileReader{},
				Model: ptrModel(speech.ModelSaarika),
			},
			wantErr: false,
		},
		{
			name: "saaras model should pass",
			req: &transcribeRequest{
				File:  &testFileReader{},
				Model: ptrModel(speech.ModelSaaras),
			},
			wantErr: false,
		},

		// Mode tests
		{
			name: "nil mode should pass",
			req: &transcribeRequest{
				File: &testFileReader{},
				Mode: nil,
			},
			wantErr: false,
		},
		{
			name: "mode with nil model should pass (defaults to saaras which supports mode)",
			req: &transcribeRequest{
				File:  &testFileReader{},
				Model: nil,
				Mode:  ptrMode(speech.ModeTranscribe),
			},
			wantErr: false,
		},
		{
			name: "mode with saarika should error",
			req: &transcribeRequest{
				File:  &testFileReader{},
				Model: ptrModel(speech.ModelSaarika),
				Mode:  ptrMode(speech.ModeTranscribe),
			},
			wantErr: true,
			errMsg:  "mode is only supported with saaras:v3 model",
		},
		{
			name: "mode with saaras should pass",
			req: &transcribeRequest{
				File:  &testFileReader{},
				Model: ptrModel(speech.ModelSaaras),
				Mode:  ptrMode(speech.ModeTranscribe),
			},
			wantErr: false,
		},
		{
			name: "translate mode with saaras should pass",
			req: &transcribeRequest{
				File:  &testFileReader{},
				Model: ptrModel(speech.ModelSaaras),
				Mode:  ptrMode(speech.ModeTranslate),
			},
			wantErr: false,
		},
		{
			name: "verbatim mode with saaras should pass",
			req: &transcribeRequest{
				File:  &testFileReader{},
				Model: ptrModel(speech.ModelSaaras),
				Mode:  ptrMode(speech.ModeVerbatim),
			},
			wantErr: false,
		},
		{
			name: "translit mode with saaras should pass",
			req: &transcribeRequest{
				File:  &testFileReader{},
				Model: ptrModel(speech.ModelSaaras),
				Mode:  ptrMode(speech.ModeTranslit),
			},
			wantErr: false,
		},
		{
			name: "codemix mode with saaras should pass",
			req: &transcribeRequest{
				File:  &testFileReader{},
				Model: ptrModel(speech.ModelSaaras),
				Mode:  ptrMode(speech.ModeCodemix),
			},
			wantErr: false,
		},

		// Language tests - key edge cases
		{
			name: "nil language should pass (auto-detect)",
			req: &transcribeRequest{
				File:     &testFileReader{},
				Language: nil,
			},
			wantErr: false,
		},
		{
			name: "language with nil model should validate against saaras (default)",
			req: &transcribeRequest{
				File:     &testFileReader{},
				Model:    nil,
				Language: ptrLanguageCode(languages.Code("hi-IN")),
			},
			wantErr: false,
		},
		{
			name: "language hi-IN supported by saarika should pass",
			req: &transcribeRequest{
				File:     &testFileReader{},
				Model:    ptrModel(speech.ModelSaarika),
				Language: ptrLanguageCode(languages.Code("hi-IN")),
			},
			wantErr: false,
		},
		{
			name: "language en-IN supported by saarika should pass",
			req: &transcribeRequest{
				File:     &testFileReader{},
				Model:    ptrModel(speech.ModelSaarika),
				Language: ptrLanguageCode(languages.Code("en-IN")),
			},
			wantErr: false,
		},
		{
			name: "language as-IN not supported by saarika should error",
			req: &transcribeRequest{
				File:     &testFileReader{},
				Model:    ptrModel(speech.ModelSaarika),
				Language: ptrLanguageCode(languages.Code("as-IN")),
			},
			wantErr: true,
			errMsg:  "as-IN is not supported by saarika:v2.5 model",
		},
		{
			name: "language as-IN supported by saaras should pass",
			req: &transcribeRequest{
				File:     &testFileReader{},
				Model:    ptrModel(speech.ModelSaaras),
				Language: ptrLanguageCode(languages.Code("as-IN")),
			},
			wantErr: false,
		},
		{
			name: "language ur-IN supported by saaras should pass",
			req: &transcribeRequest{
				File:     &testFileReader{},
				Model:    ptrModel(speech.ModelSaaras),
				Language: ptrLanguageCode(languages.Code("ur-IN")),
			},
			wantErr: false,
		},
		{
			name: "language ne-IN supported by saaras should pass",
			req: &transcribeRequest{
				File:     &testFileReader{},
				Model:    ptrModel(speech.ModelSaaras),
				Language: ptrLanguageCode(languages.Code("ne-IN")),
			},
			wantErr: false,
		},
		{
			name: "language not in saaras set should error",
			req: &transcribeRequest{
				File:     &testFileReader{},
				Language: ptrLanguageCode(languages.Code("fr-FR")),
			},
			wantErr: true,
			errMsg:  "fr-FR is not supported by",
		},
		{
			name: "language zh-CN not in saaras set should error",
			req: &transcribeRequest{
				File:     &testFileReader{},
				Language: ptrLanguageCode(languages.Code("zh-CN")),
			},
			wantErr: true,
			errMsg:  "zh-CN is not supported by",
		},

		// Complete valid requests
		{
			name: "all nil except file should pass",
			req: &transcribeRequest{
				File:     &testFileReader{},
				Model:    nil,
				Mode:     nil,
				Language: nil,
			},
			wantErr: false,
		},
		{
			name: "complete valid request with saaras model",
			req: &transcribeRequest{
				File:     &testFileReader{},
				Model:    ptrModel(speech.ModelSaaras),
				Mode:     ptrMode(speech.ModeTranscribe),
				Language: ptrLanguageCode(languages.Code("hi-IN")),
			},
			wantErr: false,
		},
		{
			name: "complete valid request with saarika model",
			req: &transcribeRequest{
				File:     &testFileReader{},
				Model:    ptrModel(speech.ModelSaarika),
				Mode:     nil,
				Language: ptrLanguageCode(languages.Code("ta-IN")),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateTranscribeRequest(tt.req)
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got nil")
					return
				}
				if tt.errMsg != "" && !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("got %q, want to contain %q", err.Error(), tt.errMsg)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			}
		})
	}
}

func TestValidateStreamConfig(t *testing.T) {
	tests := []struct {
		name    string
		cfg     *streamTranscribeRequest
		wantErr bool
		errMsg  string
	}{
		// Language tests - stream config uses languages.Code directly (not pointer)
		{
			name: "empty language should pass (will default to unknown)",
			cfg: &streamTranscribeRequest{
				Language: "",
			},
			wantErr: false,
		},
		{
			name: "valid language hi-IN should pass",
			cfg: &streamTranscribeRequest{
				Language: languages.Code("hi-IN"),
			},
			wantErr: false,
		},
		{
			name: "valid language unknown should pass",
			cfg: &streamTranscribeRequest{
				Language: languages.Code("unknown"),
			},
			wantErr: false,
		},
		{
			name: "language as-IN should pass (saaras supported)",
			cfg: &streamTranscribeRequest{
				Language: languages.Code("as-IN"),
			},
			wantErr: false,
		},
		{
			name: "invalid language fr-FR should error",
			cfg: &streamTranscribeRequest{
				Language: languages.Code("fr-FR"),
			},
			wantErr: true,
			errMsg:  "fr-FR is not supported by",
		},

		// Model tests
		{
			name: "nil model should pass",
			cfg: &streamTranscribeRequest{
				Model: nil,
			},
			wantErr: false,
		},
		{
			name: "saarika model should pass",
			cfg: &streamTranscribeRequest{
				Model: ptrModel(speech.ModelSaarika),
			},
			wantErr: false,
		},
		{
			name: "saaras model should pass",
			cfg: &streamTranscribeRequest{
				Model: ptrModel(speech.ModelSaaras),
			},
			wantErr: false,
		},

		// Mode + Model combination
		{
			name: "nil mode should pass",
			cfg: &streamTranscribeRequest{
				Mode: nil,
			},
			wantErr: false,
		},
		{
			name: "mode with nil model should pass",
			cfg: &streamTranscribeRequest{
				Model: nil,
				Mode:  ptrMode(speech.ModeTranscribe),
			},
			wantErr: false,
		},
		{
			name: "mode with saarika should error",
			cfg: &streamTranscribeRequest{
				Model: ptrModel(speech.ModelSaarika),
				Mode:  ptrMode(speech.ModeTranscribe),
			},
			wantErr: true,
			errMsg:  "mode is only supported with saaras:v3 model",
		},
		{
			name: "mode with saaras should pass",
			cfg: &streamTranscribeRequest{
				Model: ptrModel(speech.ModelSaaras),
				Mode:  ptrMode(speech.ModeTranscribe),
			},
			wantErr: false,
		},

		// Complete valid configs
		{
			name: "complete valid stream config",
			cfg: &streamTranscribeRequest{
				Language: languages.Code("hi-IN"),
				Model:    ptrModel(speech.ModelSaaras),
				Mode:     ptrMode(speech.ModeTranslate),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateStreamConfig(tt.cfg)
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got nil")
					return
				}
				if tt.errMsg != "" && !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("got %q, want to contain %q", err.Error(), tt.errMsg)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			}
		})
	}
}

func TestWithModel(t *testing.T) {
	tests := []struct {
		name    string
		model   speech.Model
		wantErr bool
		errMsg  string
	}{
		{
			name:    "valid model saarika",
			model:   speech.ModelSaarika,
			wantErr: false,
		},
		{
			name:    "valid model saaras",
			model:   speech.ModelSaaras,
			wantErr: false,
		},
		{
			name:    "invalid model should error",
			model:   speech.Model("invalid"),
			wantErr: true,
			errMsg:  "invalid model",
		},
		{
			name:    "empty model should error",
			model:   speech.Model(""),
			wantErr: true,
			errMsg:  "invalid model",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &transcribeRequest{File: &testFileReader{}}
			opt := WithModel(tt.model)
			opt(req)
			err := validateTranscribeRequest(req)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got nil")
					return
				}
				if tt.errMsg != "" && !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("got %q, want to contain %q", err.Error(), tt.errMsg)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if req.Model == nil || *req.Model != tt.model {
					t.Errorf("expected model %v, got %v", tt.model, req.Model)
				}
			}
		})
	}
}

func TestWithLanguage(t *testing.T) {
	tests := []struct {
		name    string
		lang    languages.Code
		wantErr bool
		errMsg  string
	}{
		{
			name:    "valid language hi-IN",
			lang:    languages.Code("hi-IN"),
			wantErr: false,
		},
		{
			name:    "valid language en-IN",
			lang:    languages.Code("en-IN"),
			wantErr: false,
		},
		{
			name:    "valid language as-IN (saaras only)",
			lang:    languages.Code("as-IN"),
			wantErr: false,
		},
		{
			name:    "valid language unknown (auto-detect)",
			lang:    languages.Code("unknown"),
			wantErr: false,
		},
		{
			name:    "empty language returns nil (no-op)",
			lang:    "",
			wantErr: false,
		},
		{
			name:    "invalid language fr-FR should error",
			lang:    languages.Code("fr-FR"),
			wantErr: true,
			errMsg:  "is not supported by",
		},
		{
			name:    "invalid language zh-CN should error",
			lang:    languages.Code("zh-CN"),
			wantErr: true,
			errMsg:  "is not supported by",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &transcribeRequest{File: &testFileReader{}}
			opt := WithLanguage(tt.lang)
			opt(req)
			err := validateTranscribeRequest(req)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got nil")
					return
				}
				if tt.errMsg != "" && !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("got %q, want to contain %q", err.Error(), tt.errMsg)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if tt.lang != "" {
					if req.Language == nil || *req.Language != tt.lang {
						t.Errorf("expected language %v, got %v", tt.lang, req.Language)
					}
				}
			}
		})
	}
}

func TestWithStreamSampleRate(t *testing.T) {
	tests := []struct {
		name    string
		rate    speech.StreamSampleRate
		wantErr bool
		errMsg  string
	}{
		{
			name:    "valid sample rate 8000",
			rate:    speech.SampleRate8000,
			wantErr: false,
		},
		{
			name:    "valid sample rate 16000",
			rate:    speech.SampleRate16000,
			wantErr: false,
		},
		{
			name:    "invalid sample rate 22050 should error",
			rate:    speech.SampleRate22050,
			wantErr: true,
			errMsg:  "invalid sample rate for streaming",
		},
		{
			name:    "invalid sample rate 24000 should error",
			rate:    speech.SampleRate24000,
			wantErr: true,
			errMsg:  "invalid sample rate for streaming",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &streamTranscribeRequest{}
			opt := WithStreamSampleRate(tt.rate)
			opt(cfg)
			err := validateStreamConfig(cfg)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got nil")
					return
				}
				if tt.errMsg != "" && !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("got %q, want to contain %q", err.Error(), tt.errMsg)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if cfg.SampleRate == nil || *cfg.SampleRate != tt.rate {
					t.Errorf("expected sample rate %v, got %v", tt.rate, cfg.SampleRate)
				}
			}
		})
	}
}

func TestWithStreamInputAudioCodec(t *testing.T) {
	tests := []struct {
		name    string
		codec   speech.InputAudioCodec
		wantErr bool
		errMsg  string
	}{
		{
			name:    "valid codec wav",
			codec:   speech.CodecWAV,
			wantErr: false,
		},
		{
			name:    "valid codec pcm_s16le",
			codec:   speech.CodecPCMS16LE,
			wantErr: false,
		},
		{
			name:    "valid codec pcm_l16",
			codec:   speech.CodecPCML16,
			wantErr: false,
		},
		{
			name:    "valid codec pcm_raw",
			codec:   speech.CodecPCMRAW,
			wantErr: false,
		},
		{
			name:    "invalid codec mp3 should error",
			codec:   speech.CodecMP3,
			wantErr: true,
			errMsg:  "unsupported audio codec for streaming",
		},
		{
			name:    "invalid codec flac should error",
			codec:   speech.CodecFLAC,
			wantErr: true,
			errMsg:  "unsupported audio codec for streaming",
		},
		{
			name:    "invalid codec aac should error",
			codec:   speech.CodecAAC,
			wantErr: true,
			errMsg:  "unsupported audio codec for streaming",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &streamTranscribeRequest{}
			opt := WithStreamInputAudioCodec(tt.codec)
			opt(cfg)
			err := validateStreamConfig(cfg)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got nil")
					return
				}
				if tt.errMsg != "" && !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("got %q, want to contain %q", err.Error(), tt.errMsg)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if cfg.InputAudioCodec == nil || *cfg.InputAudioCodec != tt.codec {
					t.Errorf("expected codec %v, got %v", tt.codec, cfg.InputAudioCodec)
				}
			}
		})
	}
}

func TestWithAudioCodec(t *testing.T) {
	tests := []struct {
		name  string
		codec speech.InputAudioCodec
	}{
		{
			name:  "valid codec wav",
			codec: speech.CodecWAV,
		},
		{
			name:  "valid codec mp3",
			codec: speech.CodecMP3,
		},
		{
			name:  "valid codec flac",
			codec: speech.CodecFLAC,
		},
		{
			name:  "valid codec pcm_s16le",
			codec: speech.CodecPCMS16LE,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &transcribeRequest{File: &testFileReader{}}
			opt := WithAudioCodec(tt.codec)
			opt(req)
			err := validateTranscribeRequest(req)

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if req.AudioCodec == nil || *req.AudioCodec != tt.codec {
				t.Errorf("expected codec %v, got %v", tt.codec, req.AudioCodec)
			}
		})
	}
}

func TestValidationErrorInterface(t *testing.T) {
	t.Run("ValidationError implements error interface", func(t *testing.T) {
		err := &sarvamaierrors.ValidationError{
			Field:   "test_field",
			Message: "test error message",
		}

		expected := "test_field: test error message"
		if err.Error() != expected {
			t.Errorf("expected %q, got %q", expected, err.Error())
		}
	})
}
