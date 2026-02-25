package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Shreehari-Acharya/sarvam-go-sdk"
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
		// Example 1: Detect Hindi in Devanagari script
		// Identifies the language and script of the input sarvamai.
		resp, err := client.Text.DetectLanguage(ctx,
			"मैं ऑफिस जा रहा हूँ",
		)
		if err != nil {
			log.Fatalf("DetectLanguage failed: %v", err)
		}

		fmt.Printf("Detected Language: %s\n", *resp.LanguageCode)
		fmt.Printf("Detected Script: %s\n", *resp.ScriptCode)
		// Output: Language: hi-IN, Script: Deva
	}

	{
		// Example 2: Detect English in Latin script
		resp, err := client.Text.DetectLanguage(ctx,
			"I am going to the office",
		)

		if err != nil {
			log.Fatalf("DetectLanguage failed: %v", err)
		}

		fmt.Printf("Detected Language: %s\n", *resp.LanguageCode)
		fmt.Printf("Detected Script: %s\n", *resp.ScriptCode)
		// Output: Language: en-IN, Script: Latn
	}

	{
		// Example 3: Detect Bengali in Bengali script
		resp, err := client.Text.DetectLanguage(ctx,
			"আমি বাংলায় কথা বলি",
		)
		if err != nil {
			log.Fatalf("DetectLanguage failed: %v", err)
		}

		fmt.Printf("Detected Language: %s\n", *resp.LanguageCode)
		fmt.Printf("Detected Script: %s\n", *resp.ScriptCode)
		// Output: Language: bn-IN, Script: Beng
	}

	{
		// Example 4: Detect Tamil in Tamil script
		resp, err := client.Text.DetectLanguage(ctx,
			"வணக்கம்",
		)
		if err != nil {
			log.Fatalf("DetectLanguage failed: %v", err)
		}

		fmt.Printf("Detected Language: %s\n", *resp.LanguageCode)
		fmt.Printf("Detected Script: %s\n", *resp.ScriptCode)
		// Output: Language: ta-IN, Script: Taml
	}

	{
		// Example 5: Detect Punjabi in Gurmukhi script
		resp, err := client.Text.DetectLanguage(ctx,
			"ਸਤ ਸ੍ਰੀ ਅਕਾਲ",
		)
		if err != nil {
			log.Fatalf("DetectLanguage failed: %v", err)
		}

		fmt.Printf("Detected Language: %s\n", *resp.LanguageCode)
		fmt.Printf("Detected Script: %s\n", *resp.ScriptCode)
		// Output: Language: pa-IN, Script: Guru
	}

	{
		// Example 6: Detect mixed text with Hindi and English
		resp, err := client.Text.DetectLanguage(ctx,
			"मैं ऑफिस जा रहा हूँ। मेरे पास बहुत काम है।",
		)
		if err != nil {
			log.Fatalf("DetectLanguage failed: %v", err)
		}

		fmt.Printf("Detected Language: %s\n", *resp.LanguageCode)
		fmt.Printf("Detected Script: %s\n", *resp.ScriptCode)
		// Output: Language: hi-IN, Script: Deva
	}

	{
		// Example 7: Detect Telugu in Telugu script
		resp, err := client.Text.DetectLanguage(ctx,
			"నమస్కారం",
		)
		if err != nil {
			log.Fatalf("DetectLanguage failed: %v", err)
		}

		fmt.Printf("Detected Language: %s\n", *resp.LanguageCode)
		fmt.Printf("Detected Script: %s\n", *resp.ScriptCode)
		// Output: Language: te-IN, Script: Telu
	}

	_ = ctx
}
