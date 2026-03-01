package sarvamai

import (
	"errors"

	"github.com/Shreehari-Acharya/sarvamai-go/chat"
	"github.com/Shreehari-Acharya/sarvamai-go/docintel"
	"github.com/Shreehari-Acharya/sarvamai-go/internal/transport"
	"github.com/Shreehari-Acharya/sarvamai-go/stt"
	sttjob "github.com/Shreehari-Acharya/sarvamai-go/sttjob"
	"github.com/Shreehari-Acharya/sarvamai-go/text"
	"github.com/Shreehari-Acharya/sarvamai-go/translate"
	translatejob "github.com/Shreehari-Acharya/sarvamai-go/translatejob"
	"github.com/Shreehari-Acharya/sarvamai-go/tts"
)

// Client provides access to Sarvam AI services.
//
// # Services
//
// The client exposes the following services:
//
//   - Text: For translation, transliteration, and language detection
//   - SpeechToText: For speech-to-text transcription (REST and streaming)
//   - TextToSpeech: For text-to-speech conversion
//   - Chat: For conversational AI interactions
//
// # Example
//
//	client, err := sarvamai.NewClient(sarvamai.Config{
//	    APIKey: "your-api-key",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Use text translation
//	resp, err := client.Text.Translate(ctx, translate.Request{...})
//
//	// Use speech-to-text
//	resp, err := client.SpeechToText.Transcribe(ctx, stt.TranscribeRequest{...})
//
//	// Use text-to-speech
//	resp, err := client.TextToSpeech.Convert(ctx, tts.ConvertRequest{...})
//
//	// Use chat
//	resp, err := client.Chat.Completions(ctx, chat.ChatRequest{...})
type Client struct {
	transport *transport.Transport

	Text                     *text.TextClient
	SpeechToText             *stt.STTClient
	SpeechToTextTranslate    *translate.TranslateClient
	SpeechToTextJob          *sttjob.SttJobClient
	SpeechToTextTranslateJob *translatejob.TranslateJobClient
	TextToSpeech             *tts.TTSClient
	Chat                     *chat.ChatClient
	DocumentIntelligence     *docintel.DocIntelClient
}

// NewClient creates a new Sarvam AI client with the given configuration.
//
// # Parameters
//
//	cfg: Configuration containing API key and optional settings. APIKey is required.
//
// # Returns
//
//	A new Client instance or an error if the API key is missing.
//
// # Example
//
//	client, err := sarvamai.NewClient(sarvamai.Config{
//	    APIKey:     "your-api-key",
//	    BaseURL:    "https://api.sarvam.ai",  // optional, default: https://api.sarvam.ai
//	    HTTPClient: &http.Client{Timeout: 30 * time.Second}, // optional
//	})
func NewClient(cfg Config) (*Client, error) {
	if cfg.APIKey == "" {
		return nil, errors.New("SARVAM API key is required. Create one at https://dashboard.sarvam.ai/key-management")
	}

	def := defaultConfig()

	if cfg.BaseURL == "" {
		cfg.BaseURL = def.BaseURL
	}

	if cfg.HTTPClient == nil {
		cfg.HTTPClient = def.HTTPClient
	}

	if cfg.MaxRetries <= 0 {
		cfg.MaxRetries = def.MaxRetries
	}

	t := &transport.Transport{
		APIKey:     cfg.APIKey,
		BaseURL:    cfg.BaseURL,
		HTTPClient: cfg.HTTPClient,
		MaxRetries: cfg.MaxRetries,
	}

	c := &Client{
		transport: t,
	}

	c.Text = text.NewTextClient(t)
	c.SpeechToText = stt.NewSTTClient(t)
	c.SpeechToTextTranslate = translate.NewTranslateClient(t)
	c.SpeechToTextJob = sttjob.NewSttJobClient(t)
	c.SpeechToTextTranslateJob = translatejob.NewTranslateJobClient(t)
	c.TextToSpeech = tts.NewTTSClient(t)
	c.Chat = chat.NewChatClient(t)
	c.DocumentIntelligence = docintel.NewDocIntelClient(t)
	return c, nil
}
