package docintel

import (
	"strings"
	"testing"

	"github.com/Shreehari-Acharya/sarvamai-go/languages"
)

func TestValidateCallbackURL(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "valid https URL",
			url:     "https://example.com/webhook",
			wantErr: false,
		},
		{
			name:    "valid https URL with path",
			url:     "https://example.com/webhook/path",
			wantErr: false,
		},
		{
			name:    "valid https URL with query params",
			url:     "https://example.com/webhook?token=abc",
			wantErr: false,
		},
		{
			name:    "http URL should error",
			url:     "http://example.com/webhook",
			wantErr: true,
			errMsg:  "HTTPS",
		},
		{
			name:    "invalid URL should error",
			url:     "not-a-url",
			wantErr: true,
			errMsg:  "valid URL",
		},
		{
			name:    "empty URL should error",
			url:     "",
			wantErr: true,
			errMsg:  "valid URL",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateCallbackURL(tt.url)
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
		lang    *languages.Code
		wantErr bool
		errMsg  string
	}{
		{
			name:    "valid Hindi",
			lang:    ptrLanguageCode(languages.Code("hi-IN")),
			wantErr: false,
		},
		{
			name:    "valid English",
			lang:    ptrLanguageCode(languages.Code("en-IN")),
			wantErr: false,
		},
		{
			name:    "valid Tamil",
			lang:    ptrLanguageCode(languages.Code("ta-IN")),
			wantErr: false,
		},
		{
			name:    "valid Bengali",
			lang:    ptrLanguageCode(languages.Code("bn-IN")),
			wantErr: false,
		},
		{
			name:    "valid Nepali",
			lang:    ptrLanguageCode(languages.Code("ne-IN")),
			wantErr: false,
		},
		{
			name:    "valid Sanskrit",
			lang:    ptrLanguageCode(languages.Code("sa-IN")),
			wantErr: false,
		},
		{
			name:    "nil language is valid",
			lang:    nil,
			wantErr: false,
		},
		{
			name:    "invalid language",
			lang:    ptrLanguageCode(languages.Code("fr-FR")),
			wantErr: true,
			errMsg:  "not supported",
		},
		{
			name:    "unsupported language code",
			lang:    ptrLanguageCode(languages.Code("zh-CN")),
			wantErr: true,
			errMsg:  "not supported",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateLanguage(tt.lang)
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

func TestValidateOutputFormat(t *testing.T) {
	tests := []struct {
		name    string
		format  *OutputFormat
		wantErr bool
		errMsg  string
	}{
		{
			name:    "valid HTML format",
			format:  ptrOutputFormat(OutputFormatHTML),
			wantErr: false,
		},
		{
			name:    "valid MD format",
			format:  ptrOutputFormat(OutputFormatMD),
			wantErr: false,
		},
		{
			name:    "valid JSON format",
			format:  ptrOutputFormat(OutputFormatJSON),
			wantErr: false,
		},
		{
			name:    "nil format is valid",
			format:  nil,
			wantErr: false,
		},
		{
			name:    "invalid format",
			format:  ptrOutputFormat(OutputFormat("invalid")),
			wantErr: true,
			errMsg:  "html, md, json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateOutputFormat(tt.format)
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

func TestValidateJobID(t *testing.T) {
	tests := []struct {
		name    string
		jobID   string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "valid UUID",
			jobID:   "550e8400-e29b-41d4-a716-446655440000",
			wantErr: false,
		},
		{
			name:    "valid job ID with letters and numbers",
			jobID:   "abc123-def456",
			wantErr: false,
		},
		{
			name:    "empty job ID should error",
			jobID:   "",
			wantErr: true,
			errMsg:  "job_id is required",
		},
		{
			name:    "whitespace only job ID should error",
			jobID:   "   ",
			wantErr: true,
			errMsg:  "job_id is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateJobID(tt.jobID)
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

func TestValidateFile(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		wantErr  bool
		errMsg   string
	}{
		{
			name:     "valid PDF file",
			filename: "document.pdf",
			wantErr:  false,
		},
		{
			name:     "valid ZIP file",
			filename: "images.zip",
			wantErr:  false,
		},
		{
			name:     "valid file with path",
			filename: "/path/to/document.pdf",
			wantErr:  false,
		},
		{
			name:     "valid file with spaces",
			filename: "my document.pdf",
			wantErr:  false,
		},
		{
			name:     "empty filename should error",
			filename: "",
			wantErr:  true,
			errMsg:   "filename is required",
		},
		{
			name:     "whitespace only should error",
			filename: "   ",
			wantErr:  true,
			errMsg:   "filename is required",
		},
		{
			name:     "invalid extension should error",
			filename: "document.txt",
			wantErr:  true,
			errMsg:   ".pdf or .zip",
		},
		{
			name:     "no extension should error",
			filename: "document",
			wantErr:  true,
			errMsg:   ".pdf or .zip",
		},
		{
			name:     "doc file should error",
			filename: "document.doc",
			wantErr:  true,
			errMsg:   ".pdf or .zip",
		},
		{
			name:     "png file should error",
			filename: "image.png",
			wantErr:  true,
			errMsg:   ".pdf or .zip",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateFile(tt.filename)
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

func TestValidateInitializeRequest(t *testing.T) {
	tests := []struct {
		name    string
		req     *docIntelInitializeRequest
		wantErr bool
		errMsg  string
	}{
		{
			name:    "valid request with no options",
			req:     &docIntelInitializeRequest{},
			wantErr: false,
		},
		{
			name: "valid request with language",
			req: &docIntelInitializeRequest{
				JobParameters: &JobParameters{
					Language: ptrLanguageCode(languages.Code("en-IN")),
				},
			},
			wantErr: false,
		},
		{
			name: "valid request with output format",
			req: &docIntelInitializeRequest{
				JobParameters: &JobParameters{
					OutputFormat: ptrOutputFormat(OutputFormatHTML),
				},
			},
			wantErr: false,
		},
		{
			name: "valid request with all options",
			req: &docIntelInitializeRequest{
				JobParameters: &JobParameters{
					Language:     ptrLanguageCode(languages.Code("hi-IN")),
					OutputFormat: ptrOutputFormat(OutputFormatJSON),
				},
			},
			wantErr: false,
		},
		{
			name: "invalid language should error",
			req: &docIntelInitializeRequest{
				JobParameters: &JobParameters{
					Language: ptrLanguageCode(languages.Code("fr-FR")),
				},
			},
			wantErr: true,
			errMsg:  "not supported",
		},
		{
			name: "invalid output format should error",
			req: &docIntelInitializeRequest{
				JobParameters: &JobParameters{
					OutputFormat: ptrOutputFormat(OutputFormat("xml")),
				},
			},
			wantErr: true,
			errMsg:  "html, md, json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateInitializeRequest(tt.req)
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
		name     string
		jobID    string
		filename string
		wantErr  bool
		errMsg   string
	}{
		{
			name:     "valid request",
			jobID:    "job-123",
			filename: "document.pdf",
			wantErr:  false,
		},
		{
			name:     "valid request with ZIP",
			jobID:    "job-456",
			filename: "images.zip",
			wantErr:  false,
		},
		{
			name:     "empty job ID should error",
			jobID:    "",
			filename: "document.pdf",
			wantErr:  true,
			errMsg:   "job_id is required",
		},
		{
			name:     "empty filename should error",
			jobID:    "job-123",
			filename: "",
			wantErr:  true,
			errMsg:   "filename is required",
		},
		{
			name:     "invalid filename extension should error",
			jobID:    "job-123",
			filename: "document.txt",
			wantErr:  true,
			errMsg:   ".pdf or .zip",
		},
		{
			name:     "both invalid should return job_id error first",
			jobID:    "",
			filename: "",
			wantErr:  true,
			errMsg:   "job_id is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateGetUploadLinksRequest(tt.jobID, tt.filename)
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

func ptrLanguageCode(c languages.Code) *languages.Code {
	return &c
}

func ptrOutputFormat(f OutputFormat) *OutputFormat {
	return &f
}
