package translatejob

import (
	"github.com/Shreehari-Acharya/sarvamai-go/shared/speech"
)

// initJobOption is a functional option for configuring an Initialize request.
type initJobOption func(*initJobRequest)

// WithPrompt sets an optional prompt to assist the transcription.
//
// The prompt can provide context or instructions to improve transcription accuracy.
// For example, it can include domain-specific vocabulary or formatting guidelines.
//
// # Example
//
//	client.SpeechToTextTranslateJob.Initialize(ctx, translatejob.WithPrompt("This is a medical transcription"))
func WithPrompt(prompt string) initJobOption {
	return func(req *initJobRequest) {
		req.JobParameters.Prompt = &prompt
	}
}

// WithModel sets the translation model for the job.
//
// Currently, only "saaras:v2.5" (speech.ModelSaarasV25) is supported for
// speech-to-text translation. If not specified, the API defaults to this model.
//
// # Example
//
//	client.SpeechToTextTranslateJob.Initialize(ctx, translatejob.WithModel(speech.ModelSaarasV25))
func WithModel(model speech.Model) initJobOption {
	return func(req *initJobRequest) {
		req.JobParameters.Model = &model
	}
}

// WithPtuId sets the PTU (Processing Time Unit) ID for the request.
//
// The PTU ID is an optional parameter that may be required for certain
// subscription tiers.
//
// # Example
//
//	client.SpeechToTextTranslateJob.Initialize(ctx, translatejob.WithPtuId(1))
func WithPtuId(ptuId int) initJobOption {
	return func(req *initJobRequest) {
		req.PtuID = &ptuId
	}
}

// WithDiarization enables speaker diarization for the translation job.
//
// When enabled, the API will identify and separate different speakers in the
// audio file. This is useful for multi-speaker audio transcription.
// Note: This feature is currently in Beta.
//
// # Example
//
//	client.SpeechToTextTranslateJob.Initialize(ctx, translatejob.WithDiarization(true))
func WithDiarization(withDiarization bool) initJobOption {
	return func(req *initJobRequest) {
		req.JobParameters.WithDiarization = &withDiarization
	}
}

// WithNumSpeakers sets the expected number of speakers for diarization.
//
// This is used when WithDiarization is enabled. If not specified, the API
// will attempt to automatically detect the number of speakers.
//
// # Example
//
//	client.SpeechToTextTranslateJob.Initialize(ctx, translatejob.WithDiarization(true), translatejob.WithNumSpeakers(2))
func WithNumSpeakers(numSpeakers int) initJobOption {
	return func(req *initJobRequest) {
		req.JobParameters.NumSpeakers = &numSpeakers
	}
}

// WithCallback sets a callback URL to be notified when the job completes.
//
// When the job finishes (success or failure), the API will send a POST
// request to the specified URL with the job status. The authToken will be
// included in the request headers for verification.
//
// # Parameters
//
//	callbackURL: The URL to call when the job completes
//	authToken: Optional authorization token to include in the callback request
//
// # Example
//
//	token := "optional-auth-token"
//	client.SpeechToTextTranslateJob.Initialize(ctx, translatejob.WithCallback("https://myapp.com/callback", &token))
func WithCallback(callbackURL string, authToken *string) initJobOption {
	return func(req *initJobRequest) {
		req.Callback = &speech.Callback{
			URL:       callbackURL,
			AuthToken: authToken,
		}
	}
}

// getUploadLinksOption is a functional option for configuring a GetUploadLinks request.
type getUploadLinksOption func(*getUploadLinksRequest)

// WithGetUploadLinksPtuId sets the PTU ID for the upload links request.
//
// # Example
//
//	client.SpeechToTextTranslateJob.GetUploadLinks(ctx, jobID, files, translatejob.WithGetUploadLinksPtuId(1))
func WithGetUploadLinksPtuId(ptuId int) getUploadLinksOption {
	return func(req *getUploadLinksRequest) {
		req.PtuID = &ptuId
	}
}

// startJobOption is a functional option for configuring a Start request.
type startJobOption func(*startJobRequest)

// WithStartJobPtuId sets the PTU ID for the start job request.
//
// # Example
//
//	client.SpeechToTextTranslateJob.Start(ctx, jobID, translatejob.WithStartJobPtuId(1))
func WithStartJobPtuId(ptuId int) startJobOption {
	return func(req *startJobRequest) {
		req.PtuID = &ptuId
	}
}

// getDownloadLinksOption is a functional option for configuring a GetDownloadLinks request.
type getDownloadLinksOption func(*getDownloadLinksRequest)

// WithGetDownloadLinksPtuId sets the PTU ID for the download links request.
//
// # Example
//
//	client.SpeechToTextTranslateJob.GetDownloadLinks(ctx, jobID, files, translatejob.WithGetDownloadLinksPtuId(1))
func WithGetDownloadLinksPtuId(ptuId int) getDownloadLinksOption {
	return func(req *getDownloadLinksRequest) {
		req.PtuID = &ptuId
	}
}
