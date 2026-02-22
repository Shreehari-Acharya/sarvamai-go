package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	sarvamai "github.com/Shreehari-Acharya/sarvam-go-sdk"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/stt"
)

func main() {

	client, err := sarvamai.NewSarvamAIClient(sarvamai.Config{
		APIKey: "your-api-key-here",
	})
	if err != nil {
		log.Fatal("client error:", err)
	}

	stream, err := client.SpeechToText.TranscribeStream(context.Background(), stt.StreamConfig{
		SampleRate: 16000,
	})
	if err != nil {
		log.Fatal("stream error:", err)
	}
	defer stream.Close()

	// Handle errors from the stream in background
	go func() {
		for err := range stream.Errors() {
			log.Println("stream error:", err)
		}
	}()

	// Print messages as they arrive
	go func() {
		for msg := range stream.Messages() {
			fmt.Printf("[%s] %s\n", msg.Type, msg.Data)
		}
	}()

	// Open and stream audio file
	file, err := os.Open("examples/speech-to-text/sample-audio.wav")
	if err != nil {
		log.Fatal("file error:", err)
	}
	defer file.Close()

	buf := make([]byte, 3200)
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("read error:", err)
		}
		if err := stream.SendAudio(buf[:n]); err != nil {
			log.Fatal("send error:", err)
		}
		time.Sleep(100 * time.Millisecond)
	}

	// Flush to finalize transcription
	if err := stream.Flush(); err != nil {
		log.Fatal("flush error:", err)
	}

	// Wait for final response
	time.Sleep(2 * time.Second)

	fmt.Println("\n--- Full Transcript ---")
	fmt.Println(stream.Transcript())
}
