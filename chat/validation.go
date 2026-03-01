package chat

import (
	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/sarvamaierrors"
)

func validateModel(model string) error {
	if model == "" {
		return &sarvamaierrors.ValidationError{
			Field:   "model",
			Message: "model is required",
		}
	}
	return nil
}

func validateMessages(messages []ChatMessage) error {
	if len(messages) == 0 {
		return &sarvamaierrors.ValidationError{
			Field:   "messages",
			Message: "messages is required and must contain at least one message",
		}
	}

	for i, msg := range messages {
		if msg.Content == "" {
			return &sarvamaierrors.ValidationError{
				Field:   "messages",
				Message: "message content cannot be empty",
			}
		}
		if msg.Role == "" {
			return &sarvamaierrors.ValidationError{
				Field:   "messages",
				Message: "message role cannot be empty",
			}
		}
		_ = i
	}

	return nil
}

var allowedReasoningEffortValues = map[ReasoningEffort]bool{
	ReasoningEffortLow:    true,
	ReasoningEffortMedium: true,
	ReasoningEffortHigh:   true,
}

func validateChatRequest(req *chatRequest) error {
	if err := validateModel(req.Model); err != nil {
		return err
	}
	if err := validateMessages(req.Messages); err != nil {
		return err
	}

	if req.Temperature != nil {
		temp := *req.Temperature
		if temp < 0.0 || temp > 2.0 {
			return &sarvamaierrors.ValidationError{
				Message: "temperature must be between 0.0 and 2.0",
				Field:   "temperature",
			}
		}
	}

	if req.TopP != nil {
		topP := *req.TopP
		if topP < 0.0 || topP > 1.0 {
			return &sarvamaierrors.ValidationError{
				Message: "top_p must be between 0.0 and 1.0",
				Field:   "top_p",
			}
		}
	}

	if req.ReasoningEffort != nil {
		effort := ReasoningEffort(*req.ReasoningEffort)
		if !allowedReasoningEffortValues[effort] {
			return &sarvamaierrors.ValidationError{
				Message: "reasoning_effort must be one of 'low', 'medium', or 'high'",
				Field:   "reasoning_effort",
			}
		}
	}

	if req.N != nil {
		n := *req.N
		if n < 1 || n > 128 {
			return &sarvamaierrors.ValidationError{
				Message: "n must be between 1 and 128",
				Field:   "n",
			}
		}
	}

	if req.FrequencyPenalty != nil {
		penalty := *req.FrequencyPenalty
		if penalty < -2.0 || penalty > 2.0 {
			return &sarvamaierrors.ValidationError{
				Message: "frequency_penalty must be between -2.0 and 2.0",
				Field:   "frequency_penalty",
			}
		}
	}

	if req.PresencePenalty != nil {
		penalty := *req.PresencePenalty
		if penalty < -2.0 || penalty > 2.0 {
			return &sarvamaierrors.ValidationError{
				Message: "presence_penalty must be between -2.0 and 2.0",
				Field:   "presence_penalty",
			}
		}
	}

	return nil
}
