package text

import (
	"context"
)

// Translation converts text from one language to another while preserving its meaning.
// For Example: ‘मैं ऑफिस जा रहा हूँ’ translates to ‘I am going to the office’ in English,
// where the script and language change, but the original meaning remains the same.
func (s *Service) Translate(
	ctx context.Context,
	req TranslateRequest,
) (*TranslateResponse, error) {

	var resp TranslateResponse

	err := s.transport.DoRequest(
		ctx,
		"POST",
		"/translate",
		req,
		&resp,
		"application/json",
	)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
