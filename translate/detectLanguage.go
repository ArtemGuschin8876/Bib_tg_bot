package translate



import (
	"context"
	"fmt"

	"cloud.google.com/go/translate"
)


func DetectLanguage(text string) (*translate.Detection, error) {
	// text := "こんにちは世界"
	ctx := context.Background()
	client, err := translate.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("translate.NewClient: %w", err)
	}
	defer client.Close()
	lang, err := client.DetectLanguage(ctx, []string{text})
	if err != nil {
		return nil, fmt.Errorf("DetectLanguage: %w", err)
	}
	if len(lang) == 0 || len(lang[0]) == 0 {
		return nil, fmt.Errorf("DetectLanguage return value empty")
	}
	return &lang[0][0], nil

}

// [END translate_detect_language]