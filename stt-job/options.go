package sttjob

import (
	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/sarvamaierrors"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/languages"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/shared/speech"
)

// InitJobOption is a functional option for configuring an Initialize request.
type InitJobOption func(*initJobRequest) error

// WithLanguage sets the language code for the transcription job.
//
// The language code follows BCP-47 format (e.g., "hi-IN" for Hindi).
// Use "unknown" (default) to let the server auto-detect the language.
//
// # Supported Languages
//
// Saarika v2.5 (12 languages): unknown (auto), hi-IN, bn-IN, kn-IN, ml-IN, mr-IN,
// or-IN, pa-IN, ta-IN, te-IN, en-IN, gu-IN
//
// Saaras v3 (23 languages): All above plus as-IN, ur-IN, ne-IN, kok-IN, ks-IN,
// sd-IN, sa-IN, sat-IN, mni-IN, brx-IN, mai-IN, doi-IN
//
// # Example
//
//	client.Initialize(ctx, WithLanguage(languages.ToLanguageCode("hi-IN")))
func WithLanguage(language languages.Code) InitJobOption {
	return func(req *initJobRequest) error {

		if language == "" {
			return nil
		}
		if !languages.SaarasLanguages[language] {
			return &sarvamaierrors.ValidationError{
				Field:   "language",
				Message: "invalid language code.",
			}
		}
		req.JobParameters.LanguageCode = &language
		return nil
	}
}

// WithModel sets the speech recognition model for the transcription job.
//
// # Models
//
//   - ModelSaarika (saarika:v2.5): Default multilingual model, supports 12 languages
//   - ModelSaaras (saaras:v3): Advanced model, supports 23 languages with multiple modes
//
// # Example
//
//	client.Initialize(ctx, WithModel(speech.ModelSaaras))
func WithModel(model speech.Model) InitJobOption {

	return func(req *initJobRequest) error {
		if model != speech.ModelSaarika && model != speech.ModelSaaras {
			return &sarvamaierrors.ValidationError{
				Field:   "model",
				Message: "invalid model. supported models are saarika:v2.5 and saaras:v3",
			}
		}
		req.JobParameters.Model = &model
		return nil
	}
}

// WithMode sets the processing mode for the transcription job.
//
// This option is only applicable when using saaras:v3 model.
//
// # Modes
//
//   - ModeTranscribe (default): Standard transcription with proper formatting
//   - ModeTranslate: Translates speech from Indic languages to English
//   - ModeVerbatim: Exact word-for-word transcription without normalization
//   - ModeTranslit: Romanization - converts to Latin script only
//   - ModeCodemix: Mixed script (English in English, Indic in native script)
//
// # Example
//
//	client.Initialize(ctx, WithMode(speech.ModeTranslate))
func WithMode(mode speech.Mode) InitJobOption {
	return func(req *initJobRequest) error {
		req.JobParameters.Mode = &mode
		return nil
	}
}

// WithTimeStamps enables or disables word-level timestamps in the transcription response.
//
// When enabled, the response will include start and end timestamps for each word.
// This is useful for applications that need precise timing information.
//
// # Default
//
// false (timestamps disabled)
//
// # Example
//
//	client.Initialize(ctx, WithTimeStamps(true))
func WithTimeStamps(enabled bool) InitJobOption {
	return func(req *initJobRequest) error {
		req.JobParameters.WithTimestamps = &enabled
		return nil
	}
}

// WithDiarization enables or disables speaker diarization.
//
// Speaker diarization identifies and separates different speakers in the audio.
// This is useful for multi-speaker audio transcription.
//
// Note: This feature is currently in beta mode.
//
// # Default
//
// false (diarization disabled)
//
// # Example
//
//	client.Initialize(ctx, WithDiarization(true))
func WithDiarization(enabled bool) InitJobOption {
	return func(req *initJobRequest) error {
		req.JobParameters.WithDiarization = &enabled
		return nil
	}
}

// WithNumSpeakers sets the expected number of speakers for diarization.
//
// This option helps the model identify and separate speakers more accurately.
// It requires that WithDiarization(true) is also set.
//
// Note: This feature is currently in beta mode and requires diarization to be enabled.
//
// # Default
//
// nil (auto-detect number of speakers)
//
// # Example
//
//	client.Initialize(ctx,
//	    WithDiarization(true),
//	    WithNumSpeakers(2),
//	)
func WithNumSpeakers(num int) InitJobOption {
	return func(req *initJobRequest) error {
		req.JobParameters.NumSpeakers = &num
		return nil
	}
}

// WithCallback sets the webhook URL and optional auth token for job completion notifications.
//
// The server will send a POST request to the specified URL when the job completes.
// The auth_token (if provided) will be included in the callback request headers
// for verification.
//
// Note: The callback feature requires your server to be publicly accessible.
//
// # Parameters
//
//	url: The webhook URL to call when the job completes (required)
//	authToken: Optional authorization token for callback verification (optional)
//
// # Example
//
//	// Simple callback
//	token := "my-secret-token"
//	client.Initialize(ctx, WithCallback("https://my-server.com/webhook", &token))
//
//	// Without auth token
//	client.Initialize(ctx, WithCallback("https://my-server.com/webhook", nil))
func WithCallback(url string, authToken *string) InitJobOption {
	return func(req *initJobRequest) error {
		if url == "" {
			return &sarvamaierrors.ValidationError{
				Field:   "callback_url",
				Message: "callback URL cannot be empty",
			}
		}
		req.Callback = &speech.Callback{
			URL:       url,
			AuthToken: authToken,
		}
		return nil
	}
}

// StartJobOption is a functional option for configuring a Start request.
type StartJobOption func(*startJobRequest) error

// WithPtuID sets the Pre-Trained Unit (PTU) ID for the transcription job.
//
// The PTU ID is used when you want to use a custom fine-tuned model for
// transcription. This requires prior setup and configuration of custom models.
//
// # Parameters
//
//	ptuID: The PTU ID for custom model fine-tuning
//
// # Example
//
//	client.Start(ctx, jobID, WithPtuID(12345))
func WithPtuID(ptuID int) StartJobOption {
	return func(req *startJobRequest) error {
		req.PtuID = &ptuID
		return nil
	}
}
