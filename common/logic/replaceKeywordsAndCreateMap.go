package logic

import (
	"fmt"
	"sort"
	"strings"
)

// 替换文本中的关键字，并为每个实例创建特定标记，然后生成标记到关键字的映射
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
