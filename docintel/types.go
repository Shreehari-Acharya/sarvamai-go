// Package docintel provides types for Document Intelligence API requests and responses.
package docintel

import (
	"time"

	"github.com/Shreehari-Acharya/sarvam-go-sdk/languages"
)

// Aliases for common types and constants to improve developer experience.

type LanguageCode = languages.Code

const (
	LanguageBnIN = languages.CodeBnIN
	LanguageEnIN = languages.CodeEnIN
	LanguageGuIN = languages.CodeGuIN
	LanguageHiIN = languages.CodeHiIN
	LanguageKnIN = languages.CodeKnIN
	LanguageMlIN = languages.CodeMlIN
	LanguageMrIN = languages.CodeMrIN
	LanguageOrIN = languages.CodeOrIN
	LanguagePaIN = languages.CodePaIN
	LanguageTaIN = languages.CodeTaIN
	LanguageTeIN = languages.CodeTeIN
)

// OutputFormat specifies the format for the extracted document content.
//
// Output is delivered as a ZIP file containing the processed documents.
//
//   - OutputFormatHTML: Structured HTML files with layout preservation
//   - OutputFormatMD: Markdown files (default)
//   - OutputFormatJSON: Structured JSON files for programmatic processing
type OutputFormat string

const (
	OutputFormatHTML OutputFormat = "html"
	OutputFormatMD   OutputFormat = "md"
	OutputFormatJSON OutputFormat = "json"
)

// ContainerType specifies the storage backend type for document processing.
type ContainerType string

const (
	ContainerTypeAzure   ContainerType = "Azure"
	ContainerTypeLocal   ContainerType = "Local"
	ContainerTypeGoogle  ContainerType = "Google"
	ContainerTypeAzureV1 ContainerType = "Azure_V1"
)

// JobState represents the current state of a document intelligence job.
//
//   - JobStateAccepted: Job created, awaiting file upload
//   - JobStatePending: File uploaded, waiting to start processing
//   - JobStateRunning: Processing in progress
//   - JobStateCompleted: All pages processed successfully
//   - JobStatePartiallyCompleted: Some pages succeeded, some failed
//   - JobStateFailed: All pages failed or job-level error
type JobState string

const (
	JobStateAccepted           JobState = "Accepted"
	JobStatePending            JobState = "Pending"
	JobStateRunning            JobState = "Running"
	JobStateCompleted          JobState = "Completed"
	JobStatePartiallyCompleted JobState = "PartiallyCompleted"
	JobStateFailed             JobState = "Failed"
)

// JobDetailState represents the processing state for an individual file within a job.
//
//   - JobDetailStatePending: File queued for processing
//   - JobDetailStateRunning: File being processed
//   - JobDetailStateSuccess: All pages processed successfully
//   - JobDetailStatePartialSuccess: Some pages succeeded, some failed
//   - JobDetailStateFailed: All pages failed
type JobDetailState string

const (
	JobDetailStatePending        JobDetailState = "Pending"
	JobDetailStateRunning        JobDetailState = "Running"
	JobDetailStateSuccess        JobDetailState = "Success"
	JobDetailStatePartialSuccess JobDetailState = "PartialSuccess"
	JobDetailStateFailed         JobDetailState = "Failed"
)

// JobParameters contains configuration options for a document intelligence job.
//
//   - Language: Primary language of the document in BCP-47 format (defaults to hi-IN)
//   - OutputFormat: Output format for extracted content (defaults to md)
type JobParameters struct {
	Language     *languages.Code `json:"language,omitempty"`
	OutputFormat *OutputFormat   `json:"output_format,omitempty"`
}

// Callback contains webhook configuration for job completion notifications.
//
//   - URL: HTTPS webhook URL to call upon job completion
//   - AuthToken: Optional authorization token sent as X-SARVAM-JOB-CALLBACK-TOKEN header
type Callback struct {
	URL       string  `json:"url"`
	AuthToken *string `json:"auth_token,omitempty"`
}

// DocIntelInitializeResponse represents the response from creating a new document intelligence job.
//
// # Fields
//
//   - JobID: Unique job identifier (UUID)
//   - StorageContainerType: Storage backend type
//   - JobParameters: Job configuration that was applied
//   - JobState: Current state of the job
type DocIntelInitializeResponse struct {
	JobID                string        `json:"job_id"`
	StorageContainerType ContainerType `json:"storage_container_type"`
	JobParameters        JobParameters `json:"job_parameters"`
	JobState             JobState      `json:"job_state"`
}

// FileSignedURLDetails contains details about a presigned URL for file upload/download.
//
// # Fields
//
//   - FileURL: Presigned URL for file access
//   - FileMetadata: Optional metadata about the file
type FileSignedURLDetails struct {
	FileURL      string                 `json:"file_url"`
	FileMetadata map[string]interface{} `json:"file_metadata,omitempty"`
}

// DocIntelGetUploadLinksResponse represents the response containing presigned upload URLs.
//
// # Fields
//
//   - JobID: Job identifier
//   - JobState: Current job state
//   - UploadUrls: Map of filename to presigned upload URL details
//   - StorageContainerType: Storage backend type
type DocIntelGetUploadLinksResponse struct {
	JobID                string                          `json:"job_id"`
	JobState             JobState                        `json:"job_state"`
	UploadUrls           map[string]FileSignedURLDetails `json:"upload_urls"`
	StorageContainerType ContainerType                   `json:"storage_container_type"`
}

// PageError contains error information for a specific page that failed processing.
//
// # Fields
//
//   - PageNumber: Page number that failed
//   - ErrorCode: Standardized error code
//   - ErrorMessage: Human-readable error description
type PageError struct {
	PageNumber   int    `json:"page_number"`
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

// TaskFileDetails contains information about an input or output file.
//
// # Fields
//
//   - FileName: Name of the file
//   - FileID: Unique identifier for the file
type TaskFileDetails struct {
	FileName string `json:"file_name"`
	FileID   string `json:"file_id"`
}

// JobDetail contains detailed processing information for a single file within a job.
//
// # Fields
//
//   - Inputs: Input file(s) for this task
//   - Outputs: Output file(s) produced
//   - State: Processing state for this file
//   - TotalPages: Total pages/images in the input file
//   - PagesProcessed: Number of pages processed so far
//   - PagesSucceeded: Number of pages successfully processed
//   - PagesFailed: Number of pages that failed processing
//   - ErrorMessage: Error message if processing failed
//   - ErrorCode: Standardized error code if failed
//   - PageErrors: Detailed errors for each failed page
type JobDetail struct {
	Inputs         []TaskFileDetails `json:"inputs"`
	Outputs        []TaskFileDetails `json:"outputs"`
	State          JobDetailState    `json:"state"`
	TotalPages     int               `json:"total_pages,omitempty"`
	PagesProcessed int               `json:"pages_processed,omitempty"`
	PagesSucceeded int               `json:"pages_succeeded,omitempty"`
	PagesFailed    int               `json:"pages_failed,omitempty"`
	ErrorMessage   string            `json:"error_message,omitempty"`
	ErrorCode      *string           `json:"error_code,omitempty"`
	PageErrors     []PageError       `json:"page_errors,omitempty"`
}

// DocIntelJobStatusResponse represents the detailed status of a document intelligence job.
//
// # Fields
//
//   - JobID: Job identifier (UUID)
//   - JobState: Current job state
//   - CreatedAt: Job creation timestamp (ISO 8601)
//   - UpdatedAt: Last update timestamp (ISO 8601)
//   - StorageContainerType: Storage backend type
//   - TotalFiles: Total input files (always 1)
//   - SuccessfulFilesCount: Files that completed successfully
//   - FailedFilesCount: Files that failed
//   - ErrorMessage: Job-level error message
//   - JobDetails: Per-file processing details with page metrics
type DocIntelJobStatusResponse struct {
	JobID                string        `json:"job_id"`
	JobState             JobState      `json:"job_state"`
	CreatedAt            time.Time     `json:"created_at"`
	UpdatedAt            time.Time     `json:"updated_at"`
	StorageContainerType ContainerType `json:"storage_container_type"`
	TotalFiles           int           `json:"total_files,omitempty"`
	SuccessfulFilesCount int           `json:"successful_files_count,omitempty"`
	FailedFilesCount     int           `json:"failed_files_count,omitempty"`
	ErrorMessage         string        `json:"error_message,omitempty"`
	JobDetails           []JobDetail   `json:"job_details,omitempty"`
}

// DocIntelGetDownloadLinksResponse represents the response containing presigned download URLs.
//
// # Fields
//
//   - JobID: Job identifier (UUID)
//   - JobState: Current job state
//   - StorageContainerType: Storage backend type
//   - DownloadUrls: Map of filename to presigned download URL details
//   - ErrorCode: Error code if retrieval failed
//   - ErrorMessage: Error message if retrieval failed
type DocIntelGetDownloadLinksResponse struct {
	JobID                string                          `json:"job_id"`
	JobState             JobState                        `json:"job_state"`
	StorageContainerType ContainerType                   `json:"storage_container_type"`
	DownloadURLs         map[string]FileSignedURLDetails `json:"download_urls"`
	ErrorCode            *string                         `json:"error_code,omitempty"`
	ErrorMessage         *string                         `json:"error_message,omitempty"`
}
