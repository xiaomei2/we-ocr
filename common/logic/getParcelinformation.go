package logic

type Content struct {
	ExpressCompany string `json:"express_company"` //快递公司
	PickupCode     string `json:"pickup_code"`     //
}

func GetFirstContentMap(text string, keywords []string) []Content {
	sortedKeywords := getSortedKeywordsPositions(keywords, text)
	keywordsContentMap := getBehindContentMap(sortedKeywords, text)
	var contents []Content
	for k, v := range keywordsContentMap {
		content := Content{
			ExpressCompany: k,
			PickupCode:     v,
		}
		contents = append(contents, content)
	}
	return contents
}

func GetSecondContentMap(text string, keywords []string) []Content {
	sortedKeywords := getSortedKeywordsPositions(keywords, text)
	keywordsContentMap := getFrontContentMap(sortedKeywords, text)
	var contents []Content
	for k, v := range keywordsContentMap {
		content := Content{
			ExpressCompany: k,
			PickupCode:     v,
		}
		contents = append(contents, content)
	}
	return contents
}
