package languages

type Code string

func (c Code) String() string {
	return string(c)
}

func (c Code) IsValid() bool {
	_, ok := Languages[c]
	return ok
}

var (
	Languages = map[Code]bool{
		"unknown": true, // unknown
		"auto":    true, // auto
		"bn-IN":   true, // Bengali
		"en-IN":   true, // English
		"gu-IN":   true, // Gujarati
		"hi-IN":   true, // Hindi
		"kn-IN":   true, // Kannada
		"ml-IN":   true, // Malayalam
		"mr-IN":   true, // Marathi
		"od-IN":   true, // Odia
		"pa-IN":   true, // Punjabi
		"ta-IN":   true, // Tamil
		"te-IN":   true, // Telugu
		"as-IN":   true, // Assamese
		"brx-IN":  true, // Bodo
		"doi-IN":  true, // Dogri
		"kok-IN":  true, // Konkani
		"ks-IN":   true, // Kashmiri
		"mai-IN":  true, // Maithili
		"mni-IN":  true, // Manipuri
		"ne-IN":   true, // Nepali
		"sa-IN":   true, // Sanskrit
		"sat-IN":  true, // Santali
		"sd-IN":   true, // Sindhi
		"ur-IN":   true, // Urdu
	}
)
