package text

import (
	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/transport"
)

// Service provides methods for interacting with Sarvam AI's text-related APIs.
type Service struct {
	transport *transport.Transport
}

// NewService creates a new Service instance with the given Sarvam AI requester.
func NewService(t *transport.Transport) *Service {
	return &Service{transport: t}
}
