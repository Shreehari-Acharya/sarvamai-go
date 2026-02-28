package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Shreehari-Acharya/sarvam-go-sdk"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/shared/speech"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/stt"
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

	audioFile, err := os.Open("examples/speech-to-text/sample-audio.wav")
	if err != nil {
		log.Fatalf("Failed to open audio file: %v", err)
	}

	defer audioFile.Close()

	resp, err := client.SpeechToText.Transcribe(
		ctx,
		audioFile,
		stt.WithModel(speech.ModelSaaras),
		stt.WithMode(speech.ModeTranslate),
	)
	if err != nil {
		log.Fatalf("Transcription failed: %v", err)
	}

	fmt.Printf("Transcribed Text: %s\n", resp.Transcript)
	fmt.Printf("Language Detected: %s\n", *resp.LanguageCode)
	fmt.Printf("Request ID: %s\n", *resp.RequestID)
}
