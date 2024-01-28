package translate

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	translate "cloud.google.com/go/translate/apiv3"
	"cloud.google.com/go/translate/apiv3/translatepb"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("[ERR] Error loading .env file")
	}
}


func detectLanguage(w io.Writer,  text string, req *translatepb.DetectLanguageRequest) (string,error) {
	 projectID := os.Getenv("ID_PROJECT")
	 text = "Hello, world!"

	 ctx := context.Background()
	 client, err := translate.NewTranslationClient(ctx)
	 if err != nil {
		 return "", fmt.Errorf("NewTranslationClient: %w", err)
	 }

	defer client.Close()

	req = &translatepb.DetectLanguageRequest{
			Parent:   fmt.Sprintf("projects/%s/locations/global", projectID),
			MimeType: "text/plain", // Mime types: "text/plain", "text/html"
			Source: &translatepb.DetectLanguageRequest_Content{
					Content: text,
			},
	}

	
	  resp, err := client.DetectLanguage(ctx, req)
    if err != nil {
        return "", fmt.Errorf("DetectLanguage: %w", err)
    }

	// Display list of detected languages sorted by detection confidence.
	// The most probable language is first.
	for _, language := range resp.GetLanguages() {
			// The language detected.
			fmt.Fprintf(w, "Language code: %v\n", language.GetLanguageCode())
			// Confidence of detection result for this language.
			fmt.Fprintf(w, "Confidence: %v\n", language.GetConfidence())
	}

	return "", fmt.Errorf("No languages detected")

}




func TranslateText(text, sourceLang, targetLang string) (string, error) {
	ctx := context.Background()

	sourceLang, err := detectLanguage(nil,text, &translatepb.DetectLanguageRequest{})
	if err != nil {
		return "", fmt.Errorf("Error detecting language: %w", err)
	}

	translateAPIKey := os.Getenv("TOKEN_TRANSLATE")
	


	client, err := translate.NewTranslationClient(ctx, option.WithAPIKey(translateAPIKey))
	if err != nil {
		log.Fatal("[ERR] Unable to create translation client:", err)
	}
	defer client.Close()

	// Вызов Google Translate API
	translation, err := client.TranslateText(ctx, &translatepb.TranslateTextRequest{})
	if err != nil {
		log.Fatal("[ERR] Unable to translate text:", err)
		return "", err
	}

	fmt.Printf("[DEBUG] Google Translate API Response: %+v\n", translation)

	return translation.Translations[0].String(), nil
}

