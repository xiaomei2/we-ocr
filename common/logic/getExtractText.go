package logic

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// 获取并排序文本中出现的关键字及其位置
func getSortedKeywordsPositions(keywords []string, text string) []string {
	keywordPositions := make(map[string]int)

	for _, keyword := range keywords {
		position := strings.Index(text, keyword)
		if position != -1 {
			keywordPositions[keyword] = position
		}
	}

	sortedKeywords := make([]string, 0, len(keywordPositions))
	for k := range keywordPositions {
		sortedKeywords = append(sortedKeywords, k)
	}
	sort.Slice(sortedKeywords, func(i, j int) bool {
		return keywordPositions[sortedKeywords[i]] < keywordPositions[sortedKeywords[j]]
	})

	return sortedKeywords
}

// 根据排序后的关键字获取它们之间的文本内容
// 关键字后面
func getBehindContentMap(sortedKeywords []string, text string) map[string]string {
	contentMap := make(map[string]string)
	textLength := len(text)
	pattern := `\w+-[\w-]+`
	re := regexp.MustCompile(pattern)

	for i, keyword := range sortedKeywords {
		start := strings.Index(text, keyword) + len(keyword)
		var end int
		if i+1 < len(sortedKeywords) {
			nextKeyword := sortedKeywords[i+1]
			end = strings.Index(text, nextKeyword)
		} else {
			end = textLength
		}
		content := text[start:end]
		var contentTest string
		// 应用正则表达式提纯
		matches := re.FindAllString(content, -1)
		contentTest = strings.Join(matches, ", ")
		// 判断子提取内容是否符合要求
		if len(matches) > 0 && len(contentTest) < 15 {
			//关键字字内容匹配就直接输出
			contentMap[keyword] = strings.Join(matches, ", ")
		} else {
			//关键字字内容匹配不满足要求就换FindAndPrintKeywords()
			printKeytext := FindAndPrintKeywords(content)
			if printKeytext != "2000" {
				contentMap[keyword] = printKeytext
			} else {
				//关键字字内容第二种方法不匹配就找关键字前面的内容
				precedingContent := GetPrecedingContent(keyword, text)
				precedingtext := FindAndPrintKeywords(precedingContent)
				if precedingtext != "2000" {
					contentMap[keyword] = precedingtext
				}
			}
		}
	}
	return contentMap
}

// 关键字前面
func getFrontContentMap(sortedKeywords []string, text string) map[string]string {
	contentMap := make(map[string]string)
	textLength := len(text)
	pattern := `\w+-[\w-]+`
	re := regexp.MustCompile(pattern)

	for i, keyword := range sortedKeywords {
		var start int
		if i == 0 {
			start = 0
		} else {
			prevKeyword := sortedKeywords[i-1]
			start = strings.Index(text, prevKeyword) + len(prevKeyword)
		}
		end := strings.Index(text, keyword)
		if end == -1 {
			end = textLength
		}
		content := text[start:end]
		var contentTest string
		matches := re.FindAllString(content, -1)
		fmt.Println(keyword, content)
		contentTest = strings.Join(matches, ", ")
		if len(matches) > 0 && len(contentTest) < 15 {
			contentMap[keyword] = strings.Join(matches, ", ")
		} else {
			//关键字字内容匹配不满足要求就换FindAndPrintKeywords()
			printKeytext := FindAndPrintKeywords(content)
			if printKeytext != "2000" {
				contentMap[keyword] = printKeytext
			}
		}
	}

	return contentMap
}

// 局部内容提取
func FindAndPrintKeywords(text string) string {
	rePickupCode := regexp.MustCompile(`取件码([a-zA-Z0-9]+)`)
	reAickupCode := regexp.MustCompile(`取件码:([a-zA-Z0-9]+)`)
	reDeliveryCode := regexp.MustCompile(`凭([a-zA-Z0-9]+)`)

	var result string

	AMatches := rePickupCode.FindAllStringSubmatch(text, -1)
	if len(AMatches) > 0 {
		for i, match := range AMatches {
			if i == 0 {
				result += match[1]
			} else {
				result += "\n" + match[1]
			}
		}
		return result
	}
	BMatches := reAickupCode.FindAllStringSubmatch(text, -1)
	if len(BMatches) > 0 {
		for i, match := range BMatches {
			if i == 0 {
				result += match[1]
			} else {
				result += "\n" + match[1]
			}
		}
		return result
	}
	CMatches := reDeliveryCode.FindAllStringSubmatch(text, -1)
	if len(CMatches) > 0 {
		for i, match := range CMatches {
			if i == 0 {
				result += match[1]
			} else {
				result += "\n" + match[1]
			}
		}
		return result
	}

	return "2000"
}

// 全文提取
func findFirstKeyword(text string, keywords []string) string {
	var foundKeyword string

	for _, keyword := range keywords {
		if strings.Contains(text, keyword) {
			fmt.Printf("找到关键字：%s", keyword)
			foundKeyword = keyword
			break
		}
	}

	return foundKeyword
}
func GetPrecedingContent(keyword string, text string) string {
	index := strings.Index(text, keyword)
	if index == -1 {
		return ""
	}
	return text[:index]
}
