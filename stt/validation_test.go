package stt

import (
	"strings"
	"testing"

	"github.com/Shreehari-Acharya/sarvam-go-sdk/languages"
)

func TestValidateFile(t *testing.T) {
	tests := []struct {
		name    string
		req     *transcribeRequest
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid request with file",
			req: &transcribeRequest{
				File: &strings.Reader{},
			},
			wantErr: false,
		},
		{
			name: "nil file should error",
			req: &transcribeRequest{
				File: nil,
			},
			wantErr: true,
			errMsg:  "file is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateFile(tt.req)
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

func TestValidateCodec(t *testing.T) {
	tests := []struct {
		name    string
		req     *transcribeRequest
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid with wav codec",
			req: &transcribeRequest{
				File: &strings.Reader{},

				AudioCodec: ptrAudioCodec(CodecWAV),
			},
			wantErr: false,
		},
		{
			name: "valid with mp3 codec",
			req: &transcribeRequest{
				File: &strings.Reader{},

				AudioCodec: ptrAudioCodec(CodecMP3),
			},
			wantErr: false,
		},
		{
			name: "valid with nil codec",
			req: &transcribeRequest{
				File: &strings.Reader{},
			},
			wantErr: false,
		},
		{
			name: "invalid codec",
			req: &transcribeRequest{
				File: &strings.Reader{},

				AudioCodec: ptrAudioCodec("invalid"),
			},
			wantErr: true,
			errMsg:  "unsupported audio codec",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateCodec(tt.req)
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

func TestValidateLanguage(t *testing.T) {
	tests := []struct {
		name    string
		req     *transcribeRequest
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid with Hindi language for saarika",
			req: &transcribeRequest{
				File: &strings.Reader{},

				Model:    ptrModel(ModelSaarika),
				Language: ptrLanguage(languages.Code("hi-IN")),
			},
			wantErr: false,
		},
		{
			name: "valid with English language for saarika",
			req: &transcribeRequest{
				File: &strings.Reader{},

				Model:    ptrModel(ModelSaarika),
				Language: ptrLanguage(languages.Code("en-IN")),
			},
			wantErr: false,
		},
		{
			name: "valid with nil language",
			req: &transcribeRequest{
				File: &strings.Reader{},
			},
			wantErr: false,
		},
		{
			name: "valid with Tamil for saaras",
			req: &transcribeRequest{
				File: &strings.Reader{},

				Model:    ptrModel(ModelSaaras),
				Language: ptrLanguage(languages.Code("ta-IN")),
			},
			wantErr: false,
		},
		{
			name: "invalid language for saarika",
			req: &transcribeRequest{
				File: &strings.Reader{},

				Model:    ptrModel(ModelSaarika),
				Language: ptrLanguage(languages.Code("fr-FR")),
			},
			wantErr: true,
			errMsg:  "not supported",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateLanguage(tt.req)
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

func TestValidateMode(t *testing.T) {
	tests := []struct {
		name    string
		model   *Model
		mode    *Mode
		wantErr bool
		errMsg  string
	}{
		{
			name:    "valid mode with saaras",
			model:   ptrModel(ModelSaaras),
			mode:    ptrMode(ModeTranscribe),
			wantErr: false,
		},
		{
			name:    "valid translate mode with saaras",
			model:   ptrModel(ModelSaaras),
			mode:    ptrMode(ModeTranslate),
			wantErr: false,
		},
		{
			name:    "nil mode is valid",
			model:   ptrModel(ModelSaaras),
			mode:    nil,
			wantErr: false,
		},
		{
			name:    "mode not supported with saarika",
			model:   ptrModel(ModelSaarika),
			mode:    ptrMode(ModeTranscribe),
			wantErr: true,
			errMsg:  "mode is only supported with saaras",
		},
		{
			name:    "nil model with mode should error",
			model:   nil,
			mode:    ptrMode(ModeTranscribe),
			wantErr: true,
			errMsg:  "mode is only supported with saaras",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateMode(tt.model, tt.mode)
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

func TestValidateStreamSampleRate(t *testing.T) {
	tests := []struct {
		name    string
		cfg     streamTranscribeRequest
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid sample rate 16000",
			cfg: streamTranscribeRequest{
				SampleRate: ptrStreamSampleRate(SampleRate16000),
			},
			wantErr: false,
		},
		{
			name: "valid sample rate 8000",
			cfg: streamTranscribeRequest{
				SampleRate: ptrStreamSampleRate(SampleRate8000),
			},
			wantErr: false,
		},
		{
			name: "zero sample rate is valid (uses default)",
			cfg: streamTranscribeRequest{
				SampleRate: nil,
			},
			wantErr: false,
		},
		{
			name: "invalid sample rate 44100",
			cfg: streamTranscribeRequest{
				SampleRate: ptrStreamSampleRate(44100),
			},
			wantErr: true,
			errMsg:  "invalid sample rate",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateStreamSampleRate(tt.cfg)
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

func TestValidateStreamCodec(t *testing.T) {
	tests := []struct {
		name    string
		cfg     streamTranscribeRequest
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid wav codec",
			cfg: streamTranscribeRequest{
				InputAudioCodec: ptrAudioCodec(CodecWAV),
			},
			wantErr: false,
		},
		{
			name: "valid pcm_s16le codec",
			cfg: streamTranscribeRequest{
				InputAudioCodec: ptrAudioCodec(CodecPCMS16LE),
			},
			wantErr: false,
		},
		{
			name: "nil codec is valid",
			cfg: streamTranscribeRequest{
				InputAudioCodec: nil,
			},
			wantErr: false,
		},
		{
			name: "invalid codec mp3",
			cfg: streamTranscribeRequest{
				InputAudioCodec: ptrAudioCodec(CodecMP3),
			},
			wantErr: true,
			errMsg:  "unsupported audio codec for streaming",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateStreamCodec(tt.cfg)
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

func TestValidateTranscribeRequest(t *testing.T) {
	tests := []struct {
		name    string
		req     *transcribeRequest
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid request with all options",
			req: &transcribeRequest{
				File: &strings.Reader{},

				Model:      ptrModel(ModelSaarika),
				Language:   ptrLanguage(languages.Code("hi-IN")),
				AudioCodec: ptrAudioCodec(CodecWAV),
			},
			wantErr: false,
		},
		{
			name: "valid request with saaras model",
			req: &transcribeRequest{
				File: &strings.Reader{},

				Model: ptrModel(ModelSaaras),
				Mode:  ptrMode(ModeTranscribe),
			},
			wantErr: false,
		},
		{
			name: "missing file should error",
			req: &transcribeRequest{
				File: nil,
			},
			wantErr: true,
			errMsg:  "file is required",
		},
		{
			name: "mode with saarika should error",
			req: &transcribeRequest{
				File: &strings.Reader{},

				Model: ptrModel(ModelSaarika),
				Mode:  ptrMode(ModeTranscribe),
			},
			wantErr: true,
			errMsg:  "mode is only supported with saaras",
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
		{
			name: "valid config with defaults",
			cfg: &streamTranscribeRequest{
				SampleRate: ptrStreamSampleRate(SampleRate16000),
			},
			wantErr: false,
		},
		{
			name: "valid config with language",
			cfg: &streamTranscribeRequest{
				Language:   languages.Code("en-IN"),
				SampleRate: ptrStreamSampleRate(SampleRate16000),
			},
			wantErr: false,
		},
		{
			name: "valid config with model and mode",
			cfg: &streamTranscribeRequest{
				Model:      ptrModel(ModelSaaras),
				Mode:       ptrMode(ModeTranslate),
				SampleRate: ptrStreamSampleRate(SampleRate16000),
			},
			wantErr: false,
		},
		{
			name: "invalid sample rate",
			cfg: &streamTranscribeRequest{
				SampleRate: ptrStreamSampleRate(44100),
			},
			wantErr: true,
			errMsg:  "invalid sample rate",
		},
		{
			name: "invalid codec for streaming",
			cfg: &streamTranscribeRequest{
				SampleRate:      ptrStreamSampleRate(SampleRate16000),
				InputAudioCodec: ptrAudioCodec(CodecMP3),
			},
			wantErr: true,
			errMsg:  "unsupported audio codec for streaming",
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

func ptrModel(m Model) *Model {
	return &m
}

func ptrMode(m Mode) *Mode {
	return &m
}

func ptrLanguage(c languages.Code) *languages.Code {
	return &c
}

func ptrAudioCodec(c InputAudioCodec) *InputAudioCodec {
	return &c
}

func ptrStreamSampleRate(r StreamSampleRate) *StreamSampleRate {
	return &r
}
