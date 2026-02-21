package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Shreehari-Acharya/sarvam-go-sdk"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/stt"
)

func main() {
	ctx := context.Background()

	client, err := sarvamai.NewSarvamAIClient(sarvamai.Config{
		APIKey: "your-api-key-here",
	})

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	audioFile, err := os.Open("examples/speech-to-text/sample-audio.wav")
	if err != nil {
		log.Fatalf("Failed to open audio file: %v", err)
	}

	defer audioFile.Close()

	model := stt.ModelSaaras
	mode := stt.ModeTranslit
	resp, err := client.SpeechToText.Transcribe(ctx, stt.TranscribeRequest{
		File:     audioFile,
		FileName: "sample-audio.wav",
		Model:    &model,
		Mode:     &mode,
	})
	if err != nil {
		log.Fatalf("Transcription failed: %v", err)
	}

	fmt.Printf("Transcribed Text: %s\n", resp.Transcript)
	fmt.Printf("Language Detected: %s\n", *resp.LanguageCode)
	fmt.Printf("Request ID: %s\n", *resp.RequestID)
}
