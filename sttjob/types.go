package sttjob

import (
	"time"

	"github.com/Shreehari-Acharya/sarvam-go-sdk/languages"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/shared/speech"
)

// Aliases for common types and constants to improve developer experience.

type Model = speech.Model

const (
	ModelSaarika = speech.ModelSaarika
	ModelSaaras  = speech.ModelSaaras
)

type Mode = speech.Mode

const (
	ModeTranscribe = speech.ModeTranscribe
	ModeTranslate  = speech.ModeTranslate
	ModeVerbatim   = speech.ModeVerbatim
	ModeTranslit   = speech.ModeTranslit
	ModeCodemix    = speech.ModeCodemix
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

type FileState = speech.FileState

const (
	FileStateSuccess             = speech.FileStateSuccess
	FileStateAPIError            = speech.FileStateAPIError
	FileStateInternalServerError = speech.FileStateInternalServerError
)

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

// JobParameters defines the configuration for a speech-to-text bulk job.
//
// This struct is used when initializing a new job to specify the language,
// model, and processing options for transcribing audio files.
//
// # Language Selection
//
// The language_code field specifies the language of the input audio in BCP-47 format.
// Use "unknown" (default) for auto-detection of the language.
//
// # Model Selection
//
//   - saarika:v2.5 (default): Best for multilingual audio, supports 12 languages
//   - saaras:v3: State-of-the-art model supporting 23 languages with multiple modes
//
// # Mode (saaras:v3 only)
//
// When using saaras:v3 model, you can specify the processing mode:
//
//   - transcribe: Standard transcription with proper formatting (default)
//   - translate: Translates speech from Indic languages to English
//   - verbatim: Exact word-for-word transcription preserving filler words
//   - translit: Romanization - transliterates to Latin script
//   - codemix: Code-mixed text with English words in English script
//
// # Diarization
//
// Speaker diarization identifies and separates different speakers in the audio.
// This is currently in beta mode and requires with_diarization to be enabled.
// When enabled, num_speakers can optionally specify the expected number of speakers.
//
// # Example
//
//	params := sttjob.JobParameters{
//	    LanguageCode:     languages.ToLanguageCode("hi-IN"),
//	    Model:            speech.ModelSaaras,
//	    Mode:             speech.ModeTranscribe,
//	    WithTimestamps:  ptr(true),
//	    WithDiarization:  ptr(true),
//	    NumSpeakers:      ptr(2),
//	}
type JobParameters struct {
	// LanguageCode specifies the language of the input audio in BCP-47 format.
	//
	// Supported languages for saarika:v2.5 (12 languages):
	// unknown (auto-detect), hi-IN, bn-IN, kn-IN, ml-IN, mr-IN, or-IN, pa-IN, ta-IN, te-IN, en-IN, gu-IN
	//
	// Additional languages for saaras:v3 (23 languages):
	// as-IN, ur-IN, ne-IN, kok-IN, ks-IN, sd-IN, sa-IN, sat-IN, mni-IN, brx-IN, mai-IN, doi-IN
	//
	// Default: "unknown" (auto-detect)
	LanguageCode *languages.Code `json:"language_code,omitempty"`

	// Model specifies the speech recognition model to use.
	//
	//   - saarika:v2.5 (default): Multilingual model, best for audio with multiple languages
	//   - saaras:v3: Advanced model with flexible output modes, supports 23 languages
	//
	// Default: saarika:v2.5
	Model *speech.Model `json:"model,omitempty"`

	// Mode specifies the processing mode for speech recognition.
	//
	// This field is only applicable when using saaras:v3 model.
	// Different modes provide different output formats:
	//
	//   - transcribe: Standard transcription in original language with proper formatting
	//   - translate: Translates speech from any supported Indic language to English
	//   - verbatim: Exact word-for-word transcription without normalization
	//   - translit: Romanization - converts to Latin script only
	//   - codemix: Mixed script with English words in English, Indic words in native script
	//
	// Example audio: "मेरा फोन नंबर है 9840950950"
	//   - transcribe: "मेरा फोन नंबर है 9840950950"
	//   - translate: "My phone number is 9840950950"
	//   - verbatim: "मेरा फोन नंबर है नौ आठ चार zero नौ पांच zero नौ पांच zero"
	//   - translit: "mera phone number hai 9840950950"
	//   - codemix: "मेरा phone number है 9840950950"
	//
	// Default: transcribe
	Mode *speech.Mode `json:"mode,omitempty"`

	// WithTimestamps enables word-level timestamps in the transcription response.
	//
	// When enabled, the response will include start and end times for each word.
	// Useful for applications that need precise timing information.
	//
	// Default: false
	WithTimestamps *bool `json:"with_timestamps,omitempty"`

	// WithDiarization enables speaker diarization.
	//
	// Speaker diarization identifies and separates different speakers in the audio.
	// This is useful for multi-speaker audio transcription.
	//
	// Note: This feature is currently in beta mode.
	//
	// Default: false
	WithDiarization *bool `json:"with_diarization,omitempty"`

	// NumSpeakers specifies the expected number of speakers in the audio.
	//
	// This field is used when with_diarization is enabled to help the model
	// identify and separate speakers more accurately.
	//
	// Note: This feature is currently in beta mode and requires with_diarization to be true.
	//
	// Default: nil (auto-detect number of speakers)
	NumSpeakers *int `json:"num_speakers,omitempty"`
}

// InitJobTranscribeResponse contains the response from initializing a new speech-to-text bulk job.
//
// This response includes the job ID, storage container details, and initial job status.
// Use the job_id to upload files, start the job, and check its status.
type InitJobTranscribeResponse struct {
	// JobID is the unique identifier for the bulk transcription job.
	// Use this ID to upload files, start the job, and track progress.
	JobID string `json:"job_id"`

	// StorageContainerType specifies the type of storage container used for the job.
	// This determines where the uploaded files will be stored and processed.
	// Possible values: Azure, Local, Google, Azure_V1
	StorageContainerType speech.ContainerType `json:"storage_container_type"`

	// JobParameters contains the configuration used for this job.
	// This is an empty object in the current API version.
	JobParameters struct{} `json:"job_parameters"`

	// JobState represents the current state of the job.
	// Possible values: Accepted, Pending, Running, Completed, Failed
	JobState speech.JobState `json:"job_state"`
}

// UploadUrl contains the presigned URL and metadata for uploading an audio file.
//
// The file_url is a pre-authenticated URL that can be used to upload
// the audio file directly to cloud storage.
type UploadUrl struct {
	// FileUrl is the presigned URL for uploading the audio file.
	// Make a PUT request to this URL with the audio file as the body.
	FileUrl string `json:"file_url"`

	// FileMetadata contains additional metadata about the uploaded file.
	// This may include information like file size, format, etc.
	FileMetadata map[string]any `json:"file_metadata"`
}

// GetUploadLinksResponse contains the response for generating upload URLs.
//
// This response provides presigned URLs for uploading audio files to the job's
// storage container. Each file in the request will have a corresponding
// upload URL in the response.
type GetUploadLinksResponse struct {
	// JobID is the unique identifier for the bulk transcription job.
	JobId string `json:"job_id"`

	// JobState represents the current state of the job.
	// Possible values: Accepted, Pending, Running, Completed, Failed
	JobState speech.JobState `json:"job_state"`

	// UploadUrls is a map of file names to their presigned upload URLs.
	// The key is the original filename, and the value contains the URL
	// to use for uploading that file.
	UploadUrls map[string]UploadUrl `json:"upload_urls"`

	// StorageContainerType specifies the type of storage container used.
	StorageContainerType speech.ContainerType `json:"storage_container_type"`
}

// JobTranscribeResponse contains the status and details of a speech-to-text bulk job.
//
// This response is returned by Start() and GetStatus() methods and includes
// overall job progress as well as per-file details.
type JobTranscribeResponse struct {
	// JobID is the unique identifier for the bulk transcription job.
	JobId string `json:"job_id"`

	// JobState represents the current state of the job.
	//   - Accepted: Job has been accepted but not yet started
	//   - Pending: Job is queued and waiting to be processed
	//   - Running: Job is currently being processed
	//   - Completed: All files have been processed successfully
	//   - Failed: Job failed due to an error
	JobState speech.JobState `json:"job_state"`

	// CreatedAt is the timestamp when the job was created.
	CreatedAt time.Time `json:"created_at"`

	// UpdatedAt is the timestamp when the job was last updated.
	UpdatedAt time.Time `json:"updated_at"`

	// StorageContainerType specifies the type of storage container used.
	StorageContainerType speech.ContainerType `json:"storage_container_type"`

	// TotalFiles is the total number of files in the job.
	// This field is nil until files are uploaded.
	TotalFiles *int `json:"total_files,omitempty"`

	// SuccessfulFilesCount is the number of files that have been
	// successfully transcribed.
	SuccessfulFilesCount *int `json:"successful_files_count,omitempty"`

	// FailedFilesCount is the number of files that failed to transcribe.
	FailedFilesCount *int `json:"failed_files_count,omitempty"`

	// ErrorMessage contains an error message if the job failed.
	// This field is nil if the job is still running or completed successfully.
	ErrorMessage *string `json:"error_message,omitempty"`

	// JobDetails contains detailed information about each file in the job.
	// This includes input/output file information and per-file status.
	JobDetails *[]speech.JobDetail `json:"job_details,omitempty"`
}

// DownloadUrl is a type alias for UploadUrl, representing a presigned URL
// for downloading a transcribed file.
//
// The file_url can be used to download the transcription output directly
// from cloud storage.
type DownloadUrl = UploadUrl

// GetDownloadLinksResponse contains the response for generating download URLs.
//
// This response provides presigned URLs for downloading the transcribed
// output files from a completed job.
type GetDownloadLinksResponse struct {
	// JobID is the unique identifier for the bulk transcription job.
	JobId string `json:"job_id"`

	// JobState represents the current state of the job.
	// Only returns download links when job is Completed.
	JobState speech.JobState `json:"job_state"`

	// DownloadUrls is a map of file IDs to their presigned download URLs.
	// The key is the file identifier, and the value contains the URL
	// to use for downloading that file's transcription.
	DownloadUrls map[string]DownloadUrl `json:"download_urls,omitempty"`

	// StorageContainerType specifies the type of storage container used.
	StorageContainerType speech.ContainerType `json:"storage_container_type"`
}
