package languages

// Language sets for different API services

// Mayura languages - 12 languages + auto (for translation)
var MayuraLanguages = map[Code]bool{
	"auto":  true,
	"bn-IN": true,
	"en-IN": true,
	"gu-IN": true,
	"hi-IN": true,
	"kn-IN": true,
	"ml-IN": true,
	"mr-IN": true,
	"or-IN": true,
	"pa-IN": true,
	"ta-IN": true,
	"te-IN": true,
}

// SarvamTranslate languages - 22 languages + auto (for translation)
var SarvamTranslateLanguages = map[Code]bool{
	"auto":   true,
	"bn-IN":  true,
	"en-IN":  true,
	"gu-IN":  true,
	"hi-IN":  true,
	"kn-IN":  true,
	"ml-IN":  true,
	"mr-IN":  true,
	"or-IN":  true,
	"pa-IN":  true,
	"ta-IN":  true,
	"te-IN":  true,
	"as-IN":  true,
	"brx-IN": true,
	"doi-IN": true,
	"kok-IN": true,
	"ks-IN":  true,
	"mai-IN": true,
	"mni-IN": true,
	"ne-IN":  true,
	"sa-IN":  true,
	"sat-IN": true,
	"sd-IN":  true,
	"ur-IN":  true,
}

// Transliterate languages - 12 languages + auto (same as Mayura)
var TransliterateLanguages = MayuraLanguages

// Saarika languages - 12 languages + unknown (for Speech services)
var SaarikaLanguages = map[Code]bool{
	"unknown": true, // auto-detect
	"hi-IN":   true, // Hindi
	"bn-IN":   true, // Bengali
	"kn-IN":   true, // Kannada
	"ml-IN":   true, // Malayalam
	"mr-IN":   true, // Marathi
	"or-IN":   true, // Odia
	"pa-IN":   true, // Punjabi
	"ta-IN":   true, // Tamil
	"te-IN":   true, // Telugu
	"en-IN":   true, // English
	"gu-IN":   true, // Gujarati
}

// Saaras languages - 22 languages - Saarika languages + 11 other languages (for Speech services)
var SaarasLanguages = map[Code]bool{
	"unknown": true, // auto-detect
	"hi-IN":   true, // Hindi
	"bn-IN":   true, // Bengali
	"kn-IN":   true, // Kannada
	"ml-IN":   true, // Malayalam
	"mr-IN":   true, // Marathi
	"or-IN":   true, // Odia
	"pa-IN":   true, // Punjabi
	"ta-IN":   true, // Tamil
	"te-IN":   true, // Telugu
	"en-IN":   true, // English
	"gu-IN":   true, // Gujarati
	"as-IN":   true, // Assamese
	"ur-IN":   true, // Urdu
	"ne-IN":   true, // Nepali
	"kok-IN":  true, // Konkani
	"ks-IN":   true, // Kashmiri
	"sd-IN":   true, // Sindhi
	"sa-IN":   true, // Sanskrit
	"sat-IN":  true, // Santali
	"mni-IN":  true, // Manipuri
	"brx-IN":  true, // Bodo
	"mai-IN":  true, // Maithili
	"doi-IN":  true, // Dogri
}

// TargetLanguages for TTS - 11 languages (same as Mayura without auto)
var TargetLanguages = map[Code]bool{
	"bn-IN": true, // Bengali
	"en-IN": true, // English
	"gu-IN": true, // Gujarati
	"hi-IN": true, // Hindi
	"kn-IN": true, // Kannada
	"ml-IN": true, // Malayalam
	"mr-IN": true, // Marathi
	"or-IN": true, // Odia
	"pa-IN": true, // Punjabi
	"ta-IN": true, // Tamil
	"te-IN": true, // Telugu
}

// allowed Language codes for Document Intelligence - docintel

var AllowedDocIntelLanguages = map[Code]bool{
	"hi-IN":   true, // Hindi
	"en-IN":   true, // English
	"bn-IN":   true, // Bengali
	"gu-IN":   true, // Gujarati
	"kn-IN":   true, // Kannada
	"ml-IN":   true, // Malayalam
	"mr-IN":   true, // Marathi
	"or-IN":   true, // Odia
	"pa-IN":   true, // Punjabi
	"ta-IN":   true, // Tamil
	"te-IN":   true, // Telugu
	"ur-IN":   true, // Urdu
	"as-IN":   true, // Assamese
	"bodo-IN": true, // Bodo
	"doi-IN":  true, // Dogri
	"ks-IN":   true, // Kashmiri
	"kok-IN":  true, // Konkani
	"mai-IN":  true, // Maithili
	"mni-IN":  true, // Manipuri
	"sa-IN":   true, // Sanskrit
	"sat-IN":  true, // Santali
	"sd-IN":   true, // Sindhi
	"ne-IN":   true, // Nepali
}
