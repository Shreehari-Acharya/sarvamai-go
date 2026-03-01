package text

import (
	"strings"
	"testing"

	"github.com/Shreehari-Acharya/sarvamai-go/languages"
)

func TestValidateTranslateRequest(t *testing.T) {
	tests := []struct {
		name    string
		req     translateRequest
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid mayura:v1 request with all options",
			req: translateRequest{
				Input:              "Hello",
				SourceLanguageCode: "en-IN",
				TargetLanguageCode: "hi-IN",
				Model:              ptrTranslateModel(ModelMayura),
				Mode:               ptrTranslateMode(ModeFormal),
				OutputScript:       ptrOutputScript(OutputScriptRoman),
			},
			wantErr: false,
		},
		{
			name: "valid sarvam-translate:v1 request",
			req: translateRequest{
				Input:              "Hello",
				SourceLanguageCode: "en-IN",
				TargetLanguageCode: "ta-IN",
				Model:              ptrTranslateModel(ModelSarvamTranslate),
			},
			wantErr: false,
		},
		{
			name: "valid request with auto source language (mayura only)",
			req: translateRequest{
				Input:              "नमस्ते",
				SourceLanguageCode: "auto",
				TargetLanguageCode: "en-IN",
				Model:              ptrTranslateModel(ModelMayura),
			},
			wantErr: false,
		},
		{
			name: "empty input should error",
			req: translateRequest{
				Input:              "",
				SourceLanguageCode: "en-IN",
				TargetLanguageCode: "hi-IN",
			},
			wantErr: true,
			errMsg:  "input text cannot be empty",
		},
		{
			name: "input exceeds 1000 chars for mayura",
			req: translateRequest{
				Input:              string(make([]byte, 1001)),
				SourceLanguageCode: "en-IN",
				TargetLanguageCode: "hi-IN",
				Model:              ptrTranslateModel(ModelMayura),
			},
			wantErr: true,
			errMsg:  "input text cannot exceed 1000 characters for mayura model",
		},
		{
			name: "input exceeds 2000 chars for sarvam-translate",
			req: translateRequest{
				Input:              string(make([]byte, 2001)),
				SourceLanguageCode: "en-IN",
				TargetLanguageCode: "hi-IN",
				Model:              ptrTranslateModel(ModelSarvamTranslate),
			},
			wantErr: true,
			errMsg:  "input text cannot exceed 2000 characters for sarvam-translate:v1 model",
		},
		{
			name: "invalid source language for mayura",
			req: translateRequest{
				Input:              "Hello",
				SourceLanguageCode: "as-IN",
				TargetLanguageCode: "hi-IN",
				Model:              ptrTranslateModel(ModelMayura),
			},
			wantErr: true,
			errMsg:  "invalid source language code for mayura model",
		},
		{
			name: "invalid target language for mayura",
			req: translateRequest{
				Input:              "Hello",
				SourceLanguageCode: "en-IN",
				TargetLanguageCode: "as-IN",
				Model:              ptrTranslateModel(ModelMayura),
			},
			wantErr: true,
			errMsg:  "invalid target language code for mayura model",
		},
		{
			name: "invalid mode for mayura",
			req: translateRequest{
				Input:              "Hello",
				SourceLanguageCode: "en-IN",
				TargetLanguageCode: "hi-IN",
				Model:              ptrTranslateModel(ModelMayura),
				Mode:               ptrTranslateMode("invalid-mode"),
			},
			wantErr: true,
			errMsg:  "invalid mode for mayura model",
		},
		{
			name: "invalid output script for mayura",
			req: translateRequest{
				Input:              "Hello",
				SourceLanguageCode: "en-IN",
				TargetLanguageCode: "hi-IN",
				Model:              ptrTranslateModel(ModelMayura),
				OutputScript:       ptrOutputScript("invalid-script"),
			},
			wantErr: true,
			errMsg:  "invalid output script for mayura model",
		},
		{
			name: "non-formal mode with sarvam-translate should error",
			req: translateRequest{
				Input:              "Hello",
				SourceLanguageCode: "en-IN",
				TargetLanguageCode: "hi-IN",
				Model:              ptrTranslateModel(ModelSarvamTranslate),
				Mode:               ptrTranslateMode(ModeModernColloquial),
			},
			wantErr: true,
			errMsg:  "only 'formal' mode is supported for sarvam-translate model",
		},
		{
			name: "output_script with sarvam-translate should error",
			req: translateRequest{
				Input:              "Hello",
				SourceLanguageCode: "en-IN",
				TargetLanguageCode: "hi-IN",
				Model:              ptrTranslateModel(ModelSarvamTranslate),
				OutputScript:       ptrOutputScript(OutputScriptRoman),
			},
			wantErr: true,
			errMsg:  "transliteration is not supported for sarvam-translate model",
		},
		{
			name: "valid mayura with modern-colloquial mode",
			req: translateRequest{
				Input:              "Hello",
				SourceLanguageCode: "en-IN",
				TargetLanguageCode: "hi-IN",
				Model:              ptrTranslateModel(ModelMayura),
				Mode:               ptrTranslateMode(ModeModernColloquial),
			},
			wantErr: false,
		},
		{
			name: "valid mayura with classic-colloquial mode",
			req: translateRequest{
				Input:              "Hello",
				SourceLanguageCode: "en-IN",
				TargetLanguageCode: "hi-IN",
				Model:              ptrTranslateModel(ModelMayura),
				Mode:               ptrTranslateMode(ModeClassicColloquial),
			},
			wantErr: false,
		},
		{
			name: "valid mayura with code-mixed mode",
			req: translateRequest{
				Input:              "Hello",
				SourceLanguageCode: "en-IN",
				TargetLanguageCode: "hi-IN",
				Model:              ptrTranslateModel(ModelMayura),
				Mode:               ptrTranslateMode(ModeCodeMixed),
			},
			wantErr: false,
		},
		{
			name: "sarvam-translate with new language (assamese)",
			req: translateRequest{
				Input:              "Hello",
				SourceLanguageCode: "as-IN",
				TargetLanguageCode: "en-IN",
				Model:              ptrTranslateModel(ModelSarvamTranslate),
			},
			wantErr: false,
		},
		{
			name: "sarvam-translate with new language (urdu)",
			req: translateRequest{
				Input:              "Hello",
				SourceLanguageCode: "ur-IN",
				TargetLanguageCode: "en-IN",
				Model:              ptrTranslateModel(ModelSarvamTranslate),
			},
			wantErr: false,
		},
		{
			name: "default model is sarvam-translate",
			req: translateRequest{
				Input:              string(make([]byte, 1500)),
				SourceLanguageCode: "en-IN",
				TargetLanguageCode: "hi-IN",
			},
			wantErr: false,
		},
		{
			name: "default model with 2001 chars should error",
			req: translateRequest{
				Input:              string(make([]byte, 2001)),
				SourceLanguageCode: "en-IN",
				TargetLanguageCode: "hi-IN",
			},
			wantErr: true,
			errMsg:  "input text cannot exceed 2000 characters for sarvam-translate:v1 model",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateTranslateRequest(tt.req)
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got nil")
					return
				}
				if tt.errMsg != "" && !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("got %q, want to contain %q", err.Error(), tt.errMsg)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			}
		})
	}
}

func TestValidateTransliterateRequest(t *testing.T) {
	tests := []struct {
		name    string
		req     transliterateRequest
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid transliteration en-IN to hi-IN",
			req: transliterateRequest{
				Input:              "Hello",
				SourceLanguageCode: "en-IN",
				TargetLanguageCode: "hi-IN",
			},
			wantErr: false,
		},
		{
			name: "valid transliteration hi-IN to en-IN",
			req: transliterateRequest{
				Input:              "नमस्ते",
				SourceLanguageCode: "hi-IN",
				TargetLanguageCode: "en-IN",
			},
			wantErr: false,
		},
		{
			name: "valid transliteration with auto source",
			req: transliterateRequest{
				Input:              "नमस्ते",
				SourceLanguageCode: "auto",
				TargetLanguageCode: "en-IN",
			},
			wantErr: false,
		},
		{
			name: "empty input should error",
			req: transliterateRequest{
				Input:              "",
				SourceLanguageCode: "en-IN",
				TargetLanguageCode: "hi-IN",
			},
			wantErr: true,
			errMsg:  "input text cannot be empty",
		},
		{
			name: "input exceeds 1000 chars",
			req: transliterateRequest{
				Input:              string(make([]byte, 1001)),
				SourceLanguageCode: "en-IN",
				TargetLanguageCode: "hi-IN",
			},
			wantErr: true,
			errMsg:  "input text cannot exceed 1000 characters",
		},
		{
			name: "invalid source language",
			req: transliterateRequest{
				Input:              "Hello",
				SourceLanguageCode: "fr-FR",
				TargetLanguageCode: "hi-IN",
			},
			wantErr: true,
			errMsg:  "invalid source language code for transliterate",
		},
		{
			name: "invalid target language",
			req: transliterateRequest{
				Input:              "Hello",
				SourceLanguageCode: "en-IN",
				TargetLanguageCode: "fr-FR",
			},
			wantErr: true,
			errMsg:  "invalid target language code for transliterate",
		},
		{
			name: "valid with numerals format native",
			req: transliterateRequest{
				Input:              "Hello 123",
				SourceLanguageCode: "en-IN",
				TargetLanguageCode: "hi-IN",
				NumeralsFormat:     ptrNumeralsFormat(NumeralsNative),
			},
			wantErr: false,
		},
		{
			name: "valid with spoken form",
			req: transliterateRequest{
				Input:              "Hello",
				SourceLanguageCode: "en-IN",
				TargetLanguageCode: "hi-IN",
				SpokenForm:         ptrBool(true),
			},
			wantErr: false,
		},
		{
			name: "valid with all options",
			req: transliterateRequest{
				Input:                      "Hello 123",
				SourceLanguageCode:         "en-IN",
				TargetLanguageCode:         "hi-IN",
				NumeralsFormat:             ptrNumeralsFormat(NumeralsNative),
				SpokenForm:                 ptrBool(true),
				SpokenFormNumeralsLanguage: ptrSpokenFormNumeralsLanguage(SpokenFormNumeralsEnglish),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateTransliterateRequest(tt.req)
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got nil")
					return
				}
				if tt.errMsg != "" && !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("got %q, want to contain %q", err.Error(), tt.errMsg)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			}
		})
	}
}

func TestValidateDetectLanguageRequest(t *testing.T) {
	tests := []struct {
		name    string
		req     detectLanguageRequest
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid Hindi input",
			req: detectLanguageRequest{
				Input: "नमस्ते",
			},
			wantErr: false,
		},
		{
			name: "valid English input",
			req: detectLanguageRequest{
				Input: "Hello world",
			},
			wantErr: false,
		},
		{
			name: "empty input should error",
			req: detectLanguageRequest{
				Input: "",
			},
			wantErr: true,
			errMsg:  "input text cannot be empty",
		},
		{
			name: "input exceeds 1000 chars",
			req: detectLanguageRequest{
				Input: string(make([]byte, 1001)),
			},
			wantErr: true,
			errMsg:  "input text cannot exceed 1000 characters",
		},
		{
			name: "valid with 1000 chars exactly",
			req: detectLanguageRequest{
				Input: string(make([]byte, 1000)),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateDetectLanguageRequest(tt.req)
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got nil")
					return
				}
				if tt.errMsg != "" && !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("got %q, want to contain %q", err.Error(), tt.errMsg)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			}
		})
	}
}

func TestValidateInputTextForTranslation(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		model   TranslateModel
		wantErr bool
		errMsg  string
	}{
		{
			name:    "empty input for mayura",
			input:   "",
			model:   ModelMayura,
			wantErr: true,
			errMsg:  "input text cannot be empty",
		},
		{
			name:    "valid input for mayura (1000 chars)",
			input:   string(make([]byte, 1000)),
			model:   ModelMayura,
			wantErr: false,
		},
		{
			name:    "input exceeds 1000 for mayura",
			input:   string(make([]byte, 1001)),
			model:   ModelMayura,
			wantErr: true,
			errMsg:  "input text cannot exceed 1000 characters for mayura model",
		},
		{
			name:    "empty input for sarvam-translate",
			input:   "",
			model:   ModelSarvamTranslate,
			wantErr: true,
			errMsg:  "input text cannot be empty",
		},
		{
			name:    "valid input for sarvam-translate (2000 chars)",
			input:   string(make([]byte, 2000)),
			model:   ModelSarvamTranslate,
			wantErr: false,
		},
		{
			name:    "input exceeds 2000 for sarvam-translate",
			input:   string(make([]byte, 2001)),
			model:   ModelSarvamTranslate,
			wantErr: true,
			errMsg:  "input text cannot exceed 2000 characters for sarvam-translate:v1 model",
		},
		{
			name:    "invalid model",
			input:   "hello",
			model:   TranslateModel("invalid"),
			wantErr: true,
			errMsg:  "invalid model",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateInputTextForTranslation(tt.input, tt.model)
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got nil")
					return
				}
				if tt.errMsg != "" && !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("got %q, want to contain %q", err.Error(), tt.errMsg)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			}
		})
	}
}

func TestValidateInputTextForDetectionAndTransliteration(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "empty input",
			input:   "",
			wantErr: true,
			errMsg:  "input text cannot be empty",
		},
		{
			name:    "valid input (100 chars)",
			input:   "Hello world",
			wantErr: false,
		},
		{
			name:    "valid input (1000 chars)",
			input:   string(make([]byte, 1000)),
			wantErr: false,
		},
		{
			name:    "input exceeds 1000 chars",
			input:   string(make([]byte, 1001)),
			wantErr: true,
			errMsg:  "input text cannot exceed 1000 characters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateInputTextForDetectionAndTransliteration(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got nil")
					return
				}
				if tt.errMsg != "" && !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("got %q, want to contain %q", err.Error(), tt.errMsg)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			}
		})
	}
}

func ptrTranslateModel(m TranslateModel) *TranslateModel {
	return &m
}

func ptrTranslateMode(m TranslateMode) *TranslateMode {
	return &m
}

func ptrOutputScript(o OutputScript) *OutputScript {
	return &o
}

func ptrNumeralsFormat(n NumeralsFormat) *NumeralsFormat {
	return &n
}

func ptrSpokenFormNumeralsLanguage(s SpokenFormNumeralsLanguage) *SpokenFormNumeralsLanguage {
	return &s
}

func ptrBool(b bool) *bool {
	return &b
}

func ptrLanguageCode(c languages.Code) *languages.Code {
	return &c
}
