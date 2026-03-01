package chat

//
// Option Type and Functions
//

// option is a functional option that modifies a chatRequest.
type option func(*chatRequest)

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
	return func(r *chatRequest) {
		r.Temperature = &temp
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
	return func(r *chatRequest) {
		r.TopP = &topP
	}
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
	return func(r *chatRequest) {
		e := string(effort)
		r.ReasoningEffort = &e
	}
}

// WithMaxTokens sets the maximum number of tokens to generate in the chat response.
//
// The maximum length of the model's response in tokens.
// If not set, the model will use its default token limit.
func WithMaxTokens(maxTokens uint32) option {
	return func(r *chatRequest) {
		r.MaxTokens = &maxTokens
	}
}

// WithStop sets the stop sequences for the chat request.
//
// The model will stop generating further tokens when any of the specified sequences are generated.
// This can be used to control the end of the response.
func WithStop(stop []string) option {
	return func(r *chatRequest) {
		r.Stop = &stop
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
	return func(r *chatRequest) {
		r.N = &n
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
	return func(r *chatRequest) {
		r.Seed = &seed
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
	return func(r *chatRequest) {
		r.FrequencyPenalty = &penalty
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
	return func(r *chatRequest) {
		r.PresencePenalty = &penalty
	}
}

// WithWikiGrounding enables or disables Wikipedia grounding for the chat request.
//
// When enabled, the model will use information from Wikipedia to generate
// more accurate and informative responses.
// Default: false
func WithWikiGrounding(wikiGrounding bool) option {
	return func(r *chatRequest) {
		r.WikiGrounding = &wikiGrounding
	}
}
