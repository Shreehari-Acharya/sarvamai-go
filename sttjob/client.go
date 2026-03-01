package sttjob

import (
	"context"
	"net/url"
	"strconv"

	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/transport"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/shared/speech"
)

// SttJobClient provides access to the Speech-to-Text Bulk Job API.
//
// This client enables batch processing of multiple audio files for transcription.
// The workflow consists of:
//
//  1. Initialize a new job with Initialize()
//  2. Upload audio files using GetUploadLinks() and uploading to the provided URLs
//  3. Start processing with Start()
//  4. Poll for status using GetStatus()
//  5. Download results using GetDownloadLinks()
//
// # Rate Limiting
//
// When polling for job status, it's recommended to implement a minimum delay
// of 5 milliseconds between consecutive requests to prevent rate limiting errors.
type SttJobClient struct {
	transport *transport.Transport
}

// NewSttJobClient creates a new Speech-to-Text Bulk Job client.
//
// # Parameters
//
//	t: Transport instance configured with API key and base URL
//
// # Returns
//
//	A new SttJobClient instance for making bulk transcription requests
func NewSttJobClient(t *transport.Transport) *SttJobClient {
	return &SttJobClient{
		transport: t,
	}
}

// initJobRequest is the internal request structure for initializing a bulk job.
type initJobRequest struct {
	JobParameters JobParameters    `json:"job_parameters"`
	Callback      *speech.Callback `json:"callback,omitempty"`
}

// Initialize creates a new speech-to-text bulk job.
//
// This is the first step in the bulk transcription workflow. It creates a new
// job and returns the job ID and storage container details. After initialization,
// you can upload audio files to be transcribed.
//
// # Parameters
//
//	ctx: Context for the request (used for cancellation and timeouts)
//	options: Optional functional options to configure the job
//
// # Returns
//
//	InitJobTranscribeResponse containing the job ID and storage details, or an error
//
// # Functional Options
//
//	WithLanguage(languages.Code)     - Language code for transcription (default: auto-detect)
//	WithModel(speech.Model)          - Model to use (saarika:v2.5 or saaras:v3)
//	WithMode(speech.Mode)            - Processing mode (transcribe, translate, verbatim, translit, codemix)
//	WithTimeStamps(bool)             - Include word-level timestamps in response
//	WithDiarization(bool)            - Enable speaker diarization (beta)
//	WithNumSpeakers(int)             - Expected number of speakers (requires diarization)
//	WithCallback(string, *string)    - Webhook URL and optional auth token for job completion
//
// # Example
//
//	resp, err := client.SpeechToTextJob.Initialize(
//	    context.Background(),
//	    sttjob.WithLanguage(languages.ToLanguageCode("hi-IN")),
//	    sttjob.WithModel(speech.ModelSaaras),
//	    sttjob.WithMode(speech.ModeTranscribe),
//	    sttjob.WithTimeStamps(true),
//	    sttjob.WithDiarization(true),
//	    sttjob.WithNumSpeakers(2),
//	    sttjob.WithCallback("https://your-webhook.com/callback", nil),
//	)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Job ID: %s\n", resp.JobID)
//	fmt.Printf("Storage Type: %s\n", resp.StorageContainerType)
//
// # API Reference
//
// https://docs.sarvam.ai/api-reference-docs/speech-to-text/stt/job/initiate
func (c *SttJobClient) Initialize(ctx context.Context,
	options ...InitJobOption,
) (*InitJobTranscribeResponse, error) {
	var resp InitJobTranscribeResponse

	req := &initJobRequest{}

	for _, option := range options {
		option(req)
	}

	if err := validateInitJobRequest(req); err != nil {
		return nil, err
	}

	err := c.transport.DoRequest(
		ctx,
		"POST",
		"/speech-to-text/job/v1",
		req,
		&resp,
		"application/json",
	)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// getUploadLinksRequest is the internal request structure for getting upload links.
type getUploadLinksRequest struct {
	JobID string   `json:"job_id"`
	Files []string `json:"files"`
}

// GetUploadLinks generates presigned URLs for uploading audio files to a bulk job.
//
// After initializing a job, use this method to get presigned URLs for each audio
// file you want to transcribe. These URLs allow direct upload to cloud storage
// without requiring additional authentication.
//
// # Parameters
//
//	ctx: Context for the request (used for cancellation and timeouts)
//	jobID: The unique identifier returned from Initialize()
//	files: List of file names to upload (e.g., ["audio1.mp3", "audio2.wav"])
//
// # Returns
//
//	GetUploadLinksResponse containing presigned upload URLs, or an error
//
// # Example
//
//	resp, err := client.SpeechToTextJob.GetUploadLinks(
//	    context.Background(),
//	    "job-id-from-initialize",
//	    []string{"audio1.mp3", "audio2.wav", "audio3.flac"},
//	)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Upload each file to its presigned URL
//	for filename, uploadInfo := range resp.UploadUrls {
//	    fmt.Printf("Uploading %s to %s\n", filename, uploadInfo.FileUrl)
//	    // Make PUT request with audio file content
//	}
//
// # API Reference
//
// https://docs.sarvam.ai/api-reference-docs/speech-to-text/stt/job/upload
func (c *SttJobClient) GetUploadLinks(ctx context.Context,
	jobID string,
	files []string,
) (*GetUploadLinksResponse, error) {

	var resp GetUploadLinksResponse

	req := &getUploadLinksRequest{
		JobID: jobID,
		Files: files,
	}

	err := validateGetUploadLinksRequest(req)
	if err != nil {
		return nil, err
	}

	err = c.transport.DoRequest(
		ctx,
		"POST",
		"/speech-to-text/job/v1/upload-files",
		req,
		&resp,
		"application/json",
	)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// startJobRequest is the internal request structure for starting a bulk job.
type startJobRequest struct {
	JobID string `json:"job_id"`
	PtuID *int   `json:"ptu_id,omitempty"`
}

// Start begins processing a speech-to-text bulk job.
//
// Call this method after all audio files have been uploaded to start the
// transcription process. The job will transition from Accepted/Pending to
// Running state.
//
// # Parameters
//
//	ctx: Context for the request (used for cancellation and timeouts)
//	jobID: The unique identifier returned from Initialize()
//	opts: Optional functional options to configure the start request
//
// # Returns
//
//	JobTranscribeResponse containing the current job status, or an error
//
// # Functional Options
//
//	WithPtuID(int) - Pre-trained unit ID for custom model fine-tuning
//
// # Example
//
//	resp, err := client.SpeechToTextJob.Start(
//	    context.Background(),
//	    "job-id-from-initialize",
//	)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Job State: %s\n", resp.JobState)
//
// # Note
//
// Ensure all audio files are uploaded before calling Start().
// The job will only process files that have been successfully uploaded.
//
// # API Reference
//
// https://docs.sarvam.ai/api-reference-docs/speech-to-text/stt/job/start
func (c *SttJobClient) Start(ctx context.Context,
	jobID string,
	opts ...StartJobOption,
) (*JobTranscribeResponse, error) {
	var resp JobTranscribeResponse

	req := &startJobRequest{
		JobID: jobID,
	}

	for _, opt := range opts {
		opt(req)
	}

	if err := validateStartJobRequest(req); err != nil {
		return nil, err
	}

	var reqUrl string

	if req.PtuID != nil {
		reqUrl = "/speech-to-text/job/v1/" + url.PathEscape(jobID) + "/start?ptu_id=" + strconv.Itoa(*req.PtuID)
	} else {
		reqUrl = "/speech-to-text/job/v1/" + url.PathEscape(jobID) + "/start"
	}

	err := c.transport.DoRequest(
		ctx,
		"POST",
		reqUrl,
		nil,
		&resp,
		"application/json",
	)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetStatus retrieves the current status of a speech-to-text bulk job.
//
// Use this method to poll for job progress. The response includes overall
// job status as well as per-file details including success/failure counts.
//
// # Parameters
//
//	ctx: Context for the request (used for cancellation and timeouts)
//	jobID: The unique identifier returned from Initialize()
//
// # Returns
//
//	JobTranscribeResponse containing the current job status and file details, or an error
//
// # Rate Limiting
//
// To prevent rate limit errors, implement a minimum delay of 5 milliseconds
// between consecutive status polling requests.
//
// # Example
//
//	// Poll until job is completed or failed
//	for {
//	    resp, err := client.SpeechToTextJob.GetStatus(context.Background(), jobID)
//	    if err != nil {
//	        log.Fatal(err)
//	    }
//
//	    fmt.Printf("Job State: %s\n", resp.JobState)
//	    fmt.Printf("Progress: %d/%d files\n", *resp.SuccessfulFilesCount, *resp.TotalFiles)
//
//	    if resp.JobState == speech.JobStateCompleted || resp.JobState == speech.JobStateFailed {
//	        break
//	    }
//
//	    time.Sleep(2 * time.Second) // Wait before next poll
//	}
//
// # API Reference
//
// https://docs.sarvam.ai/api-reference-docs/speech-to-text/stt/job/status
func (c *SttJobClient) GetStatus(ctx context.Context, jobID string) (*JobTranscribeResponse, error) {
	var resp JobTranscribeResponse

	if err := validateGetStatusRequest(jobID); err != nil {
		return nil, err
	}

	url := "/speech-to-text/job/v1/" + url.PathEscape(jobID) + "/status"
	err := c.transport.DoRequest(
		ctx,
		"GET",
		url,
		nil,
		&resp,
		"application/json",
	)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// getDownloadLinksRequest is the internal request structure for getting download links.
type getDownloadLinksRequest struct {
	JobID string   `json:"job_id"`
	Files []string `json:"files"`
}

// GetDownloadLinks generates presigned URLs for downloading transcription results.
//
// After a job is completed, use this method to get presigned URLs for downloading
// the transcription output files. You can download all files or specify specific
// file IDs to download.
//
// # Parameters
//
//	ctx: Context for the request (used for cancellation and timeouts)
//	jobID: The unique identifier returned from Initialize()
//	files: List of file IDs to download (from JobDetails in GetStatus response)
//
// # Returns
//
//	GetDownloadLinksResponse containing presigned download URLs, or an error
//
// # Example
//
//	// First get the job status to find file IDs
//	status, err := client.SpeechToTextJob.GetStatus(context.Background(), jobID)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Collect file IDs from job details
//	var fileIDs []string
//	if status.JobDetails != nil {
//	    for _, detail := range *status.JobDetails {
//	        for _, input := range detail.Inputs {
//	            fileIDs = append(fileIDs, input.FileId)
//	        }
//	    }
//	}
//
//	// Get download URLs
//	resp, err := client.SpeechToTextJob.GetDownloadLinks(
//	    context.Background(),
//	    jobID,
//	    fileIDs,
//	)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Download each transcription
//	for fileID, downloadInfo := range resp.DownloadUrls {
//	    fmt.Printf("Downloading %s from %s\n", fileID, downloadInfo.FileUrl)
//	    // Make GET request to download the file
//	}
//
// # API Reference
//
// https://docs.sarvam.ai/api-reference-docs/speech-to-text/stt/job/download
func (c *SttJobClient) GetDownloadLinks(ctx context.Context,
	jobID string,
	files []string,
) (*GetDownloadLinksResponse, error) {

	var resp GetDownloadLinksResponse

	req := &getDownloadLinksRequest{
		JobID: jobID,
		Files: files,
	}

	err := validateGetDownloadLinksRequest(req)
	if err != nil {
		return nil, err
	}

	err = c.transport.DoRequest(
		ctx,
		"POST",
		"/speech-to-text/job/v1/download-files",
		req,
		&resp,
		"application/json",
	)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}
