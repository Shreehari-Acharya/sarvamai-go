package docintel

import (
	"context"
	"net/url"

	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/transport"
)

// DocIntelClient provides access to the Document Intelligence API.
type DocIntelClient struct {
	transport *transport.Transport
}

// NewDocIntelClient creates a new Document Intelligence client.
func NewDocIntelClient(t *transport.Transport) *DocIntelClient {
	return &DocIntelClient{
		transport: t,
	}
}

type docIntelInitializeRequest struct {
	JobParameters *JobParameters `json:"job_parameters,omitempty"`
	Callback      *Callback      `json:"callback,omitempty"`
}

// Initialize creates a new Document Intelligence job.
//
// # Parameters
//
//	ctx: Context for the request
//	opts: Optional functional options to configure the job
//
// # Returns
//
//	DocIntelInitializeResponse containing the job ID and status, or an error
//
// # Functional Options
//
//	WithLanguage(languages.Code)      - Language of the document (BCP-47 format)
//	WithOutputFormat(OutputFormat)   - Output format (html, md, json)
//	WithCallback(url string, authToken *string)    - Webhook URL and optional auth token
//
// # Example
//
//	resp, err := client.DocumentIntelligence.Initialize(
//	    ctx,
//	    docintel.WithLanguage("hi-IN"),
//	    docintel.WithOutputFormat(docintel.OutputFormatMD),
//	)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(resp.JobID)
//
// # API Reference
//
// https://docs.sarvam.ai/api-reference-docs/document-intelligence/initialise
func (c *DocIntelClient) Initialize(ctx context.Context, options ...docIntelOption) (*DocIntelInitializeResponse, error) {

	req := &docIntelInitializeRequest{}

	for _, option := range options {
		if err := option(req); err != nil {
			return nil, err
		}
	}

	if err := validateInitializeRequest(req); err != nil {
		return nil, err
	}

	var resp DocIntelInitializeResponse

	err := c.transport.DoRequest(
		ctx,
		"POST",
		"/doc-digitization/job/v1",
		req,
		&resp,
		"application/json",
	)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

type docIntelGetUploadLinksRequest struct {
	JobID string   `json:"job_id"`
	Files []string `json:"files"`
}

// GetUploadLinks retrieves presigned URLs for uploading files to a job.
//
// # Parameters
//
//	ctx: Context for the request
//	jobID: Job identifier returned from Initialize
//	filename: Name of the file to upload (must be .pdf or .zip)
//
// # Returns
//
//	DocIntelGetUploadLinksResponse containing upload URLs, or an error
//
// # Example
//
//	resp, err := client.DocumentIntelligence.GetUploadLinks(ctx, jobID, "document.pdf")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(resp.UploadUrls["document.pdf"].FileURL)
//
// # API Reference
//
// https://docs.sarvam.ai/api-reference-docs/document-intelligence/get-upload-links
func (c *DocIntelClient) GetUploadLinks(ctx context.Context, jobID string, filename string) (*DocIntelGetUploadLinksResponse, error) {

	if err := validateGetUploadLinksRequest(jobID, filename); err != nil {
		return nil, err
	}

	var resp DocIntelGetUploadLinksResponse

	req := &docIntelGetUploadLinksRequest{
		JobID: jobID,
		Files: []string{filename},
	}

	err := c.transport.DoRequest(
		ctx,
		"POST",
		"/doc-digitization/job/v1/upload-files",
		req,
		&resp,
		"application/json",
	)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// Start initiates processing of a document intelligence job.
//
// Call this after uploading the file to start processing.
// The job processes asynchronously; use GetStatus or webhooks to monitor progress.
//
// # Parameters
//
//	ctx: Context for the request
//	jobID: Job identifier returned from Initialize
//
// # Returns
//
//	DocIntelJobStatusResponse containing the job status, or an error
//
// # Example
//
//	resp, err := client.DocumentIntelligence.Start(ctx, jobID)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(resp.JobState)
//
// # API Reference
//
// https://docs.sarvam.ai/api-reference-docs/document-intelligence/start
func (c *DocIntelClient) Start(ctx context.Context, jobID string) (*DocIntelJobStatusResponse, error) {

	if err := validateJobID(jobID); err != nil {
		return nil, err
	}

	var resp DocIntelJobStatusResponse

	url := "/doc-digitization/job/v1/" + url.PathEscape(jobID) + "/start"
	err := c.transport.DoRequest(
		ctx,
		"POST",
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

// GetStatus retrieves the current status of a document intelligence job.
//
// # Parameters
//
//	ctx: Context for the request
//	jobID: Job identifier returned from Initialize
//
// # Returns
//
//	DocIntelJobStatusResponse containing detailed job status, or an error
//
// # Example
//
//	resp, err := client.DocumentIntelligence.GetStatus(ctx, jobID)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("State: %s, Total Pages: %d\n", resp.JobState, resp.JobDetails[0].TotalPages)
//
// # API Reference
//
// https://docs.sarvam.ai/api-reference-docs/document-intelligence/get-status
func (c *DocIntelClient) GetStatus(ctx context.Context, jobID string) (*DocIntelJobStatusResponse, error) {

	if err := validateJobID(jobID); err != nil {
		return nil, err
	}

	var resp DocIntelJobStatusResponse

	url := "/doc-digitization/job/v1/" + url.PathEscape(jobID) + "/status"

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

// GetDownloadLinks retrieves presigned URLs for downloading processed output files.
//
// Job must be in Completed or PartiallyCompleted state before calling this.
//
// # Parameters
//
//	ctx: Context for the request
//	jobID: Job identifier returned from Initialize
//
// # Returns
//
//	DocIntelGetDownloadLinksResponse containing download URLs, or an error
//
// # Example
//
//	resp, err := client.DocumentIntelligence.GetDownloadLinks(ctx, jobID)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for name, details := range resp.DownloadURLs {
//	    fmt.Printf("Download %s: %s\n", name, details.FileURL)
//	}
//
// # API Reference
//
// https://docs.sarvam.ai/api-reference-docs/document-intelligence/get-download-links
func (c *DocIntelClient) GetDownloadLinks(ctx context.Context, jobID string) (*DocIntelGetDownloadLinksResponse, error) {

	if err := validateJobID(jobID); err != nil {
		return nil, err
	}

	var resp DocIntelGetDownloadLinksResponse

	url := "/doc-digitization/job/v1/" + url.PathEscape(jobID) + "/download-files"

	err := c.transport.DoRequest(
		ctx,
		"POST",
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
