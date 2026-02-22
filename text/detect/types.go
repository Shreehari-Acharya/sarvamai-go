package detect

// Request represents a language detection request.
//
// # Fields
//
//   - Input: The text to analyze (required). Maximum 1000 characters.
//
// # Example
//
//	req := detect.Request{
//	    Input: "नमस्ते कैसे हैं आप",
//	}
type Request struct {
	Input string `json:"input"`
}

// Response represents a language detection response.
//
// # Fields
//
//   - RequestID: Unique identifier for the request
//   - LanguageCode: Detected language code (e.g., "hi-IN" for Hindi)
//   - ScriptCode: Detected script code (e.g., "Devanagari")
type Response struct {
	RequestID    *string `json:"request_id"`
	LanguageCode *string `json:"language_code"`
	ScriptCode   *string `json:"script_code"`
}

// Validate validates the detection request.
// Returns an error if input is empty or exceeds 1000 characters.
func (r Request) Validate() error {
	return validateDetectLanguageInput(r.Input)
}
