package tts

import (
	"fmt"
	"slices"

	"github.com/Shreehari-Acharya/sarvam-go-sdk/internal/sarvamaierrors"
	"github.com/Shreehari-Acharya/sarvam-go-sdk/languages"
)

func validateTargetLanguage(code string) error {
	if code == "" || !languages.TargetLanguages[languages.Code(code)] {
		return &sarvamaierrors.ValidationError{
			Field:   "target_language_code",
			Message: "target_language_code is required and must be a supported language code",
		}
	}
	return nil
}

func validateText(text string, model Model) error {
	maxLen := 2500
	if model == BulbulV2 {
		maxLen = 1500
	}
	if len(text) == 0 || len(text) > maxLen {
		return &sarvamaierrors.ValidationError{
			Field:   "text",
			Message: fmt.Sprintf("text must be between 1 and %d characters", maxLen),
		}
	}
	return nil
}

func validateTTSRequest(req *ttsRequest) error {
	if err := validateTargetLanguage(req.TargetLanguageCode); err != nil {
		return err
	}

	model := BulbulV3
	if req.Model != nil {
		model = *req.Model
	}

	if err := validateText(req.Text, model); err != nil {
		return err
	}

	switch model {
	case BulbulV3:
		if req.Pitch != nil {
			return &sarvamaierrors.ValidationError{
				Field:   "pitch",
				Message: "pitch is not supported for bulbul:v3",
			}
		}
		if req.Loudness != nil {
			return &sarvamaierrors.ValidationError{
				Field:   "loudness",
				Message: "loudness is not supported for bulbul:v3",
			}
		}
		if req.EnablePreprocessing != nil {
			return &sarvamaierrors.ValidationError{
				Field:   "enable_preprocessing",
				Message: "enable_preprocessing is not supported for bulbul:v3",
			}
		}
		if req.Pace != nil {
			if *req.Pace < 0.5 || *req.Pace > 2.0 {
				return &sarvamaierrors.ValidationError{
					Field:   "pace",
					Message: "pace must be between 0.5 and 2.0 for bulbul:v3",
				}
			}
		}
		if req.Temperature != nil {
			if *req.Temperature < 0.01 || *req.Temperature > 2.0 {
				return &sarvamaierrors.ValidationError{
					Field:   "temperature",
					Message: "temperature must be between 0.01 and 2.0",
				}
			}
		}
		if req.SpeechSampleRate != nil {
			allowedRates := []SpeechSampleRate{SampleRate8000, SampleRate16000, SampleRate22050, SampleRate24000, SampleRate32000, SampleRate44100, SampleRate48000}
			if !slices.Contains(allowedRates, *req.SpeechSampleRate) {
				return &sarvamaierrors.ValidationError{
					Field:   "speech_sample_rate",
					Message: "sample rate must be 8000, 16000, 22050, 24000, 32000, 44100, or 48000 for bulbul:v3",
				}
			}
		}

	case BulbulV2:
		if req.Pitch != nil {
			if *req.Pitch < -0.75 || *req.Pitch > 0.75 {
				return &sarvamaierrors.ValidationError{
					Field:   "pitch",
					Message: "pitch must be between -0.75 and 0.75",
				}
			}
		}
		if req.Loudness != nil {
			if *req.Loudness < 0.3 || *req.Loudness > 3.0 {
				return &sarvamaierrors.ValidationError{
					Field:   "loudness",
					Message: "loudness must be between 0.3 and 3.0",
				}
			}
		}
		if req.Temperature != nil {
			return &sarvamaierrors.ValidationError{
				Field:   "temperature",
				Message: "temperature is not supported for bulbul:v2",
			}
		}
		if req.Pace != nil {
			if *req.Pace < 0.3 || *req.Pace > 3.0 {
				return &sarvamaierrors.ValidationError{
					Field:   "pace",
					Message: "pace must be between 0.3 and 3.0 for bulbul:v2",
				}
			}
		}
		if req.SpeechSampleRate != nil {
			allowedRates := []SpeechSampleRate{SampleRate8000, SampleRate16000, SampleRate22050, SampleRate24000}
			if !slices.Contains(allowedRates, *req.SpeechSampleRate) {
				return &sarvamaierrors.ValidationError{
					Field:   "speech_sample_rate",
					Message: "sample rate must be 8000, 16000, 22050, or 24000 for bulbul:v2",
				}
			}
		}
	}

	return nil
}

func validateTTSStreamRequest(cfg *ttsStreamRequest) error {
	if err := validateTargetLanguage(cfg.TargetLanguageCode); err != nil {
		return err
	}

	if cfg.Speaker == "" {
		return &sarvamaierrors.ValidationError{
			Field:   "speaker",
			Message: "speaker is required for streaming",
		}
	}

	model := BulbulV2
	if cfg.Model != nil {
		model = *cfg.Model
	}

	switch model {
	case BulbulV2:
		if cfg.Pitch != nil {
			if *cfg.Pitch < -0.75 || *cfg.Pitch > 0.75 {
				return &sarvamaierrors.ValidationError{
					Field:   "pitch",
					Message: "pitch must be between -0.75 and 0.75 for bulbul:v2",
				}
			}
		}
		if cfg.Loudness != nil {
			if *cfg.Loudness < 0.3 || *cfg.Loudness > 3.0 {
				return &sarvamaierrors.ValidationError{
					Field:   "loudness",
					Message: "loudness must be between 0.3 and 3.0 for bulbul:v2",
				}
			}
		}
		if cfg.Temperature != nil {
			return &sarvamaierrors.ValidationError{
				Field:   "temperature",
				Message: "temperature is not supported for bulbul:v2",
			}
		}
		if cfg.Pace != nil {
			if *cfg.Pace < 0.3 || *cfg.Pace > 3.0 {
				return &sarvamaierrors.ValidationError{
					Field:   "pace",
					Message: "pace must be between 0.3 and 3.0 for bulbul:v2",
				}
			}
		}
		if cfg.SpeechSampleRate != nil {
			allowedRates := []SpeechSampleRate{SampleRate8000, SampleRate16000, SampleRate22050, SampleRate24000}
			if !slices.Contains(allowedRates, *cfg.SpeechSampleRate) {
				return &sarvamaierrors.ValidationError{
					Field:   "speech_sample_rate",
					Message: "sample rate must be 8000, 16000, 22050, or 24000 for bulbul:v2",
				}
			}
		}

	case BulbulV3Beta:
		if cfg.Pitch != nil {
			return &sarvamaierrors.ValidationError{
				Field:   "pitch",
				Message: "pitch is not supported for bulbul:v3-beta",
			}
		}
		if cfg.Loudness != nil {
			return &sarvamaierrors.ValidationError{
				Field:   "loudness",
				Message: "loudness is not supported for bulbul:v3-beta",
			}
		}
		if cfg.Temperature != nil {
			if *cfg.Temperature < 0.01 || *cfg.Temperature > 1.0 {
				return &sarvamaierrors.ValidationError{
					Field:   "temperature",
					Message: "temperature must be between 0.01 and 1.0",
				}
			}
		}
		if cfg.Pace != nil {
			if *cfg.Pace < 0.5 || *cfg.Pace > 2.0 {
				return &sarvamaierrors.ValidationError{
					Field:   "pace",
					Message: "pace must be between 0.5 and 2.0 for bulbul:v3-beta",
				}
			}
		}
		if cfg.SpeechSampleRate != nil {
			allowedRates := []SpeechSampleRate{SampleRate8000, SampleRate16000, SampleRate22050, SampleRate24000}
			if !slices.Contains(allowedRates, *cfg.SpeechSampleRate) {
				return &sarvamaierrors.ValidationError{
					Field:   "speech_sample_rate",
					Message: "sample rate must be 8000, 16000, 22050, or 24000 for bulbul:v3-beta",
				}
			}
		}
	}

	return nil
}
