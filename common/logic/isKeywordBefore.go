package logic

import "strings"

// 判断文本中的 firstKeywords 的任一关键字是否在 secondKeywords 的任一关键字之前
func IsKeywordBefore(text string, firstKeywords []string, secondKeywords []string) bool {
	for _, first := range firstKeywords {
		for _, second := range secondKeywords {
			if strings.Contains(text, first) && strings.Contains(text, second) {
				if strings.Index(text, first) < strings.Index(text, second) {
					return true
				}
			}
		}
	}
	return false
}
