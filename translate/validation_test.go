package translate

import (
	"strings"
	"testing"

	"github.com/Shreehari-Acharya/sarvam-go-sdk/shared/speech"
)

func TestValidateTranslateRequest(t *testing.T) {
	tests := []struct {
		name    string
		req     *translateRequest
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid request with file",
			req: &translateRequest{
				File: strings.NewReader("audio data"),
			},
			wantErr: false,
		},
		{
			name: "nil file should error",
			req: &translateRequest{
				File: nil,
			},
			wantErr: true,
			errMsg:  "file is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateTranslateRequest(tt.req)
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

func TestWithModelForTranslateStream(t *testing.T) {
	tests := []struct {
		name    string
		model   speech.Model
		wantErr bool
		errMsg  string
	}{
		{
			name:    "valid model saaras v2.5",
			model:   speech.ModelSaarasV25,
			wantErr: false,
		},
		{
			name:    "valid model saaras v3",
			model:   speech.ModelSaaras,
			wantErr: false,
		},
		{
			name:    "invalid model saarika",
			model:   speech.ModelSaarika,
			wantErr: true,
			errMsg:  "saaras:v2.5, saaras:v3",
		},
		{
			name:    "invalid custom model",
			model:   speech.Model("custom-model"),
			wantErr: true,
			errMsg:  "saaras:v2.5, saaras:v3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &streamTranslateRequest{}
			opt := WithModelForTranslateStream(tt.model)
			opt(req)
			err := validateTranslateStreamRequest(req)

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
					t.Errorf("expected model to be set to %v, got %v", tt.model, req.Model)
				}
			}
		})
	}
}

func TestWithModeForTranslateStream(t *testing.T) {
	tests := []struct {
		name    string
		mode    speech.Mode
		wantErr bool
		errMsg  string
	}{
		{
			name:    "valid mode translate",
			mode:    speech.ModeTranslate,
			wantErr: false,
		},
		{
			name:    "valid mode transcribe",
			mode:    speech.ModeTranscribe,
			wantErr: false,
		},
		{
			name:    "valid mode verbatim",
			mode:    speech.ModeVerbatim,
			wantErr: false,
		},
		{
			name:    "valid mode translit",
			mode:    speech.ModeTranslit,
			wantErr: false,
		},
		{
			name:    "valid mode codemix",
			mode:    speech.ModeCodemix,
			wantErr: false,
		},
		{
			name:    "invalid mode",
			mode:    speech.Mode("invalid"),
			wantErr: true,
			errMsg:  "transcribe, translate, verbatim, translit, codemix",
		},
		{
			name:    "empty mode",
			mode:    speech.Mode(""),
			wantErr: true,
			errMsg:  "transcribe, translate, verbatim, translit, codemix",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &streamTranslateRequest{}
			opt := WithModeForTranslateStream(tt.mode)
			opt(req)
			err := validateTranslateStreamRequest(req)

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
				if req.Mode == nil || *req.Mode != tt.mode {
					t.Errorf("expected mode to be set to %v, got %v", tt.mode, req.Mode)
				}
			}
		})
	}
}

func TestWithAudioCodecForTranslateStream(t *testing.T) {
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
			name:    "invalid codec mp3",
			codec:   speech.CodecMP3,
			wantErr: true,
			errMsg:  "wav, pcm_s16le, pcm_l16, pcm_raw",
		},
		{
			name:    "invalid codec aac",
			codec:   speech.CodecAAC,
			wantErr: true,
			errMsg:  "wav, pcm_s16le, pcm_l16, pcm_raw",
		},
		{
			name:    "invalid codec flac",
			codec:   speech.CodecFLAC,
			wantErr: true,
			errMsg:  "wav, pcm_s16le, pcm_l16, pcm_raw",
		},
		{
			name:    "invalid codec opus",
			codec:   speech.CodecOPUS,
			wantErr: true,
			errMsg:  "wav, pcm_s16le, pcm_l16, pcm_raw",
		},
		{
			name:    "invalid codec ogg",
			codec:   speech.CodecOGG,
			wantErr: true,
			errMsg:  "wav, pcm_s16le, pcm_l16, pcm_raw",
		},
		{
			name:    "invalid codec mp4",
			codec:   speech.CodecMP4,
			wantErr: true,
			errMsg:  "wav, pcm_s16le, pcm_l16, pcm_raw",
		},
		{
			name:    "invalid empty codec",
			codec:   speech.InputAudioCodec(""),
			wantErr: true,
			errMsg:  "wav, pcm_s16le, pcm_l16, pcm_raw",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &streamTranslateRequest{}
			opt := WithAudioCodecForTranslateStream(tt.codec)
			opt(req)
			err := validateTranslateStreamRequest(req)

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
				if req.InputAudioCodec == nil || *req.InputAudioCodec != tt.codec {
					t.Errorf("expected codec to be set to %v, got %v", tt.codec, req.InputAudioCodec)
				}
			}
		})
	}
}

func TestWithSampleRateForTranslateStream(t *testing.T) {
	tests := []struct {
		name    string
		rate    speech.StreamSampleRate
		wantErr bool
		errMsg  string
	}{
		{
			name:    "valid rate 8000",
			rate:    speech.SampleRate8000,
			wantErr: false,
		},
		{
			name:    "valid rate 16000",
			rate:    speech.SampleRate16000,
			wantErr: false,
		},
		{
			name:    "invalid rate 22050",
			rate:    speech.SampleRate22050,
			wantErr: true,
			errMsg:  "8000, 16000 Hz",
		},
		{
			name:    "invalid rate 24000",
			rate:    speech.SampleRate24000,
			wantErr: true,
			errMsg:  "8000, 16000 Hz",
		},
		{
			name:    "invalid zero rate",
			rate:    speech.StreamSampleRate(0),
			wantErr: true,
			errMsg:  "8000, 16000 Hz",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &streamTranslateRequest{}
			opt := WithSampleRateForTranslateStream(tt.rate)
			opt(req)
			err := validateTranslateStreamRequest(req)

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
				if req.SampleRate == nil || *req.SampleRate != tt.rate {
					t.Errorf("expected sample rate to be set to %v, got %v", tt.rate, req.SampleRate)
				}
			}
		})
	}
}
