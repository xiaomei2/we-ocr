package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "取件码931790在17:00前领取，申通快递凭石材市场44453C在19:00前领取，EMS快递提货码2-B在18:00前领取"
	firstKeywords := []string{"取件码", "取货码"}
	secondKeywords := []string{"顺丰快递", "EMS快递", "申通快递", "圆通快递", "韵达快递", "极兔速递", "中通快递", "尼乔卫生站", "邮政快递", "韵达快递", "圆通速递"}

	result := isKeywordBefore(text, firstKeywords, secondKeywords)
	fmt.Println(result) // 输出：true
}

// 判断文本中的 firstKeywords 的任一关键字是否在 secondKeywords 的任一关键字之前
func isKeywordBefore(text string, firstKeywords []string, secondKeywords []string) int {
	for _, first := range firstKeywords {
		for _, second := range secondKeywords {
			if strings.Contains(text, first) && strings.Contains(text, second) {
				if strings.Index(text, first) < strings.Index(text, second) {
					return 1
				}
			}
		}
	}
	return 0
}
