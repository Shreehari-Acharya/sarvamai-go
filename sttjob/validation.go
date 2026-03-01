package sttjob

import (
	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/sarvamaierrors"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/shared/speech"
)

func validateInitJobRequest(req *initJobRequest) error {

	model := req.JobParameters.Model
	mode := req.JobParameters.Mode
	language := req.JobParameters.LanguageCode
	withDiarization := req.JobParameters.WithDiarization
	numSpeakers := req.JobParameters.NumSpeakers

	if err := speech.ValidateMode(model, mode); err != nil {
		return err
	}

	if language != nil {
		if err := speech.ValidateLanguageWithSpec(model, *language, true); err != nil {
			return err
		}
	}

	if numSpeakers != nil {
		if withDiarization == nil || !*withDiarization {
			return &sarvamaierrors.ValidationError{
				Field:   "num_speakers",
				Message: "num_speakers is only applicable when with_diarization is true",
			}
		}
	}

	return nil
}

func validateGetUploadLinksRequest(req *getUploadLinksRequest) error {
	if req.JobID == "" {
		return &sarvamaierrors.ValidationError{
			Field:   "job_id",
			Message: "job_id is required",
		}
	}

	if len(req.Files) == 0 {
		return &sarvamaierrors.ValidationError{
			Field:   "files",
			Message: "at least one file is required",
		}
	}

	return nil
}

func validateStartJobRequest(req *startJobRequest) error {
	if req.JobID == "" {
		return &sarvamaierrors.ValidationError{
			Field:   "job_id",
			Message: "job_id is required",
		}
	}
	return nil
}

func validateGetStatusRequest(jobID string) error {
	if jobID == "" {
		return &sarvamaierrors.ValidationError{
			Field:   "job_id",
			Message: "job_id is required",
		}
	}
	return nil
}

func validateGetDownloadLinksRequest(req *getDownloadLinksRequest) error {
	if req.JobID == "" {
		return &sarvamaierrors.ValidationError{
			Field:   "job_id",
			Message: "job_id is required",
		}
	}

	if len(req.Files) == 0 {
		return &sarvamaierrors.ValidationError{
			Field:   "files",
			Message: "at least one file is required",
		}
	}

	return nil
}
