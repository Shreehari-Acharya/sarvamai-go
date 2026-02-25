package docintel

import (
	"github.com/Shreehari-Acharya/sarvam-go-sdk/languages"
)

// docIntelOption is a functional option for configuring a Document Intelligence job.
type docIntelOption func(*docIntelInitializeRequest) error

// WithOutputFormat sets the output format for the extracted document content.
//
// Output is delivered as a ZIP file containing the processed documents.
//
// Supported formats:
//   - OutputFormatHTML: Structured HTML files with layout preservation
//   - OutputFormatMD: Markdown files (default)
//   - OutputFormatJSON: Structured JSON files for programmatic processing
//
// Default: OutputFormatMD (Markdown)
func WithOutputFormat(format OutputFormat) docIntelOption {
	return func(req *docIntelInitializeRequest) error {
		if req.JobParameters == nil {
			req.JobParameters = &JobParameters{}
		}
		req.JobParameters.OutputFormat = &format
		return nil
	}
}

// WithLanguage sets the primary language of the document being processed.
//
// This helps optimize text extraction accuracy for the specified language.
// The language should be specified in BCP-47 format (e.g., "hi-IN" for Hindi).
//
// Supported languages:
//   - hi-IN: Hindi (default)
//   - en-IN: English
//   - bn-IN: Bengali
//   - gu-IN: Gujarati
//   - kn-IN: Kannada
//   - ml-IN: Malayalam
//   - mr-IN: Marathi
//   - od-IN: Odia
//   - pa-IN: Punjabi
//   - ta-IN: Tamil
//   - te-IN: Telugu
//   - ur-IN: Urdu
//   - as-IN: Assamese
//   - bodo-IN: Bodo
//   - doi-IN: Dogri
//   - ks-IN: Kashmiri
//   - kok-IN: Konkani
//   - mai-IN: Maithili
//   - mni-IN: Manipuri
//   - ne-IN: Nepali
//   - sa-IN: Sanskrit
//   - sat-IN: Santali
//   - sd-IN: Sindhi
//
// Default: hi-IN (Hindi)
func WithLanguage(language languages.Code) docIntelOption {
	return func(req *docIntelInitializeRequest) error {
		if req.JobParameters == nil {
			req.JobParameters = &JobParameters{}
		}
		req.JobParameters.Language = &language
		return nil
	}
}

// WithCallback configures a webhook to receive notifications when job processing completes.
//
// The callback URL must use HTTPS scheme. When the job finishes processing,
// a POST request will be sent to the specified URL.
//
// The optional authToken will be sent as X-SARVAM-JOB-CALLBACK-TOKEN header
// in the webhook request, allowing you to verify the request is from Sarvam.
//
// Parameters:
//   - callbackURL: HTTPS URL to receive the webhook notification
//   - authToken: Optional authorization token for verification
//
// Example:
//
//	authToken := "my-secret-token"
//	client.DocumentIntelligence.Initialize(ctx,
//	    docintel.WithCallback("https://example.com/webhook", &authToken),
//	)
func WithCallback(callbackURL string, authToken *string) docIntelOption {
	return func(req *docIntelInitializeRequest) error {
		if err := validateCallbackURL(callbackURL); err != nil {
			return err
		}
		req.Callback = &Callback{
			URL:       callbackURL,
			AuthToken: authToken,
		}
		return nil
	}
}
