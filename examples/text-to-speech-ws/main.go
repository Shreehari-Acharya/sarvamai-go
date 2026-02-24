package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/Shreehari-Acharya/sarvam-go-sdk"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/tts"
)

func main() {
	apiKey := os.Getenv("SARVAM_API_KEY")
	if apiKey == "" {
		log.Fatal("SARVAM_API_KEY environment variable not set")
	}

	client, err := sarvamai.NewClient(sarvamai.Config{
		APIKey: apiKey,
	})
	if err != nil {
		log.Fatal("failed to initialize client:", err)
	}

	ctx := context.Background()

	// Streaming TTS using iterator pattern
	stream, err := client.TextToSpeech.StreamConvert(
		ctx,
		"hi-IN",
		tts.WithStreamSpeaker(tts.SpeakerShubh),
		tts.WithStreamModel(tts.BulbulV3Beta),
		tts.WithStreamTemperature(0.7),
	)
	if err != nil {
		log.Fatal("failed to open stream:", err)
	}
	defer stream.Close()

	// Send text chunks
	textParts := []string{
		"नमस्ते!",
		"मैं सर्वम ए-आई का बुलबुल मॉडल हूँ।",
		"आज मैं आपको टेक्स्ट-टू-स्पीच स्ट्रीमिंग दिखा रहा हूँ।",
	}

	for _, part := range textParts {
		fmt.Println("Sending:", part)
		if err := stream.SendText(part); err != nil {
			log.Fatal("send error:", err)
		}
	}

	if err := stream.Flush(); err != nil {
		log.Fatal("flush error:", err)
	}

	// Save audio chunks to file
	f, err := os.Create("output.mp3")
	if err != nil {
		log.Fatal("file error:", err)
	}
	defer f.Close()

	fmt.Println("Receiving audio...")
	for stream.Next() {
		chunk := stream.Current()
		decoded, err := base64.StdEncoding.DecodeString(chunk.Audio)
		if err != nil {
			log.Println("decode error:", err)
			continue
		}
		f.Write(decoded)
		fmt.Printf("Received chunk: %d bytes\n", len(decoded))
	}

	if err := stream.Err(); err != nil {
		log.Fatal("stream error:", err)
	}

	fmt.Println("Done. Saved to output.mp3")
}
