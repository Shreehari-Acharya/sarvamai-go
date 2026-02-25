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
		// Example 1: Basic transliteration (English to Hindi)
		// Converts "Hello" from Latin script to Devanagari script.
		resp, err := client.Text.Transliterate(ctx, "Hello", "en-IN", "hi-IN")
		if err != nil {
			log.Fatalf("Transliterate failed: %v", err)
		}

		fmt.Printf("Output: %s\n", resp.TransliteratedText)
		fmt.Printf("Detected Source Language: %s\n", resp.SourceLanguageCode)
		fmt.Printf("Request ID: %s\n", *resp.RequestID)
	}

	{
		// Example 2: Transliterate Hindi to English
		// Converts from Devanagari script to Latin script.
		resp, err := client.Text.Transliterate(ctx, "नमस्ते", "hi-IN", "en-IN")
		if err != nil {
			log.Fatalf("Transliterate failed: %v", err)
		}

		fmt.Printf("Output: %s\n", resp.TransliteratedText)
	}

	{
		// Example 3: Auto-detect source language
		// Automatically detects the source language.
		resp, err := client.Text.Transliterate(ctx, "नमस्ते", "auto", "en-IN")

		if err != nil {
			log.Fatalf("Transliterate failed: %v", err)
		}

		fmt.Printf("Output: %s\n", resp.TransliteratedText)
		fmt.Printf("Detected Source Language: %s\n", resp.SourceLanguageCode)
	}

	{
		// Example 4: Transliterate with spoken form
		// Converts text to natural spoken form.
		// Note: Has no effect when target language is en-IN.
		resp, err := client.Text.Transliterate(ctx, "मुझे कल 9:30am को appointment है", "hi-IN", "hi-IN",
			text.WithSpokenForm(true),
		)
		if err != nil {
			log.Fatalf("Transliterate failed: %v", err)
		}

		fmt.Printf("Output: %s\n", resp.TransliteratedText)
		// Output: "मुझे कल सुबह साढ़े नौ बजे को अपॉइंटमेंट है"
	}

	{
		// Example 5: Transliterate with native numerals
		// Uses language-specific native numerals instead of international (0-9).
		resp, err := client.Text.Transliterate(ctx, "मेरे पास ₹200 है", "hi-IN", "hi-IN",
			text.WithNumeralsFormatTransliteration(text.NumeralsNative),
		)
		if err != nil {
			log.Fatalf("Transliterate failed: %v", err)
		}

		fmt.Printf("Output: %s\n", resp.TransliteratedText)
	}

	{
		// Example 6: Transliterate with spoken form numerals in English
		// Converts numbers to their English spoken form.
		resp, err := client.Text.Transliterate(ctx, "मेरे पास ₹200 है", "hi-IN", "hi-IN",
			text.WithSpokenForm(true),
			text.WithSpokenFormNumeralsLanguage(text.SpokenFormNumeralsEnglish),
		)
		if err != nil {
			log.Fatalf("Transliterate failed: %v", err)
		}

		fmt.Printf("Output: %s\n", resp.TransliteratedText)
		// Output: "मेरे पास टू हन्डर्ड रूपीस है"
	}

	{
		// Example 7: Transliterate Tamil to English
		// Converts from Tamil script to Latin script.

		resp, err := client.Text.Transliterate(ctx, "வணக்கம்", "ta-IN", "en-IN")
		if err != nil {
			log.Fatalf("Transliterate failed: %v", err)
		}

		fmt.Printf("Output: %s\n", resp.TransliteratedText)
		// Output: "vanakkam"
	}

	_ = ctx
}
