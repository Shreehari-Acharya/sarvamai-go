package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Shreehari-Acharya/sarvam-go-sdk"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/text"
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

	{
		// Example 1: Basic translation (English to Hindi)
		// Translates "Hello" from English to Hindi using default mayura:v1 model.
		resp, err := client.Text.Translate(ctx, "Hello", "en-IN", "hi-IN")
		if err != nil {
			log.Fatalf("Translate failed: %v", err)
		}

		fmt.Printf("Output: %s\n", resp.TranslatedText)
		fmt.Printf("Detected Source Language: %s\n", resp.SourceLanguageCode)
		fmt.Printf("Request ID: %s\n", *resp.RequestID)
	}

	{
		// Example 2: Auto-detect source language
		// Uses "auto" to automatically detect source language (mayura:v1 only).
		resp, err := client.Text.Translate(ctx, "मैं ऑफिस जा रहा हूँ", "auto", "en-IN")
		if err != nil {
			log.Fatalf("Translate failed: %v", err)
		}

		fmt.Printf("Output: %s\n", resp.TranslatedText)
		fmt.Printf("Detected Source Language: %s\n", resp.SourceLanguageCode)
	}

	{
		// Example 3: Using sarvam-translate:v1 model with additional languages
		// This model supports 22 Indian languages but only formal mode.
		resp, err := client.Text.Translate(ctx, "Hello world", "en-IN", "ta-IN",
			text.WithModel(text.ModelSarvamTranslate),
		)
		if err != nil {
			log.Fatalf("Translate failed: %v", err)
		}

		fmt.Printf("Output: %s\n", resp.TranslatedText)
	}

	{
		// Example 4: Translate with output script (transliteration)
		// OutputScript is only supported by mayura:v1 model.
		resp, err := client.Text.Translate(ctx, "Your EMI of Rs. 3000 is pending", "en-IN", "hi-IN",
			text.WithModel(text.ModelMayura),
			text.WithMode(text.ModeFormal),
			text.WithOutputScript(text.OutputScriptRoman),
		)
		if err != nil {
			log.Fatalf("Translate failed: %v", err)
		}

		fmt.Printf("Output: %s\n", resp.TranslatedText)
	}

	{
		// Example 5: Translate with native numerals
		// Uses language-specific native numerals instead of international (0-9).
		resp, err := client.Text.Translate(ctx, "My phone number is: 9840950950", "en-IN", "hi-IN",
			text.WithNumeralsFormat(text.NumeralsNative),
		)
		if err != nil {
			log.Fatalf("Translate failed: %v", err)
		}

		fmt.Printf("Output: %s\n", resp.TranslatedText)
	}

	{
		// Example 6: Translate with speaker gender
		// Influences the translation style based on speaker gender.
		resp, err := client.Text.Translate(ctx, "I am going to the office", "en-IN", "hi-IN",
			text.WithSpeakerGender(text.GenderFemale),
		)
		if err != nil {
			log.Fatalf("Translate failed: %v", err)
		}

		fmt.Printf("Output: %s\n", resp.TranslatedText)
	}

	_ = ctx
}
