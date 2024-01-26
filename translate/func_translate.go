package translate

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/translate"
	"github.com/joho/godotenv"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("[ERR] Error loading .env file")
	}
}

// TranslateText переводит текст из одного языка на другой
func TranslateText(text, sourceLang, targetLang string) (string, error) {
	translateAPIKey := os.Getenv("TOKEN_TRANSLATE")

	ctx := context.Background()

	client, err := translate.NewClient(ctx, option.WithAPIKey(translateAPIKey))
	if err != nil {
		log.Fatal("[ERR] Unable to create translation client:", err)
	}
	defer client.Close()

	// Вызов Google Translate API
	translation, err := client.Translate(ctx, []string{text}, language.Russian, nil)
	if err != nil {
		log.Fatal("[ERR] Unable to translate text:", err)
		return "", err
	}

	fmt.Printf("[DEBUG] Google Translate API Response: %+v\n", translation)

	return translation[0].Text, nil
}

