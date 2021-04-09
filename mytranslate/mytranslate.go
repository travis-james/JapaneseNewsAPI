package mytranslate

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
