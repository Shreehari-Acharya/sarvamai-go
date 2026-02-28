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

func validateChatRequest(model string, messages []ChatMessage) error {
	if err := validateModel(model); err != nil {
		return err
	}
	if err := validateMessages(messages); err != nil {
		return err
	}
	return nil
}
