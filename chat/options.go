package chat

import (
	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/sarvamaierrors"
)

//
// Option Type and Functions
//

// option is a functional option that modifies a chatRequest.
type option func(*chatRequest) error

// WithTemperature sets the sampling temperature for the chat request.
//
// Temperature controls the randomness of the model's output.
// Higher values (e.g., 0.8) make output more random, lower values (e.g., 0.2) make it more focused.
//
// Range: 0.0 to 2.0
// Default: 0.2
//
// Note: It is generally recommended to set either temperature or top_p, but not both.
func WithTemperature(temp float64) option {
	return func(r *chatRequest) error {
		if temp < 0.0 || temp > 2.0 {
			return &sarvamaierrors.ValidationError{
				Message: "temperature must be between 0.0 and 2.0",
				Field:   "temperature",
			}
		}
		r.Temperature = &temp
		return nil
	}
}

// WithTopP sets the nucleus sampling probability for the chat request.
//
// TopP is an alternative to temperature that controls the diversity of the model's output.
// It considers the smallest set of tokens whose cumulative probability exceeds top_p.
//
// Range: 0.0 to 1.0
// Default: 1.0
//
// Note: It is generally recommended to set either temperature or top_p, but not both.
func WithTopP(topP float64) option {
	return func(r *chatRequest) error {
		if topP < 0.0 || topP > 1.0 {
			return &sarvamaierrors.ValidationError{
				Message: "top_p must be between 0.0 and 1.0",
				Field:   "top_p",
			}
		}
		r.TopP = &topP
		return nil
	}
}

var allowedReasoningEffortValues = map[ReasoningEffort]bool{
	ReasoningEffortLow:    true,
	ReasoningEffortMedium: true,
	ReasoningEffortHigh:   true,
}

// WithReasoningEffort sets the reasoning effort level for the chat request.
//
// Effort controls how much reasoning the model applies before generating a response.
// Higher effort may produce more accurate responses but can increase latency and cost.
//
//   - Low:    Fast response, minimal reasoning
//   - Medium: Balanced (default)
//   - High:   Thorough reasoning, may take longer
func WithReasoningEffort(effort ReasoningEffort) option {
	return func(r *chatRequest) error {
		if !allowedReasoningEffortValues[effort] {
			return &sarvamaierrors.ValidationError{
				Message: "reasoning_effort must be one of 'low', 'medium', or 'high'",
				Field:   "reasoning_effort",
			}
		}
		e := string(effort)
		r.ReasoningEffort = &e
		return nil
	}
}

// WithMaxTokens sets the maximum number of tokens to generate in the chat response.
//
// The maximum length of the model's response in tokens.
// If not set, the model will use its default token limit.
func WithMaxTokens(maxTokens uint32) option {
	return func(r *chatRequest) error {
		r.MaxTokens = &maxTokens
		return nil
	}
}

// WithStop sets the stop sequences for the chat request.
//
// The model will stop generating further tokens when any of the specified sequences are generated.
// This can be used to control the end of the response.
func WithStop(stop []string) option {
	return func(r *chatRequest) error {
		r.Stop = &stop
		return nil
	}
}

// WithN sets the number of chat completion choices to generate.
//
// Number of chat completion choices to generate for each input message.
// The API will return n different choices. You can then pick the one that best suits your needs.
//
// Range: 1 to 128
// Default: 1
func WithN(n int) option {
	return func(r *chatRequest) error {
		if n < 1 || n > 128 {
			return &sarvamaierrors.ValidationError{
				Message: "n must be between 1 and 128",
				Field:   "n",
			}
		}
		r.N = &n
		return nil
	}
}

// WithSeed sets the random seed for reproducibility.
//
// This feature is in Beta. If specified, the system will make a best effort to
// sample deterministically, such that repeated requests with the same seed and
// parameters should return the same result.
//
// Note: Only effective when temperature > 0.
func WithSeed(seed int64) option {
	return func(r *chatRequest) error {
		r.Seed = &seed
		return nil
	}
}

// WithFrequencyPenalty sets the frequency penalty for the chat request.
//
// Positive values penalize new tokens based on their existing frequency in the text so far,
// decreasing the model's likelihood to repeat the same line verbatim.
//
// Range: -2.0 to 2.0
// Default: 0.0
func WithFrequencyPenalty(penalty float64) option {
	return func(r *chatRequest) error {
		if penalty < -2.0 || penalty > 2.0 {
			return &sarvamaierrors.ValidationError{
				Message: "frequency_penalty must be between -2.0 and 2.0",
				Field:   "frequency_penalty",
			}
		}
		r.FrequencyPenalty = &penalty
		return nil
	}
}

// WithPresencePenalty sets the presence penalty for the chat request.
//
// Positive values penalize new tokens based on whether they appear in the text so far,
// increasing the model's likelihood to talk about new topics.
//
// Range: -2.0 to 2.0
// Default: 0.0
func WithPresencePenalty(penalty float64) option {
	return func(r *chatRequest) error {
		if penalty < -2.0 || penalty > 2.0 {
			return &sarvamaierrors.ValidationError{
				Message: "presence_penalty must be between -2.0 and 2.0",
				Field:   "presence_penalty",
			}
		}
		r.PresencePenalty = &penalty
		return nil
	}
}

// WithWikiGrounding enables or disables Wikipedia grounding for the chat request.
//
// When enabled, the model will use information from Wikipedia to generate
// more accurate and informative responses.
// Default: false
func WithWikiGrounding(wikiGrounding bool) option {
	return func(r *chatRequest) error {
		r.WikiGrounding = &wikiGrounding
		return nil
	}
}
