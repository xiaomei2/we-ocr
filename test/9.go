package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func checkAndReturnMatches(input string) (string, error) {
	// 编译正则表达式
	pattern := `\w+-[\w-]+`
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", fmt.Errorf("正则表达式编译错误: %s", err.Error())
	}

	// 查找匹配项
	match := re.FindString(input)
	if match == "" {
		return "没有找到匹配项", nil
	}

	return match, nil
}

func findKeywordsContent(text string, keywords []string) {
	// 记录上一次找到关键词的位置和名称
	lastIndex := -1
	lastKeyword := ""

	// 遍历文本的每个字符
	for i := 0; i <= len(text); i++ {
		for _, keyword := range keywords {
			// 检查当前位置开始的字符串是否以关键字开头
			if strings.HasPrefix(text[i:], keyword) {
				// 如果之前有关键字，则保存之前关键字和它的内容
				if lastKeyword != "" {
					fmt.Printf("字匹配：%s\n", text[lastIndex+len(lastKeyword):i])
				}
				// 更新最后找到的关键字和索引
				lastKeyword = keyword
				lastIndex = i
				i += len(keyword) - 1 // 跳过当前找到的关键词
				break
			}
		}
	}

	// 处理最后一个关键字后的内容
	if lastIndex != -1 && lastIndex+len(lastKeyword) <= len(text) {
		item, err := checkAndReturnMatches(text[lastIndex+len(lastKeyword):])
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("%s:%s\n", lastKeyword, item)
	}
}

func main() {
	text := "这是一个测试文本，关键词1好好1-S关键词1关键词1，最后是关键词3是a-d-1。"
	keywords := []string{"关键词1", "关键词2", "关键词3"}
	findKeywordsContent(text, keywords)
	//ExtractContentBeforeKeywords(text, keywords)
}

// KeywordPosition 用于存储关键词及其在文本中的位置
type KeywordPosition struct {
	Keyword  string
	Position int
}

// ExtractContentBeforeKeywords 提取并输出每个关键词到上一个关键词之前的内容
func ExtractContentBeforeKeywords(text string, keywords []string) {
	var positions []KeywordPosition

	// 查找所有关键词的位置，包括重复出现的情况
	for _, keyword := range keywords {
		start := 0
		for {
			index := strings.Index(text[start:], keyword)
			if index == -1 {
				break
			}
			absoluteIndex := start + index
			positions = append(positions, KeywordPosition{Keyword: keyword, Position: absoluteIndex})
			start = absoluteIndex + len(keyword) // 更新搜索起点，继续查找是否还有该关键词
		}
	}

	// 根据位置对关键词进行排序
	sort.Slice(positions, func(i, j int) bool {
		return positions[i].Position < positions[j].Position
	})

	// 提取并打印每个关键词到上一个关键词之前的内容，包括关键词本身
	lastPosition := 0
	for i, kp := range positions {
		content := ""
		if i == 0 {
			content = text[:kp.Position] // 获取第一个关键词之前的内容
		} else {
			content = text[lastPosition:kp.Position] // 获取上一个关键词到当前关键词之前的内容
		}
		fmt.Printf("%s ： %s\n", kp.Keyword, content) // 以指定格式打印
		lastPosition = kp.Position + len(kp.Keyword)
	}
	// 如果需要，打印最后一个关键词后的内容
	if lastPosition < len(text) {
		fmt.Printf("文本结束 ： %s\n", text[lastPosition:])
	}
}
