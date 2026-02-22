package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Shreehari-Acharya/sarvam-go-sdk"
)

func main() {
	ctx := context.Background()

	client, err := sarvamai.NewClient(sarvamai.Config{
		APIKey: "your-api-key-here",
	})
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	{
		// Example 1: Basic translation (English to Hindi)
		// Translates "Hello" from English to Hindi using default mayura:v1 model.
		resp, err := client.Text.Translate(ctx, sarvamai.TranslateRequest{
			Input:              "Hello",
			SourceLanguageCode: "en-IN",
			TargetLanguageCode: "hi-IN",
		})
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
		resp, err := client.Text.Translate(ctx, sarvamai.TranslateRequest{
			Input:              "मैं ऑफिस जा रहा हूँ",
			SourceLanguageCode: "auto",
			TargetLanguageCode: "en-IN",
		})
		if err != nil {
			log.Fatalf("Translate failed: %v", err)
		}

		fmt.Printf("Output: %s\n", resp.TranslatedText)
		fmt.Printf("Detected Source Language: %s\n", resp.SourceLanguageCode)
	}

	{
		// Example 3: Using sarvam-translate:v1 model with additional languages
		// This model supports 22 Indian languages but only formal mode.
		model := sarvamai.ModelSarvamTranslate
		resp, err := client.Text.Translate(ctx, sarvamai.TranslateRequest{
			Input:              "Hello world",
			SourceLanguageCode: "en-IN",
			TargetLanguageCode: "ta-IN", // Tamil
			Model:              &model,
		})
		if err != nil {
			log.Fatalf("Translate failed: %v", err)
		}

		fmt.Printf("Output: %s\n", resp.TranslatedText)
	}

	{
		// Example 4: Translate with output script (transliteration)
		// OutputScript is only supported by mayura:v1 model.
		mode := sarvamai.ModeFormal
		outputScript := sarvamai.OutputScriptRoman

		resp, err := client.Text.Translate(ctx, sarvamai.TranslateRequest{
			Input:              "Your EMI of Rs. 3000 is pending",
			SourceLanguageCode: "en-IN",
			TargetLanguageCode: "hi-IN",
			Mode:               &mode,
			OutputScript:       &outputScript,
		})
		if err != nil {
			log.Fatalf("Translate failed: %v", err)
		}

		fmt.Printf("Output: %s\n", resp.TranslatedText)
	}

	{
		// Example 5: Translate with native numerals
		// Uses language-specific native numerals instead of international (0-9).
		numerals := sarvamai.NumeralsNative

		resp, err := client.Text.Translate(ctx, sarvamai.TranslateRequest{
			Input:              "My phone number is: 9840950950",
			SourceLanguageCode: "en-IN",
			TargetLanguageCode: "hi-IN",
			NumeralsFormat:     &numerals,
		})
		if err != nil {
			log.Fatalf("Translate failed: %v", err)
		}

		fmt.Printf("Output: %s\n", resp.TranslatedText)
	}

	{
		// Example 6: Translate with speaker gender
		// Influences the translation style based on speaker gender.
		gender := sarvamai.GenderFemale
		req := sarvamai.TranslateRequest{
			Input:              "I am going to the office",
			SourceLanguageCode: "en-IN",
			TargetLanguageCode: "hi-IN",
			SpeakerGender:      &gender,
		}

		resp, err := client.Text.Translate(ctx, req)
		if err != nil {
			log.Fatalf("Translate failed: %v", err)
		}

		fmt.Printf("Input: %s\n", req.Input)
		fmt.Printf("Output: %s\n", resp.TranslatedText)
	}

	_ = ctx
}
