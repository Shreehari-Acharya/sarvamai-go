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
		"auto":   true,
		"bn-IN":  true,
		"en-IN":  true,
		"gu-IN":  true,
		"hi-IN":  true,
		"kn-IN":  true,
		"ml-IN":  true,
		"mr-IN":  true,
		"od-IN":  true,
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
)
