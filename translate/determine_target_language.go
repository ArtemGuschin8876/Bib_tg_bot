package translate


func DetermineTargetLanguage(sourceLanguage string) string {
	if sourceLanguage == "en" {
		return "ru"
	} else if sourceLanguage == "ru" {
		return "en"
	}
	return "" 
}