package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/Shreehari-Acharya/sarvam-go-sdk"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/shared/speech"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/translate"
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
		log.Fatal("client error:", err)
	}

	stream, err := client.SpeechToTextTranslate.TranslateStream(
		ctx,
		translate.WithSampleRateForTranslateStream(speech.SampleRate16000),
	)
	if err != nil {
		log.Fatal("stream error:", err)
	}
	defer stream.Close()

	// Open audio file
	file, err := os.Open("examples/speech-to-text/sample-audio.wav")
	if err != nil {
		log.Fatal("file error:", err)
	}
	defer file.Close()

	// Send audio in a goroutine
	go func() {
		buf := make([]byte, 32000)
		for {
			n, err := file.Read(buf)
			if err == io.EOF {
				stream.Flush()
				break
			}
			if err != nil {
				log.Printf("read error: %v", err)
				return
			}
			if err := stream.SendAudio(buf[:n]); err != nil {
				log.Printf("send error: %v", err)
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Iterate through responses using the iterator pattern
	for stream.Next() {
		resp := stream.Current()
		fmt.Printf("[%s] %s\n", resp.Type, resp.Data)
	}

	// Check for errors
	if err := stream.Err(); err != nil {
		log.Printf("stream error: %v", err)
	}

	fmt.Println("\n--- Full Transcript ---")
	fmt.Println(stream.Text())
}
