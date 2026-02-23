// Package chat provides types for the Chat Completions API.
package chat

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
)

// Role represents the author of a message in a conversation.
//
// # Values
//
//   - System:    System messages that set the assistant's behavior and context
//   - User:     User messages containing requests or questions
//   - Assistant: Assistant messages representing the model's responses (for conversation history)
type Role string

const (
	RoleSystem    Role = "system"
	RoleUser      Role = "user"
	RoleAssistant Role = "assistant"
)

// ChatMessage represents a single turn in the conversation.
//
// Use [SystemMessage], [UserMessage], or [AssistantMessage] helpers to create instances.
type ChatMessage struct {
	Role    Role   `json:"role"`
	Content string `json:"content"`
}

// Role returns the role of the message.
func (m ChatMessage) RoleValue() Role {
	return m.Role
}

// SystemMessage creates a system message that sets the assistant's behavior.
func SystemMessage(content string) ChatMessage {
	return ChatMessage{Role: RoleSystem, Content: content}
}

// UserMessage creates a user message containing a request or question.
func UserMessage(content string) ChatMessage {
	return ChatMessage{Role: RoleUser, Content: content}
}

// AssistantMessage creates an assistant message (for conversation history).
func AssistantMessage(content string) ChatMessage {
	return ChatMessage{Role: RoleAssistant, Content: content}
}

// Model identifiers for chat completions.
const (
	ModelSarvamM = "sarvam-m"
)

// ReasoningEffort controls how much reasoning the model applies before generating a response.
type ReasoningEffort string

const (
	ReasoningEffortLow    ReasoningEffort = "low"
	ReasoningEffortMedium ReasoningEffort = "medium"
	ReasoningEffortHigh   ReasoningEffort = "high"
)

// FinishReason represents why the model stopped generating a response.
type FinishReason string

const (
	FinishReasonStop   FinishReason = "stop"
	FinishReasonLength FinishReason = "length"
	FinishReasonError  FinishReason = "error"
)

// Choice represents a single response choice from the model.
type Choice struct {
	Index        int          `json:"index"`
	Message      ChatMessage  `json:"message"`
	FinishReason FinishReason `json:"finish_reason"`
}

// Usage represents token usage statistics for the request and response.
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// FirstChoice returns the content of the first choice, if available.
func (r *ChatResponse) FirstChoice() (string, error) {
	if len(r.Choices) == 0 {
		return "", fmt.Errorf("no choices in response")
	}
	return r.Choices[0].Message.Content, nil
}

// ChatResponse represents a chat completions response from the API.
type ChatResponse struct {
	ID      string   `json:"id"`
	Choices []Choice `json:"choices"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Usage   Usage    `json:"usage"`
}

// ChatDelta represents the incremental content in a streaming response.
type ChatDelta struct {
	Role    Role   `json:"role,omitempty"`
	Content string `json:"content,omitempty"`
}

// ChunkChoice represents a single choice in a streaming response.
type ChunkChoice struct {
	Index        int          `json:"index"`
	Delta        ChatDelta    `json:"delta"`
	FinishReason FinishReason `json:"finish_reason"`
}

// ChatChunk represents a single chunk in a streaming response.
type ChatChunk struct {
	ID      string        `json:"id"`
	Choices []ChunkChoice `json:"choices"`
	Created int64         `json:"created"`
	Model   string        `json:"model"`
}

// ChatStream is an iterator for streaming chat responses.
type ChatStream struct {
	mu          sync.Mutex
	reader      *bufio.Reader
	resp        *http.Response
	current     ChatChunk
	done        bool
	err         error
	accumulated strings.Builder
}

// NewChatStream creates a new ChatStream from an HTTP response.
func NewChatStream(resp *http.Response) *ChatStream {
	return &ChatStream{
		reader: bufio.NewReader(resp.Body),
		resp:   resp,
	}
}

// Next advances the iterator to the next chunk.
// Returns true if there is a chunk available, false if the stream is done or errored.
func (s *ChatStream) Next() bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.done || s.err != nil {
		return false
	}

	for {
		line, err := s.reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				s.done = true
			} else {
				s.err = fmt.Errorf("read stream: %w", err)
			}
			return false
		}

		line = strings.TrimSpace(line)
		if line == "" || !strings.HasPrefix(line, "data: ") {
			continue
		}

		data := strings.TrimPrefix(line, "data: ")
		if data == "[DONE]" {
			s.done = true
			return false
		}

		var chunk ChatChunk
		if err := json.Unmarshal([]byte(data), &chunk); err != nil {
			s.err = fmt.Errorf("unmarshal chunk: %w", err)
			return false
		}

		s.current = chunk
		if len(chunk.Choices) > 0 && chunk.Choices[0].Delta.Content != "" {
			s.accumulated.WriteString(chunk.Choices[0].Delta.Content)
		}
		return true
	}
}

// Current returns the current chunk.
// Valid only after Next returns true.
func (s *ChatStream) Current() ChatChunk {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.current
}

// Text returns all accumulated content as a string.
func (s *ChatStream) Text() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.accumulated.String()
}

// Err returns the error encountered during streaming, if any.
func (s *ChatStream) Err() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.err
}

// Close closes the stream and releases resources.
func (s *ChatStream) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.resp != nil && s.resp.Body != nil {
		return s.resp.Body.Close()
	}
	return nil
}

// Choice returns the first choice from the current chunk.
func (s *ChatStream) Choice() (ChunkChoice, bool) {
	chunk := s.Current()
	if len(chunk.Choices) == 0 {
		return ChunkChoice{}, false
	}
	return chunk.Choices[0], true
}
