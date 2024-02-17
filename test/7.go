package main

import (
	"cor/common/logic"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func main() {
	firstKeywords := []string{"取件码", "取货码"}
	secondKeywords := []string{"顺丰快递", "EMS快递", "申通快递", "圆通快递", "韵达快递", "极兔速递", "中通快递", "尼乔卫生站", "邮政快递", "韵达快递", "圆通速递"}
	// 获取图片文字内容
	//text := getImagesText()
	text := ":是圆通快递取货码，取件码4-a-a，"
	result := logic.IsKeywordBefore(text, firstKeywords, secondKeywords)
	newText, newKeywordMap := ReplaceKeywordsAndCreateMap(text, secondKeywords)
	keywords := GetNewKeyWords(newKeywordMap)
	fmt.Printf("result:%v\n", result)
	fmt.Printf("newText:%v\n", newText)
	fmt.Printf("Keywords:%v\n", keywords)
	keywordsContentMap := GetContentMap(keywords, newText)
	fmt.Printf("keywordsContentMap:%v\n", keywordsContentMap)
}

type Content struct {
	ExpressCompany string `json:"express_company"` //快递公司
	PickupCode     string `json:"pickup_code"`     //
}

func GetContentMap(sortedKeywords []string, text string) []Content {
	pattern := `\w+-[\w-]+`
	re := regexp.MustCompile(pattern)
	for _, keyword := range sortedKeywords {
	}
}
func ReplaceKeywordsAndCreateMap(text string, keywords []string) (string, map[string]string) {
	keywordPositions := make(map[int]string)
	tokenCounter := 1 // 用于为每个关键字实例生成递增的标记
	tokenToKeywordMap := make(map[string]string)

	// 查找关键字在文本中的所有位置
	for _, keyword := range keywords {
		startPos := 0
		for {
			index := strings.Index(text[startPos:], keyword)
			if index == -1 {
				break
			}
			// 调整index以反映全局位置
			index += startPos
			keywordPositions[index] = keyword
			startPos = index + len(keyword)
		}
	}

	// 按位置排序关键字
	var positions []int
	for pos := range keywordPositions {
		positions = append(positions, pos)
	}
	sort.Ints(positions)

	// 根据排序后的位置替换关键字并更新映射
	var sb strings.Builder
	lastPos := 0
	for _, pos := range positions {
		keyword := keywordPositions[pos]
		token := fmt.Sprintf("<*%d*>", tokenCounter)
		sb.WriteString(text[lastPos:pos]) // 添加关键字之前的文本
		sb.WriteString(token)             // 插入特定标记
		lastPos = pos + len(keyword)

		// 更新标记到关键字的映射
		tokenToKeywordMap[token] = keyword

		// 更新计数器，为下一个关键字实例生成新的标记
		tokenCounter++
	}
	sb.WriteString(text[lastPos:]) // 添加最后一部分文本

	return sb.String(), tokenToKeywordMap
}

// 赋予新关键词
func GetNewKeyWords(tokenToKeywordMap map[string]string) []string {
	var keywords []string
	for keyword := range tokenToKeywordMap {
		keywords = append(keywords, keyword)
	}
	return keywords
}
