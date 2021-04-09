package mytranslate

// func createContext() {
// 	// Create a google translate client.
// 	ctx := context.Background()
// 	tc, err := translate.NewTranslationClient(ctx, option.WithCredentialsFile("translation-api-project-307416-8cf7f95bb9a6.json"))
// 	if err != nil {
// 		fmt.Println("Client failed.")
// 		panic(err)
// 	}
// }

// func translateJP(ctx context.Context, text string, tc *translate.TranslationClient) (string, error) {
// 	req := &translatepb.TranslateTextRequest{
// 		Contents:           []string{text},
// 		SourceLanguageCode: "ja",
// 		TargetLanguageCode: "en",
// 		Parent:             "projects/translation-api-project-307416",
// 	}
// 	resp, err := tc.TranslateText(ctx, req)
// 	if err != nil {
// 		fmt.Println("Translate failed.")
// 		return "", err
// 	}
// 	return resp.Translations[0].TranslatedText, nil
// }
