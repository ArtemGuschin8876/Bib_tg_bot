package translate

import (
	"context"
	"fmt"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

func TranslateTextWithModel(targetLanguage, text, model string) (string, error) {
	
	ctx := context.Background()

	lang, err := language.Parse(targetLanguage)
	if err != nil {
			return "", fmt.Errorf("language.Parse: %w", err)
	}

	client, err := translate.NewClient(ctx)
	if err != nil {
			return "", fmt.Errorf("translate.NewClient: %w", err)
	}
	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, lang, &translate.Options{
			Model: model, // Either "nmt" or "base".
	})
	if err != nil {
			return "", fmt.Errorf("Translate: %w", err)
	}
	if len(resp) == 0 {
			return "", nil
	}
	fmt.Println("11")
	return resp[0].Text, nil
}