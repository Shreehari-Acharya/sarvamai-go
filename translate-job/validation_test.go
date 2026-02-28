package translatejob

import (
	"strings"
	"testing"
)

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
			name: "empty job_id should error",
			req: &getUploadLinksRequest{
				JobID: "",
				Files: []string{"audio.mp3"},
			},
			wantErr: true,
			errMsg:  "job_id is required",
		},
		{
			name: "empty files should error",
			req: &getUploadLinksRequest{
				JobID: "job-123",
				Files: []string{},
			},
			wantErr: true,
			errMsg:  "at least one file is required",
		},
		{
			name: "both empty should return job_id error first",
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
			name: "valid request",
			req: &startJobRequest{
				JobID: "550e8400-e29b-41d4-a716-446655440000",
			},
			wantErr: false,
		},
		{
			name: "valid job ID with letters and numbers",
			req: &startJobRequest{
				JobID: "abc123-def456",
			},
			wantErr: false,
		},
		{
			name: "empty job_id should error",
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
				Files: []string{"output.json"},
			},
			wantErr: false,
		},
		{
			name: "valid request with multiple files",
			req: &getDownloadLinksRequest{
				JobID: "job-456",
				Files: []string{"output1.json", "output2.json"},
			},
			wantErr: false,
		},
		{
			name: "empty job_id should error",
			req: &getDownloadLinksRequest{
				JobID: "",
				Files: []string{"output.json"},
			},
			wantErr: true,
			errMsg:  "job_id is required",
		},
		{
			name: "empty files should error",
			req: &getDownloadLinksRequest{
				JobID: "job-123",
				Files: []string{},
			},
			wantErr: true,
			errMsg:  "at least one file is required",
		},
		{
			name: "both empty should return job_id error first",
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
