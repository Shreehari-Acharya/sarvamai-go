package docintel

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/sarvamaierrors"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/languages"
)

func validateCallbackURL(callbackURL string) error {
	parsedURL, err := url.ParseRequestURI(callbackURL)
	if err != nil {
		return &sarvamaierrors.ValidationError{
			Field:   "callback_url",
			Message: "callback url must be a valid URL",
		}
	}
	if parsedURL.Scheme != "https" {
		return &sarvamaierrors.ValidationError{
			Field:   "callback_url",
			Message: "callback url must use HTTPS scheme",
		}
	}
	return nil
}

func validateLanguage(language *languages.Code) error {
	if language == nil {
		return nil
	}
	if !languages.AllowedDocIntelLanguages[*language] {
		return &sarvamaierrors.ValidationError{
			Field:   "language",
			Message: fmt.Sprintf("language %s is not supported for document intelligence", *language),
		}
	}
	return nil
}

func validateOutputFormat(format *OutputFormat) error {
	if format == nil {
		return nil
	}
	if !allowedOutputFormats[*format] {
		return &sarvamaierrors.ValidationError{
			Field:   "output_format",
			Message: "output_format must be one of: html, md, json",
		}
	}
	return nil
}

func validateJobID(jobID string) error {
	if strings.TrimSpace(jobID) == "" {
		return &sarvamaierrors.ValidationError{
			Field:   "job_id",
			Message: "job_id is required",
		}
	}
	return nil
}

func validateFile(filename string) error {
	if strings.TrimSpace(filename) == "" {
		return &sarvamaierrors.ValidationError{
			Field:   "filename",
			Message: "filename is required",
		}
	}
	ext := ""
	if idx := strings.LastIndex(filename, "."); idx != -1 {
		ext = filename[idx:]
	}
	if !allowedFileExtensions[ext] {
		return &sarvamaierrors.ValidationError{
			Field:   "filename",
			Message: "filename must have .pdf or .zip extension",
		}
	}
	return nil
}

func validateInitializeRequest(req *docIntelInitializeRequest) error {
	if req.JobParameters != nil {
		if err := validateLanguage(req.JobParameters.Language); err != nil {
			return err
		}
		if err := validateOutputFormat(req.JobParameters.OutputFormat); err != nil {
			return err
		}
	}
	return nil
}

func validateGetUploadLinksRequest(jobID string, filename string) error {
	if err := validateJobID(jobID); err != nil {
		return err
	}
	if err := validateFile(filename); err != nil {
		return err
	}
	return nil
}
