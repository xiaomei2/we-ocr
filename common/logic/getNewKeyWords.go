package logic

// 赋予新关键词
func GetNewKeyWords(tokenToKeywordMap map[string]string) []string {
	var keywords []string
	for keyword := range tokenToKeywordMap {
		keywords = append(keywords, keyword)
	}
	return keywords
}
