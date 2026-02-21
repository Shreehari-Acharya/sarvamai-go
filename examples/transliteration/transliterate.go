package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Shreehari-Acharya/sarvam-go-sdk"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/text"
)

func main() {
	ctx := context.Background()

	client, err := sarvamai.NewSarvamAIClient(sarvamai.Config{
		APIKey: "your-api-key-here",
	})
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	{
		// Example 1: Basic transliteration (English to Hindi)
		// Converts "Hello" from Latin script to Devanagari script.
		resp, err := client.Text.Transliterate(ctx, text.TransliterateRequest{
			Input:              "Hello",
			SourceLanguageCode: "en-IN",
			TargetLanguageCode: "hi-IN",
		})
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
		resp, err := client.Text.Transliterate(ctx, text.TransliterateRequest{
			Input:              "नमस्ते",
			SourceLanguageCode: "hi-IN",
			TargetLanguageCode: "en-IN",
		})
		if err != nil {
			log.Fatalf("Transliterate failed: %v", err)
		}

		fmt.Printf("Output: %s\n", resp.TransliteratedText)
	}

	{
		// Example 3: Auto-detect source language
		// Automatically detects the source language.
		resp, err := client.Text.Transliterate(ctx, text.TransliterateRequest{
			Input:              "नमस्ते",
			SourceLanguageCode: "auto",
			TargetLanguageCode: "en-IN",
		})

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
		spokenForm := true

		resp, err := client.Text.Transliterate(ctx, text.TransliterateRequest{
			Input:              "मुझे कल 9:30am को appointment है",
			SourceLanguageCode: "hi-IN",
			TargetLanguageCode: "hi-IN",
			SpokenForm:         &spokenForm,
		})
		if err != nil {
			log.Fatalf("Transliterate failed: %v", err)
		}

		fmt.Printf("Output: %s\n", resp.TransliteratedText)
		// Output: "मुझे कल सुबह साढ़े नौ बजे को अपॉइंटमेंट है"
	}

	{
		// Example 5: Transliterate with native numerals
		// Uses language-specific native numerals instead of international (0-9).
		numerals := text.TransliterateNumeralsNative

		resp, err := client.Text.Transliterate(ctx, text.TransliterateRequest{
			Input:              "मेरे पास ₹200 है",
			SourceLanguageCode: "hi-IN",
			TargetLanguageCode: "hi-IN",
			NumeralsFormat:     &numerals,
		})
		if err != nil {
			log.Fatalf("Transliterate failed: %v", err)
		}

		fmt.Printf("Output: %s\n", resp.TransliteratedText)
	}

	{
		// Example 6: Transliterate with spoken form numerals in English
		// Converts numbers to their English spoken form.
		spokenForm := true
		spokenNumerals := text.SpokenFormNumeralsEnglish

		resp, err := client.Text.Transliterate(ctx, text.TransliterateRequest{
			Input:                      "मेरे पास ₹200 है",
			SourceLanguageCode:         "hi-IN",
			TargetLanguageCode:         "hi-IN",
			SpokenForm:                 &spokenForm,
			SpokenFormNumeralsLanguage: &spokenNumerals,
		})
		if err != nil {
			log.Fatalf("Transliterate failed: %v", err)
		}

		fmt.Printf("Output: %s\n", resp.TransliteratedText)
		// Output: "मेरे पास टू हन्डर्ड रूपीस है"
	}

	{
		// Example 7: Transliterate Tamil to English
		// Converts from Tamil script to Latin script.

		resp, err := client.Text.Transliterate(ctx, text.TransliterateRequest{
			Input:              "வணக்கம்",
			SourceLanguageCode: "ta-IN",
			TargetLanguageCode: "en-IN",
		})
		if err != nil {
			log.Fatalf("Transliterate failed: %v", err)
		}

		fmt.Printf("Output: %s\n", resp.TransliteratedText)
		// Output: "vanakkam"
	}

	_ = ctx
}
