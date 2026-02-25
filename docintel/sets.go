package docintel

var allowedOutputFormats = map[OutputFormat]bool{
	OutputFormatHTML: true,
	OutputFormatMD:   true,
	OutputFormatJSON: true,
}

var allowedFileExtensions = map[string]bool{
	".pdf": true,
	".zip": true,
}
