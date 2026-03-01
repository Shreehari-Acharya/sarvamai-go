package translatejob

import (
	"github.com/Shreehari-Acharya/sarvamai-go/internal/sarvamaierrors"
)

func validateInitJobRequest(req *initJobRequest) error {
	if req.Callback != nil {
		if req.Callback.URL == "" {
			return &sarvamaierrors.ValidationError{
				Field:   "callback_url",
				Message: "callback URL cannot be empty",
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
