package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Shreehari-Acharya/sarvam-go-sdk"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/chat"
)

func main() {
	ctx := context.Background()

	apiKey := os.Getenv("SARVAM_API_KEY")
	if apiKey == "" {
		log.Fatal("SARVAM_API_KEY environment variable not set")
	}

	client, err := sarvamai.NewClient(sarvamai.Config{
		APIKey: apiKey,
	})
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Example 1: Non-streaming request
	fmt.Println("=== Non-streaming Request ===")
	resp, err := client.Chat.Completions(
		ctx,
		chat.ModelSarvamM,
		[]chat.ChatMessage{
			chat.SystemMessage("You are a helpful assistant."),
			chat.UserMessage("What is the capital of India?"),
		},
		chat.WithMaxTokens(100),
		chat.WithTemperature(0.7),
	)
	if err != nil {
		log.Fatalf("Chat error: %v", err)
	}

	if content, err := resp.FirstChoice(); err == nil {
		fmt.Printf("Assistant: %s\n", content)
	}

	// Example 2: Streaming request
	fmt.Println("\n=== Streaming Request ===")
	stream, err := client.Chat.StreamCompletions(
		ctx,
		chat.ModelSarvamM,
		[]chat.ChatMessage{
			chat.SystemMessage("You are a helpful assistant."),
			chat.UserMessage("write a blog about llms"),
		},
		chat.WithMaxTokens(500),
	)
	if err != nil {
		log.Fatalf("Stream error: %v", err)
	}
	defer stream.Close()

	// Approach 1: Iterate through chunks for real-time output
	fmt.Print("Streaming: ")
	for stream.Next() {
		chunk := stream.Current()
		if len(chunk.Choices) > 0 && chunk.Choices[0].Delta.Content != "" {
			fmt.Print(chunk.Choices[0].Delta.Content)
		}
	}

	if err := stream.Err(); err != nil {
		log.Fatalf("Stream read error: %v", err)
	}

	// Print token usage info (only available in non-streaming)
	fmt.Printf("\n\nTokens used: %d (Prompt: %d, Completion: %d)\n",
		resp.Usage.TotalTokens,
		resp.Usage.PromptTokens,
		resp.Usage.CompletionTokens)
}
