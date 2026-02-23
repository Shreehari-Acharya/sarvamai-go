package chat

import (
	"context"

	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/transport"
)

// ChatClient provides access to the Chat Completions API.
type ChatClient struct {
	transport *transport.Transport
}

// NewChatClient creates a new Chat client.
//
// # Parameters
//
//	t: Transport instance configured with API key and base URL
//
// # Returns
//
//	A new ChatClient instance
func NewChatClient(t *transport.Transport) *ChatClient {
	return &ChatClient{
		transport: t,
	}
}

// chatRequest is the internal request structure for the Chat Completions API.
type chatRequest struct {
	Model            string        `json:"model"`
	Messages         []ChatMessage `json:"messages"`
	Temperature      *float64      `json:"temperature,omitempty"`
	TopP             *float64      `json:"top_p,omitempty"`
	ReasoningEffort  *string       `json:"reasoning_effort,omitempty"`
	MaxTokens        *uint32       `json:"max_tokens,omitempty"`
	Stream           *bool         `json:"stream,omitempty"`
	Stop             *[]string     `json:"stop,omitempty"`
	N                *int          `json:"n,omitempty"`
	Seed             *int64        `json:"seed,omitempty"`
	FrequencyPenalty *float64      `json:"frequency_penalty,omitempty"`
	PresencePenalty  *float64      `json:"presence_penalty,omitempty"`
	WikiGrounding    *bool         `json:"wiki_grounding,omitempty"`
}

// Completions generates chat completions for the given conversation.
//
// # Parameters
//
//	ctx: Context for the request (used for cancellation and timeouts)
//	model: The model to use (e.g., "sarvam-m")
//	messages: Array of conversation messages (at least one required)
//	opts: Optional functional options to customize the request
//
// # Returns
//
//	ChatResponse containing the model's reply, or an error
//
// # Functional Options
//
//	WithTemperature(float64)        - Sampling temperature (0.0-2.0)
//	WithTopP(float64)              - Nucleus sampling (0.0-1.0)
//	WithReasoningEffort(ReasoningEffort) - Reasoning effort (low, medium, high)
//	WithMaxTokens(uint32)          - Maximum tokens to generate
//	WithStream(bool)               - Enable streaming responses
//	WithStop([]string)             - Stop sequences
//	WithN(int)                     - Number of responses (1-128)
//	WithSeed(int64)                - Random seed for reproducibility
//	WithFrequencyPenalty(float64)  - Frequency penalty (-2.0 to 2.0)
//	WithPresencePenalty(float64)   - Presence penalty (-2.0 to 2.0)
//	WithWikiGrounding(bool)        - Enable Wikipedia grounding
//
// # Example (using helper functions)
//
//	resp, err := client.Completions(
//	    context.Background(),
//	    chat.ModelSarvamM,
//	    []chat.ChatMessage{
//	        chat.SystemMessage("You are a helpful assistant."),
//	        chat.UserMessage("What is the capital of India?"),
//	    },
//	    chat.WithTemperature(0.7),
//	    chat.WithMaxTokens(500),
//	)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Or use MessageBuilder for more control
//	msg := chat.NewMessageBuilder().
//	    AsUser().
//	    WithContent("Hello").
//	    Build()
//
//	for _, choice := range resp.Choices {
//	    fmt.Println(choice.Message.Content)
//	}
func (c *ChatClient) Completions(
	ctx context.Context,
	model string,
	messages []ChatMessage,
	opts ...option,
) (*ChatResponse, error) {

	req := &chatRequest{
		Model:    model,
		Messages: messages,
	}

	for _, opt := range opts {

		if err := opt(req); err != nil {
			return nil, err
		}
	}

	var resp ChatResponse

	err := c.transport.DoRequest(
		ctx,
		"POST",
		"/v1/chat/completions",
		req,
		&resp,
		"application/json",
	)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// StreamCompletions generates a streaming chat completion for the given conversation.
//
// Use this method when you want to receive the response as it's being generated,
// character by character, for a real-time experience.
//
// # Parameters
//
//	ctx: Context for the request (used for cancellation and timeouts)
//	model: The model to use (e.g., "sarvam-m")
//	messages: Array of conversation messages (at least one required)
//	opts: Optional functional options to customize the request
//
// # Returns
//
//	*ChatStream that can be iterated to get chunks, or an error
//
// # Example
//
//	stream, err := client.StreamCompletions(
//	    context.Background(),
//	    chat.ModelSarvamM,
//	    []chat.ChatMessage{
//	        chat.SystemMessage("You are a helpful assistant."),
//	        chat.UserMessage("What is the capital of India?"),
//	    },
//	    chat.WithTemperature(0.7),
//	    chat.WithMaxTokens(500),
//	)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer stream.Close()
//
//	// Iterate through chunks
//	for stream.Next() {
//	    chunk := stream.Current()
//	    if len(chunk.Choices) > 0 {
//	        fmt.Print(chunk.Choices[0].Delta.Content)
//	    }
//	}
//
//	if err := stream.Err(); err != nil {
//	    log.Fatal(err)
//	}
//
//	// Or get all accumulated text at once
//	fmt.Println(stream.Text())
func (c *ChatClient) StreamCompletions(
	ctx context.Context,
	model string,
	messages []ChatMessage,
	opts ...option,
) (*ChatStream, error) {

	// Ensure the stream option is set to true
	streamOpt := true

	req := &chatRequest{
		Model:    model,
		Messages: messages,
		Stream:   &streamOpt,
	}

	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, err
		}
	}

	resp, err := c.transport.DoStreamRequest(
		ctx,
		"POST",
		"/v1/chat/completions",
		req,
	)
	if err != nil {
		return nil, err
	}

	return NewChatStream(resp), nil
}
