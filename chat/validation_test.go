package chat

import (
	"strings"
	"testing"

	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/sarvamaierrors"
)

func TestValidateChatRequest(t *testing.T) {
	tests := []struct {
		name      string
		model     string
		messages  []ChatMessage
		wantError error
	}{
		{
			name:      "valid request",
			model:     "sarvam-m",
			messages:  []ChatMessage{{Role: "user", Content: "Hello"}},
			wantError: nil,
		},
		{
			name:      "missing model",
			model:     "",
			messages:  []ChatMessage{{Role: "user", Content: "Hello"}},
			wantError: &sarvamaierrors.ValidationError{Field: "model", Message: "model is required"},
		},
		{
			name:      "missing messages",
			model:     "sarvam-m",
			messages:  nil,
			wantError: &sarvamaierrors.ValidationError{Field: "messages", Message: "messages is required and must contain at least one message"},
		},
		{
			name:      "empty messages",
			model:     "sarvam-m",
			messages:  []ChatMessage{},
			wantError: &sarvamaierrors.ValidationError{Field: "messages", Message: "messages is required and must contain at least one message"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateChatRequest(tt.model, tt.messages)
			if tt.wantError == nil {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
			} else {
				if err == nil {
					t.Errorf("expected error %v, got nil", tt.wantError)
				} else if !strings.Contains(err.Error(), tt.wantError.Error()) {
					t.Errorf("expected error containing %v, got %v", tt.wantError.Error(), err.Error())
				}
			}
		})
	}
}
