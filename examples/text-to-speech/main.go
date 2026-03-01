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

	// Non-streaming TTS using functional options
	resp, err := client.TextToSpeech.Convert(
		ctx,
		"Hello, yeh ek sarvam ai text to speech conversion ka example hai.",
		tts.LanguageHiIN,
		tts.WithSpeakerVoice(tts.SpeakerShubh),
		tts.WithOutputAudioCodec(tts.AudioCodecMP3),
		tts.WithModel(tts.BulbulV3),
	)

	if err != nil {
		log.Fatalf("Convert failed: %v", err)
	}

	for i, audio := range resp.Audios {
		decodedAudio, err := base64.StdEncoding.DecodeString(audio)
		if err != nil {
			log.Fatalf("Failed to decode audio: %v", err)
		}

		err = os.WriteFile(fmt.Sprintf("output-%d.mp3", i), decodedAudio, 0644)
		if err != nil {
			log.Fatalf("Failed to write audio file: %v", err)
		}
	}
}
