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
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("[ERR] Error loading .env file")
	}
}

func detectLanguage(w io.Writer, text string, req *translatepb.DetectLanguageRequest) (string, error) {
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
		MimeType: "text/plain",
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
	for _, language := range resp.Languages {
		// The language detected.
		fmt.Fprintf(w, "Language code: %v\n", language.GetLanguageCode())
		// Confidence of detection result for this language.
		fmt.Fprintf(w, "Confidence: %v\n", language.GetConfidence())
	}

	return resp.Languages[0].GetLanguageCode(), nil
}

func TranslateText(text, sourceLang, targetLang string) (string, error) {
	ctx := context.Background()

	// Set GOOGLE_APPLICATION_CREDENTIALS
	err := os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "C:\\Users\\RR\\go\\src\\projects\\BIb_bot\\translation-api-412323-7e71139aef69.json")
	if err != nil {
		log.Fatal("[ERR] Unable to set GOOGLE_APPLICATION_CREDENTIALS:", err)
	}

	client, err := translate.NewTranslationClient(ctx)
	if err != nil {
		log.Fatal("[ERR] Unable to create translation client:", err)
	}
	defer client.Close()

	// Вызов Google Translate API
	translation, err := client.TranslateText(ctx, &translatepb.TranslateTextRequest{
		Contents:       []string{text},
		TargetLanguageCode: targetLang,
		SourceLanguageCode: sourceLang,
	})
	if err != nil {
		log.Fatal("[ERR] Unable to translate text:", err)
		return "", err
	}

	fmt.Printf("[DEBUG] Google Translate API Response: %+v\n", translation)

	return translation.Translations[0].GetTranslatedText(), nil
}
	