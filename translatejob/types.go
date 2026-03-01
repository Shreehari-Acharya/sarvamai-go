// Package translatejob provides types for the Speech-to-Text Translate Batch Job API.
//
// The batch job API allows you to process multiple audio files with translation.
// The workflow consists of:
//   - [Initialize] - Create a new bulk job and get job ID and storage details
//   - [GetUploadLinks] - Get presigned URLs to upload audio files
//   - [Start] - Start processing the uploaded files
//   - [GetStatus] - Poll for job status and results
//   - [GetDownloadLinks] - Get presigned URLs to download translated outputs
package translatejob

import (
	"time"

	"github.com/Shreehari-Acharya/sarvamai-go/shared/speech"
)

// Aliases for common types and constants to improve developer experience.

type Model = speech.Model

const (
	ModelSaarasV25 = speech.ModelSaarasV25
	ModelSaaras    = speech.ModelSaaras
)

type JobState = speech.JobState

const (
	JobStateAccepted  = speech.JobStateAccepted
	JobStatePending   = speech.JobStatePending
	JobStateRunning   = speech.JobStateRunning
	JobStateCompleted = speech.JobStateCompleted
	JobStateFailed    = speech.JobStateFailed
)

type ContainerType = speech.ContainerType

const (
	ContainerTypeAzure   = speech.ContainerTypeAzure
	ContainerTypeLocal   = speech.ContainerTypeLocal
	ContainerTypeGCS     = speech.ContainerTypeGCS
	ContainerTypeAzureV1 = speech.ContainerTypeAzureV1
)

// JobParameters defines the configuration for a translation job.
//
// Use functional options from [options] to set these values.
// The API currently only supports "saaras:v2.5" model for translation.
type JobParameters struct {
	// Prompt is an optional prompt to assist the transcription.
	// Can be used to provide context or improve accuracy for specific use cases.
	Prompt *string `json:"prompt,omitempty"`

	// Model is the speech-to-text translation model to use.
	// Currently only "saaras:v2.5" (ModelSaarasV25) is supported for translation.
	// If not specified, defaults to "saaras:v2.5".
	Model *speech.Model `json:"model,omitempty"`

	// WithDiarization enables speaker diarization, which identifies and separates
	// different speakers in the audio. When true, the response will include
	// speaker-specific segments. Currently in Beta.
	WithDiarization *bool `json:"with_diarization,omitempty"`

	// NumSpeakers is the expected number of speakers in the audio.
	// Used when with_diarization is true. If not specified, the API will attempt
	// to detect the number of speakers automatically.
	NumSpeakers *int `json:"num_speakers,omitempty"`
}

// InitJobResponse is the response from initializing a new bulk translation job.
type InitJobResponse struct {
	// JobID is the unique identifier for the bulk job.
	JobID string `json:"job_id"`

	// StorageContainerType indicates where the files are stored.
	// Values: Azure, Local, Google, Azure_V1.
	StorageContainerType speech.ContainerType `json:"storage_container_type"`

	// JobParameters are the parameters that were used for this job.
	JobParameters JobParameters `json:"job_parameters"`

	// JobState is the current state of the job.
	// Values: Accepted, Pending, Running, Completed, Failed.
	JobState speech.JobState `json:"job_state"`
}

// UploadUrl contains the presigned URL details for uploading an audio file.
type UploadUrl struct {
	// FileUrl is the presigned URL where the file should be uploaded.
	FileUrl string `json:"file_url"`

	// FileMetadata contains any additional metadata about the file.
	FileMetadata map[string]any `json:"file_metadata"`
}

// GetUploadLinksResponse is the response containing presigned upload URLs.
type GetUploadLinksResponse struct {
	// JobID is the unique identifier for the bulk job.
	JobID string `json:"job_id"`

	// JobState is the current state of the job.
	JobState speech.JobState `json:"job_state"`

	// UploadUrls is a map of filenames to their presigned upload URLs.
	// Key: original filename, Value: presigned URL details.
	UploadUrls map[string]UploadUrl `json:"upload_urls"`

	// StorageContainerType indicates where the files are stored.
	StorageContainerType speech.ContainerType `json:"storage_container_type"`
}

// JobStatusResponse is the response containing the current status of a bulk job.
type JobStatusResponse struct {
	// JobID is the unique identifier for the bulk job.
	JobID string `json:"job_id"`

	// JobState is the current state of the job.
	// Values: Accepted, Pending, Running, Completed, Failed.
	JobState speech.JobState `json:"job_state"`

	// CreatedAt is the timestamp when the job was created.
	CreatedAt time.Time `json:"created_at"`

	// UpdatedAt is the timestamp when the job was last updated.
	UpdatedAt time.Time `json:"updated_at"`

	// StorageContainerType indicates where the files are stored.
	StorageContainerType speech.ContainerType `json:"storage_container_type"`

	// TotalFiles is the total number of files in the job.
	TotalFiles *int `json:"total_files,omitempty"`

	// SuccessfulFilesCount is the number of files successfully processed.
	SuccessfulFilesCount *int `json:"successful_files_count,omitempty"`

	// FailedFilesCount is the number of files that failed processing.
	FailedFilesCount *int `json:"failed_files_count,omitempty"`

	// ErrorMessage contains any error message if the job failed.
	ErrorMessage *string `json:"error_message,omitempty"`

	// JobDetails contains file-level details for each processed file.
	JobDetails []speech.JobDetail `json:"job_details,omitempty"`
}

// DownloadUrl contains the presigned URL details for downloading a result file.
// Aliases [UploadUrl] since they have the same structure.
type DownloadUrl = UploadUrl

// GetDownloadLinksResponse is the response containing presigned download URLs.
type GetDownloadLinksResponse struct {
	// JobID is the unique identifier for the bulk job.
	JobID string `json:"job_id"`

	// JobState is the current state of the job.
	JobState speech.JobState `json:"job_state"`

	// DownloadUrls is a map of filenames to their presigned download URLs.
	// Key: original filename, Value: presigned URL details.
	DownloadUrls map[string]DownloadUrl `json:"download_urls,omitempty"`

	// StorageContainerType indicates where the files are stored.
	StorageContainerType speech.ContainerType `json:"storage_container_type"`
}
