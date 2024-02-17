package main

import (
	"fmt"
	"regexp"
	"strings"
)

type PickupInfo struct {
	ExpressCompany string `json:"express_company"` // 快递公司
	PickupCode     string `json:"pickup_code"`     // 取件码
}

func GetContentMaps(sortedKeywords []string, text string) []PickupInfo {
	var result []PickupInfo
	pattern := `\w+-[\w-]+`
	re := regexp.MustCompile(pattern)

	for i := 0; i < len(sortedKeywords); i++ {
		keyword := sortedKeywords[i]
		startIndex := 0
		if i > 0 {
			//prevKeyword := sortedKeywords[i-1]
			prevIndex := re.FindStringIndex(text)
			if len(prevIndex) > 0 {
				startIndex = prevIndex[1]
			}
		}

		matches := re.FindAllStringSubmatch(text[startIndex:], -1)
		for _, match := range matches {
			if len(match) > 0 {
				content := PickupInfo{
					ExpressCompany: keyword,
					PickupCode:     match[0],
				}
				result = append(result, content)
			}
			//子匹配不匹配处理
		}
	}
	return result
}

func main() {
	sortedKeywords := []string{"圆通快递", "顺丰快递"}
	//text1 := ":是圆通快递取货码，取件码4aa，:是顺丰快递取货码，取件码1-b-c"
	text := ":是圆通快递取货码，取件码4-a-a，:是顺丰快递取货码，取件码1-b-c"
	result := GetContentMaps(sortedKeywords, text)
	//re := GetContentMapsInRange(sortedKeywords, text)
	for _, item := range result {
		fmt.Printf("关键字：%s，匹配内容：%s\n", item.ExpressCompany, item.PickupCode)
	}
	//for _, item := range re {
	//	fmt.Printf("关键字：%s，匹配内容：%s\n", item.ExpressCompany, item.PickupCode)
	//}
}
func GetContentMapsInRange(sortedKeywords []string, text string) []PickupInfo {
	var result []PickupInfo
	pattern := `\w+-[\w-]+`
	re := regexp.MustCompile(pattern)
	prevIndex := 0 // 记录上一个关键字的位置

	for i := 0; i < len(sortedKeywords); i++ {
		keyword := sortedKeywords[i]
		startIndex := prevIndex

		if i == 0 {
			// 如果是第一个关键字，子匹配内容为它前面所有内容
			content := PickupInfo{
				ExpressCompany: keyword,
				PickupCode:     text[:startIndex],
			}
			result = append(result, content)
			startIndex = 0 // 第一个关键字匹配前面所有内容
		} else {
			// 如果不是第一个关键字，子匹配内容为上一个关键字到当前关键字之间的内容
			lastKeyword := sortedKeywords[i-1]
			lastIndex := strings.Index(text, lastKeyword) + len(lastKeyword)
			content := PickupInfo{
				ExpressCompany: keyword,
				PickupCode:     text[lastIndex:startIndex],
			}
			result = append(result, content)
		}

		matches := re.FindAllStringSubmatch(text[startIndex:], -1)
		for _, match := range matches {
			if len(match) > 0 {
				content := PickupInfo{
					ExpressCompany: keyword,
					PickupCode:     match[0],
				}
				result = append(result, content)
			}
		}

		// 更新上一个关键字的位置
		if len(matches) > 0 {
			lastMatch := matches[len(matches)-1][0]
			prevIndex = strings.Index(text, lastMatch) + len(lastMatch)
		}
	}

	return result
}
