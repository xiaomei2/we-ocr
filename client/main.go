package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	text := "顺丰快递提货码1-A-1在17:00前领取"

	keywords := []string{"顺丰快递", "EMS快递", "申通快递"}
	count := 0
	for _, keyword := range keywords {
		count += strings.Count(text, keyword)
	}

	if count >= 2 {
		// 将关键字转换为正则表达式模式
		pattern := ""
		for i, keyword := range keywords {
			if i < len(keywords)-1 {
				pattern += fmt.Sprintf(`(%s).*?\b(\w+-[\w-]+)|`, keyword)
			} else {
				pattern += fmt.Sprintf(`(%s).*?\b(\w+-[\w-]+)`, keyword)
			}
		}

		re := regexp.MustCompile(pattern)
		matches := re.FindAllStringSubmatch(text, -1)

		for _, match := range matches {
			// 第一个匹配项是整个模式的匹配结果，从第二个开始才是提取的子匹配项

			for i := 1; i < len(match); i++ {
				if i%2 == 0 && match[i] != "" {
					company := match[i-1]
					code := match[i]
					fmt.Printf("公司：%s，提货码：%s\n", company, code)
				}
			}

		}
	} else {
		// 将关键字转换为正则表达式模式
		pattern := ""
		for i, keyword := range keywords {
			if i < len(keywords)-1 {
				pattern += fmt.Sprintf(`(%s).*?\b(\w+-[\w-]+)|`, keyword)
			} else {
				pattern += fmt.Sprintf(`(%s).*?\b(\w+-[\w-]+)`, keyword)
			}
		}

		re := regexp.MustCompile(pattern)
		matches := re.FindAllStringSubmatch(text, -1)

		for _, match := range matches {
			// 第一个匹配项是整个模式的匹配结果，从第二个开始才是提取的子匹配项

			for i := 1; i < len(match); i++ {
				if i%2 == 0 && match[i] != "" {
					company := match[i-1]
					code := match[i]
					fmt.Printf("公司：%s，提货码：%s\n", company, code)
				}
			}

		}
	}

}
