package detect

type Request struct {
	Input string `json:"input"`
}

type Response struct {
	RequestID    *string `json:"request_id"`
	LanguageCode *string `json:"language_code"`
	ScriptCode   *string `json:"script_code"`
}

func (r Request) Validate() error {
	return validateDetectLanguageInput(r.Input)
}
