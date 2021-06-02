package mytranslate

import (
	"context"
	"fmt"

	translate "cloud.google.com/go/translate/apiv3"
	"google.golang.org/api/option"
	translatepb "google.golang.org/genproto/googleapis/cloud/translate/v3"
)

func TranslateJP(text string) (string, error) {
	// Create a google translate client.
	ctx := context.Background()
	tc, err := translate.NewTranslationClient(ctx, option.WithCredentialsFile("mytranslate/gtranslatekeys.json"))
	if err != nil {
		fmt.Println("Client failed.")
		return "", err
	}
	req := &translatepb.TranslateTextRequest{
		Contents:           []string{text},
		SourceLanguageCode: "ja",
		TargetLanguageCode: "en",
		Parent:             "projects/translation-api-project-307416",
		MimeType:           "text/plain",
	}
	resp, err := tc.TranslateText(ctx, req)
	if err != nil {
		fmt.Println("Translate failed.")
		return "", err
	}
	return resp.Translations[0].TranslatedText, nil
}
