package sttjob

import (
	"strings"
	"testing"

	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/sarvamaierrors"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/languages"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/shared/speech"
)

func ptrBool(b bool) *bool {
	return &b
}

func ptrInt(i int) *int {
	return &i
}

func ptrLanguageCode(c languages.Code) *languages.Code {
	return &c
}

func ptrModel(m speech.Model) *speech.Model {
	return &m
}

func ptrMode(m speech.Mode) *speech.Mode {
	return &m
}

func TestValidateInitJobRequest(t *testing.T) {
	tests := []struct {
		name    string
		req     *initJobRequest
		wantErr bool
		errMsg  string
	}{
		{
			name:    "valid request with no options",
			req:     &initJobRequest{},
			wantErr: false,
		},
		{
			name: "valid request with saarika model",
			req: &initJobRequest{
				JobParameters: JobParameters{
					Model: ptrModel(speech.ModelSaarika),
				},
			},
			wantErr: false,
		},
		{
			name: "valid request with saaras model",
			req: &initJobRequest{
				JobParameters: JobParameters{
					Model: ptrModel(speech.ModelSaaras),
				},
			},
			wantErr: false,
		},
		{
			name: "valid request with saaras model and mode",
			req: &initJobRequest{
				JobParameters: JobParameters{
					Model: ptrModel(speech.ModelSaaras),
					Mode:  ptrMode(speech.ModeTranscribe),
				},
			},
			wantErr: false,
		},
		{
			name: "valid request with translate mode",
			req: &initJobRequest{
				JobParameters: JobParameters{
					Model: ptrModel(speech.ModelSaaras),
					Mode:  ptrMode(speech.ModeTranslate),
				},
			},
			wantErr: false,
		},
		{
			name: "valid request with all options",
			req: &initJobRequest{
				JobParameters: JobParameters{
					LanguageCode:    ptrLanguageCode(languages.Code("hi-IN")),
					Model:           ptrModel(speech.ModelSaaras),
					Mode:            ptrMode(speech.ModeTranscribe),
					WithTimestamps:  ptrBool(true),
					WithDiarization: ptrBool(true),
					NumSpeakers:     ptrInt(2),
				},
			},
			wantErr: false,
		},
		{
			name: "nil language (auto-detect) should pass",
			req: &initJobRequest{
				JobParameters: JobParameters{
					LanguageCode: nil,
				},
			},
			wantErr: false,
		},
		{
			name: "numSpeakers with diarization true should pass",
			req: &initJobRequest{
				JobParameters: JobParameters{
					WithDiarization: ptrBool(true),
					NumSpeakers:     ptrInt(2),
				},
			},
			wantErr: false,
		},
		{
			name: "numSpeakers with nil diarization should error",
			req: &initJobRequest{
				JobParameters: JobParameters{
					WithDiarization: nil,
					NumSpeakers:     ptrInt(2),
				},
			},
			wantErr: true,
			errMsg:  "num_speakers is only applicable when with_diarization is true",
		},
		{
			name: "numSpeakers with diarization false should error",
			req: &initJobRequest{
				JobParameters: JobParameters{
					WithDiarization: ptrBool(false),
					NumSpeakers:     ptrInt(2),
				},
			},
			wantErr: true,
			errMsg:  "num_speakers is only applicable when with_diarization is true",
		},
		{
			name: "mode with saarika model should error",
			req: &initJobRequest{
				JobParameters: JobParameters{
					Model: ptrModel(speech.ModelSaarika),
					Mode:  ptrMode(speech.ModeTranscribe),
				},
			},
			wantErr: true,
			errMsg:  "mode is only supported with saaras:v3 model",
		},
		{
			name: "language as-IN not supported by saarika should error",
			req: &initJobRequest{
				JobParameters: JobParameters{
					Model:        ptrModel(speech.ModelSaarika),
					LanguageCode: ptrLanguageCode(languages.Code("as-IN")),
				},
			},
			wantErr: true,
			errMsg:  "as-IN is not supported by saarika:v2.5 model",
		},
		{
			name: "language as-IN supported by saaras should pass",
			req: &initJobRequest{
				JobParameters: JobParameters{
					Model:        ptrModel(speech.ModelSaaras),
					LanguageCode: ptrLanguageCode(languages.Code("as-IN")),
				},
			},
			wantErr: false,
		},
		{
			name: "language hi-IN supported by saarika should pass",
			req: &initJobRequest{
				JobParameters: JobParameters{
					Model:        ptrModel(speech.ModelSaarika),
					LanguageCode: ptrLanguageCode(languages.Code("hi-IN")),
				},
			},
			wantErr: false,
		},
		{
			name: "language not in saaras set should error",
			req: &initJobRequest{
				JobParameters: JobParameters{
					LanguageCode: ptrLanguageCode(languages.Code("fr-FR")),
				},
			},
			wantErr: true,
			errMsg:  "fr-FR is not supported by",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateInitJobRequest(tt.req)
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

func TestValidateGetUploadLinksRequest(t *testing.T) {
	tests := []struct {
		name    string
		req     *getUploadLinksRequest
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid request with single file",
			req: &getUploadLinksRequest{
				JobID: "job-123",
				Files: []string{"audio.mp3"},
			},
			wantErr: false,
		},
		{
			name: "valid request with multiple files",
			req: &getUploadLinksRequest{
				JobID: "job-456",
				Files: []string{"audio1.mp3", "audio2.wav", "audio3.flac"},
			},
			wantErr: false,
		},
		{
			name: "empty jobID should error",
			req: &getUploadLinksRequest{
				JobID: "",
				Files: []string{"audio.mp3"},
			},
			wantErr: true,
			errMsg:  "job_id is required",
		},
		{
			name: "empty files slice should error",
			req: &getUploadLinksRequest{
				JobID: "job-123",
				Files: []string{},
			},
			wantErr: true,
			errMsg:  "at least one file is required",
		},
		{
			name: "valid request with empty string in files should pass (not validated)",
			req: &getUploadLinksRequest{
				JobID: "job-123",
				Files: []string{"audio.mp3", ""},
			},
			wantErr: false,
		},
		{
			name: "both jobID and files empty should return job_id error first",
			req: &getUploadLinksRequest{
				JobID: "",
				Files: []string{},
			},
			wantErr: true,
			errMsg:  "job_id is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateGetUploadLinksRequest(tt.req)
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

func TestValidateStartJobRequest(t *testing.T) {
	tests := []struct {
		name    string
		req     *startJobRequest
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid request with jobID",
			req: &startJobRequest{
				JobID: "job-123",
			},
			wantErr: false,
		},
		{
			name: "valid request with jobID and ptuID",
			req: &startJobRequest{
				JobID: "job-456",
				PtuID: ptrInt(12345),
			},
			wantErr: false,
		},
		{
			name: "empty jobID should error",
			req: &startJobRequest{
				JobID: "",
			},
			wantErr: true,
			errMsg:  "job_id is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateStartJobRequest(tt.req)
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

func TestValidateGetStatusRequest(t *testing.T) {
	tests := []struct {
		name    string
		jobID   string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "valid UUID jobID",
			jobID:   "550e8400-e29b-41d4-a716-446655440000",
			wantErr: false,
		},
		{
			name:    "valid alphanumeric jobID",
			jobID:   "abc123-def456",
			wantErr: false,
		},
		{
			name:    "empty jobID should error",
			jobID:   "",
			wantErr: true,
			errMsg:  "job_id is required",
		},
		{
			name:    "whitespace only jobID passes (not trimmed by validation)",
			jobID:   "   ",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateGetStatusRequest(tt.jobID)
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

func TestValidateGetDownloadLinksRequest(t *testing.T) {
	tests := []struct {
		name    string
		req     *getDownloadLinksRequest
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid request with single file",
			req: &getDownloadLinksRequest{
				JobID: "job-123",
				Files: []string{"file-id-1"},
			},
			wantErr: false,
		},
		{
			name: "valid request with multiple files",
			req: &getDownloadLinksRequest{
				JobID: "job-456",
				Files: []string{"file-id-1", "file-id-2", "file-id-3"},
			},
			wantErr: false,
		},
		{
			name: "empty jobID should error",
			req: &getDownloadLinksRequest{
				JobID: "",
				Files: []string{"file-id-1"},
			},
			wantErr: true,
			errMsg:  "job_id is required",
		},
		{
			name: "empty files slice should error",
			req: &getDownloadLinksRequest{
				JobID: "job-123",
				Files: []string{},
			},
			wantErr: true,
			errMsg:  "at least one file is required",
		},
		{
			name: "valid request with empty string in files should pass (not validated)",
			req: &getDownloadLinksRequest{
				JobID: "job-123",
				Files: []string{"file-id-1", ""},
			},
			wantErr: false,
		},
		{
			name: "both jobID and files empty should return job_id error first",
			req: &getDownloadLinksRequest{
				JobID: "",
				Files: []string{},
			},
			wantErr: true,
			errMsg:  "job_id is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateGetDownloadLinksRequest(tt.req)
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
			errMsg:  "invalid language code",
		},
		{
			name:    "invalid language zh-CN should error",
			lang:    languages.Code("zh-CN"),
			wantErr: true,
			errMsg:  "invalid language code",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &initJobRequest{}
			opt := WithLanguage(tt.lang)
			err := opt(req)

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
			model:   speech.Model("invalid-model"),
			wantErr: true,
			errMsg:  "invalid model",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &initJobRequest{}
			opt := WithModel(tt.model)
			err := opt(req)

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

func TestWithCallback(t *testing.T) {
	tests := []struct {
		name      string
		url       string
		authToken *string
		wantErr   bool
		errMsg    string
	}{
		{
			name:      "valid HTTPS URL",
			url:       "https://example.com/webhook",
			authToken: nil,
			wantErr:   false,
		},
		{
			name:      "valid URL with auth token",
			url:       "https://example.com/webhook",
			authToken: ptrString("secret-token"),
			wantErr:   false,
		},
		{
			name:      "empty URL should error",
			url:       "",
			authToken: nil,
			wantErr:   true,
			errMsg:    "callback URL cannot be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &initJobRequest{}
			opt := WithCallback(tt.url, tt.authToken)
			err := opt(req)

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

func TestWithPtuID(t *testing.T) {
	tests := []struct {
		name  string
		ptuID int
	}{
		{
			name:  "valid ptuID",
			ptuID: 12345,
		},
		{
			name:  "valid ptuID zero",
			ptuID: 0,
		},
		{
			name:  "valid large ptuID",
			ptuID: 999999999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &startJobRequest{}
			opt := WithPtuID(tt.ptuID)
			err := opt(req)

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if req.PtuID == nil || *req.PtuID != tt.ptuID {
				t.Errorf("expected ptuID %d, got %v", tt.ptuID, req.PtuID)
			}
		})
	}
}

func ptrString(s string) *string {
	return &s
}

func TestValidationErrorStruct(t *testing.T) {
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
