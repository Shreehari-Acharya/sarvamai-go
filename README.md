# Sarvam AI Go SDK (Unofficial)

[![Go Reference](https://pkg.go.dev/badge/github.com/Shreehari-Acharya/sarvamai-go.svg)](https://pkg.go.dev/github.com/Shreehari-Acharya/sarvamai-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/Shreehari-Acharya/sarvamai-go)](https://goreportcard.com/report/github.com/Shreehari-Acharya/sarvamai-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

An idiomatic Go client for the Sarvam AI API.

Note: This is an unofficial, community-maintained SDK. It is not affiliated with or endorsed by Sarvam AI.

## Features

- **Chat**: Conversational AI optimized for Indic contexts.
- **Speech-to-Text**: Support for both REST and WebSocket-based real-time transcription.
- **Text-to-Speech**: Natural synthesis with streaming capabilities.
- **Text Processing**: Translation, Transliteration, and Language Identification (LID).
- **Document Intelligence**: Digitization and extraction for complex documents.
- **Batch Processing**: Efficient handling of large audio volumes via Job APIs.
- **Resilience**: Automatic retries with exponential backoff and context-aware connection management.

## Installation

```bash
go get github.com/Shreehari-Acharya/sarvamai-go
```

## Quick Start

### Initialization

Initialize the client using your API key from the Sarvam dashboard.

```go
import "github.com/Shreehari-Acharya/sarvamai-go"

client, err := sarvamai.NewClient(sarvamai.Config{
    APIKey: "your-api-key",
})
```

### Chat Completion

```go
import "github.com/Shreehari-Acharya/sarvamai-go/chat"

resp, err := client.Chat.Completions(ctx, chat.ModelSarvamM, []chat.ChatMessage{
    chat.SystemMessage("You are a helpful assistant."),
    chat.UserMessage("What is the capital of India?"),
})
```

### Speech-to-Text

```go
import "github.com/Shreehari-Acharya/sarvamai-go/stt"

file, _ := os.Open("audio.wav")
defer file.Close()

resp, err := client.SpeechToText.Transcribe(ctx, file,
    stt.WithModel(stt.ModelSaaras),
    stt.WithLanguage(stt.LanguageHiIN),
)
```

### Text-to-Speech Streaming

```go
import "github.com/Shreehari-Acharya/sarvamai-go/tts"

stream, err := client.TextToSpeech.StreamConvert(ctx, tts.LanguageHiIN,
    tts.WithStreamSpeaker(tts.SpeakerShubh),
)
defer stream.Close()

stream.SendText("नमस्ते, आप कैसे हैं?")
stream.Flush()

for stream.Next() {
    chunk := stream.Current()
    // chunk.Audio contains base64 encoded data
}
```

## Resilience and Reliability

- **Retries**: The SDK automatically retries on `429 Too Many Requests` and `5xx` errors.
- **Timeout Management**: REST requests use a default 30-second timeout, configurable via the client config.
- **Resource Cleanup**: WebSocket streams respect context cancellation to ensure underlying connections and goroutines are properly terminated.

## Documentation And Examples

Full documentation and detailed examples are available in the [examples/](./examples) directory.

## Contributing

Contributions are welcome. Please refer to [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
