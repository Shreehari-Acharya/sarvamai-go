package translatejob

import (
	"context"
	"net/url"
	"strconv"

	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/sarvamaierrors"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/transport"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/shared/speech"
)

// TranslateJobClient provides access to the Speech-to-Text Translate Batch Job API.
// It allows processing multiple audio files with translation through a bulk job workflow.
//
// The typical workflow is:
//   - Initialize a new bulk job
//   - Get upload links and upload audio files
//   - Start processing
//   - Poll for status until completed
//   - Get download links for results
type TranslateJobClient struct {
	transport *transport.Transport
}

// NewTranslateJobClient creates a new TranslateJobClient.
//
// # Parameters
//
//	t: Transport instance configured with API key and base URL
//
// # Returns
//
//	A new TranslateJobClient instance
func NewTranslateJobClient(t *transport.Transport) *TranslateJobClient {
	return &TranslateJobClient{
		transport: t,
	}
}

// initJobRequest is the internal request structure for initializing a bulk job.
type initJobRequest struct {
	JobParameters JobParameters    `json:"job_parameters"`
	PtuID         *int             `json:"ptu_id,omitempty"`
	Callback      *speech.Callback `json:"callback,omitempty"`
}

// Initialize creates a new speech-to-text translate bulk job.
//
// This is the first step in the batch processing workflow. It creates a new job
// and returns the job ID and storage container details for uploading files.
//
// # Parameters
//
//	ctx: Context for the request (used for cancellation and timeouts)
//	opts: Optional functional options to configure the job
//
// # Returns
//
//	[InitJobResponse] containing job ID and storage details, or an error
//
// # Functional Options
//
//	[WithPrompt] - Set a prompt to assist transcription
//	[WithModel] - Set the translation model (currently only saaras:v2.5 supported)
//	[WithDiarization] - Enable speaker diarization
//	[WithNumSpeakers] - Set expected number of speakers
//	[WithPtuId] - Set the PTU ID for the request
//	[WithCallback] - Set a callback URL for job completion notifications
//
// # Example
//
//	resp, err := client.SpeechToTextTranslateJob.Initialize(
//	    context.Background(),
//	    translatejob.WithDiarization(true),
//	    translatejob.WithNumSpeakers(2),
//	    translatejob.WithPrompt("This is a customer service call"),
//	)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Job ID: %s\n", resp.JobID)
//
// # API Reference
//
// https://docs.sarvam.ai/api-reference-docs/speech-to-text-translate/stt-translate/job/initiate
func (c *TranslateJobClient) Initialize(ctx context.Context, opts ...initJobOption) (*InitJobResponse, error) {
	req := &initJobRequest{
		JobParameters: JobParameters{},
	}

	for _, option := range opts {
		if err := option(req); err != nil {
			return nil, err
		}
	}

	var reqUrl string
	if req.PtuID == nil {
		reqUrl = "/speech-to-text-translate/job/v1"
	} else {
		reqUrl = "/speech-to-text-translate/job/v1/?ptu_id=" + strconv.Itoa(*req.PtuID)
	}

	var resp InitJobResponse
	err := c.transport.DoRequest(
		ctx,
		"POST",
		reqUrl,
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
	PtuID *int     `json:"ptu_id,omitempty"`
	Files []string `json:"files"`
}

// GetUploadLinks generates presigned URLs for uploading audio files to the job.
//
// After initializing a job, use this method to get presigned URLs where each
// audio file should be uploaded. Upload your files to these URLs before
// starting the job.
//
// # Parameters
//
//	ctx: Context for the request
//	jobID: The job ID from [Initialize]
//	files: List of filenames to generate upload URLs for
//	opts: Optional functional options
//
// # Returns
//
//	[GetUploadLinksResponse] containing presigned upload URLs, or an error
//
// # Functional Options
//
//	[WithGetUploadLinksPtuId] - Set the PTU ID for the request
//
// # Example
//
//	resp, err := client.SpeechToTextTranslateJob.GetUploadLinks(
//	    context.Background(),
//	    "job-id-from-initialize",
//	    []string{"audio1.mp3", "audio2.mp3"},
//	)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for filename, urlDetails := range resp.UploadUrls {
//	    fmt.Printf("Upload %s to: %s\n", filename, urlDetails.FileUrl)
//	}
//
// # API Reference
//
// https://docs.sarvam.ai/api-reference-docs/speech-to-text-translate/stt-translate/job/upload
func (c *TranslateJobClient) GetUploadLinks(ctx context.Context,
	jobID string,
	files []string,
	opts ...getUploadLinksOption) (*GetUploadLinksResponse, error) {
	req := &getUploadLinksRequest{
		JobID: jobID,
		Files: files,
	}

	var resp GetUploadLinksResponse

	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, err
		}
	}

	if err := validateGetUploadLinksRequest(req); err != nil {
		return nil, err
	}

	var reqUrl string
	if req.PtuID == nil {
		reqUrl = "/speech-to-text-translate/job/v1/upload-files"
	} else {
		reqUrl = "/speech-to-text-translate/job/v1/upload-files?ptu_id=" + strconv.Itoa(*req.PtuID)
	}

	err := c.transport.DoRequest(
		ctx,
		"POST",
		reqUrl,
		req,
		&resp,
		"application/json",
	)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// startJobRequest is the internal request structure for starting a job.
type startJobRequest struct {
	JobID string `json:"job_id"`
	PtuID *int   `json:"ptu_id,omitempty"`
}

// Start begins processing a bulk translation job.
//
// Call this after uploading all audio files. The job will transition from
// Pending to Running state. Use [GetStatus] to poll for completion.
//
// # Parameters
//
//	ctx: Context for the request
//	jobID: The job ID from [Initialize]
//	opts: Optional functional options
//
// # Returns
//
//	[JobStatusResponse] with current job status, or an error
//
// # Functional Options
//
//	[WithStartJobPtuId] - Set the PTU ID for the request
//
// # Example
//
//	status, err := client.SpeechToTextTranslateJob.Start(
//	    context.Background(),
//	    "job-id-from-initialize",
//	)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Job state: %s\n", status.JobState)
//
// # API Reference
//
// https://docs.sarvam.ai/api-reference-docs/speech-to-text-translate/stt-translate/job/start
func (c *TranslateJobClient) Start(ctx context.Context, jobID string, opts ...startJobOption) (*JobStatusResponse, error) {
	req := &startJobRequest{
		JobID: jobID,
	}

	var resp JobStatusResponse

	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, err
		}
	}

	if err := validateStartJobRequest(req); err != nil {
		return nil, err
	}

	var reqUrl string
	if req.PtuID == nil {
		reqUrl = "/speech-to-text-translate/job/v1/" + url.PathEscape(jobID) + "/start"
	} else {
		reqUrl = "/speech-to-text-translate/job/v1/" + url.PathEscape(jobID) + "/start?ptu_id=" + strconv.Itoa(*req.PtuID)
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

// GetStatus retrieves the current status of a bulk translation job.
//
// Use this to poll for job completion. The API recommends adding at least
// 5ms delay between consecutive status requests.
//
// # Parameters
//
//	ctx: Context for the request
//	jobID: The job ID from [Initialize]
//
// # Returns
//
//	[JobStatusResponse] containing current job status and file-level details, or an error
//
// # Example
//
//	for {
//	    status, err := client.SpeechToTextTranslateJob.GetStatus(context.Background(), "job-id")
//	    if err != nil {
//	        log.Fatal(err)
//	    }
//	    fmt.Printf("Job state: %s\n", status.JobState)
//	    if status.JobState == "Completed" || status.JobState == "Failed" {
//	        break
//	    }
//	    time.Sleep(5 * time.Second) // Respect rate limiting
//	}
//
// # API Reference
//
// https://docs.sarvam.ai/api-reference-docs/speech-to-text-translate/stt-translate/job/status
func (c *TranslateJobClient) GetStatus(ctx context.Context, jobID string) (*JobStatusResponse, error) {
	if jobID == "" {
		return nil, &sarvamaierrors.ValidationError{
			Field:   "job_id",
			Message: "job_id is required",
		}
	}

	var resp JobStatusResponse

	reqUrl := "/speech-to-text-translate/job/v1/" + url.PathEscape(jobID) + "/status"
	err := c.transport.DoRequest(
		ctx,
		"GET",
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

// getDownloadLinksRequest is the internal request structure for getting download links.
type getDownloadLinksRequest struct {
	JobID string   `json:"job_id"`
	Files []string `json:"files"`
	PtuID *int     `json:"ptu_id,omitempty"`
}

// GetDownloadLinks generates presigned URLs for downloading translated output files.
//
// Call this after the job status is "Completed" to get download URLs for
// the translated transcription files.
//
// # Parameters
//
//	ctx: Context for the request
//	jobID: The job ID from [Initialize]
//	files: List of filenames to get download URLs for
//	opts: Optional functional options
//
// # Returns
//
//	[GetDownloadLinksResponse] containing presigned download URLs, or an error
//
// # Functional Options
//
//	[WithGetDownloadLinksPtuId] - Set the PTU ID for the request
//
// # Example
//
//	// First check job status
//	status, err := client.SpeechToTextTranslateJob.GetStatus(context.Background(), "job-id")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	if status.JobState != speech.JobStateCompleted {
//	    fmt.Println("Job not completed yet")
//	    return
//	}
//
//	// Get download links for all files
//	resp, err := client.SpeechToTextTranslateJob.GetDownloadLinks(
//	    context.Background(),
//	    "job-id",
//	    []string{"audio1.mp3", "audio2.mp3"},
//	)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for filename, urlDetails := range resp.DownloadUrls {
//	    fmt.Printf("Download %s from: %s\n", filename, urlDetails.FileUrl)
//	}
//
// # API Reference
//
// https://docs.sarvam.ai/api-reference-docs/speech-to-text-translate/stt-translate/job/download
func (c *TranslateJobClient) GetDownloadLinks(ctx context.Context, jobID string, files []string, opts ...getDownloadLinksOption) (*GetDownloadLinksResponse, error) {
	req := &getDownloadLinksRequest{
		JobID: jobID,
		Files: files,
	}

	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, err
		}
	}

	if err := validateGetDownloadLinksRequest(req); err != nil {
		return nil, err
	}

	var reqUrl string

	if req.PtuID == nil {
		reqUrl = "/speech-to-text-translate/job/v1/download-files"
	} else {
		reqUrl = "/speech-to-text-translate/job/v1/download-files?ptu_id=" + url.QueryEscape(strconv.Itoa(*req.PtuID))
	}

	var resp GetDownloadLinksResponse
	err := c.transport.DoRequest(
		ctx,
		"POST",
		reqUrl,
		req,
		&resp,
		"application/json",
	)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
