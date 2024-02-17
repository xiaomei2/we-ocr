package logic

type Content struct {
	ExpressCompany string `json:"express_company"` //快递公司
	PickupCode     string `json:"pickup_code"`     //
}

// 关键字之间内容
func GetFirstContentMap(newText string, keywords []string) []Content {
	keywordsContentMap := GetBehindContentMap(keywords, newText)
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

// 关键字前面内容
func GetSecondContentMap(newText string, keywords []string) []Content {
	keywordsContentMap := getFrontContentMap(keywords, newText)
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
