package main

import (
	"fmt"
	"sort"
	"strings"
)

func getSortedKeywordsPositions(keywords []string, text string) map[string][]int {
	keywordPositions := make(map[string][]int)

	for _, keyword := range keywords {
		positions := []int{}
		index := 0
		for {
			position := strings.Index(text[index:], keyword)
			if position == -1 {
				break
			}
			absPosition := index + position
			positions = append(positions, absPosition)
			index = absPosition + len(keyword)
		}
		if len(positions) > 0 {
			keywordPositions[keyword] = positions
		}
	}

	sortedKeywords := make([]string, 0, len(keywordPositions))
	for k := range keywordPositions {
		sortedKeywords = append(sortedKeywords, k)
	}
	sort.Slice(sortedKeywords, func(i, j int) bool {
		return keywordPositions[sortedKeywords[i]][0] < keywordPositions[sortedKeywords[j]][0]
	})

	return keywordPositions
}

func printKeywordContext(keyword string, context string) {
	fmt.Printf("%s: %s\n", keyword, context)
}

func sortedKeys(m map[string][]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func main() {
	keywords := []string{"A", "B", "C"}
	text := "lAllAllSA你好B我好C好的"
	printKeywordsWithInterContext(keywords, text)
}
func printKeywordsWithPreContext(keywords []string, text string) {
	positions := getSortedKeywordsPositions(keywords, text)
	fmt.Printf("Keywords found in the text:\n")
	for _, keyword := range sortedKeys(positions) {
		for _, pos := range positions[keyword] {
			startIndex := pos - 10
			if startIndex < 0 {
				startIndex = 0
			}
			context := text[startIndex:pos]
			printKeywordContext(keyword, context)
		}
	}
}
func printKeywordsWithInterContext(keywords []string, text string) {
	positions := getSortedKeywordsPositions(keywords, text)
	fmt.Printf("Keywords found in the text:\n")

	sortedKeywords := sortedKeys(positions)
	for i, keyword := range sortedKeywords {
		if i < len(sortedKeywords)-1 {
			nextKeyword := sortedKeywords[i+1]
			for _, pos := range positions[keyword] {
				var context string
				nextPos := findNextPosition(positions[nextKeyword], pos+len(keyword))
				if nextPos != -1 {
					context = text[pos+len(keyword) : nextPos]
				} else {
					context = ""
				}
				printKeywordContext(keyword, context)
			}
		} else {
			// For the last keyword, print the context until the end of the text
			for _, pos := range positions[keyword] {
				context := text[pos+len(keyword):]
				printKeywordContext(keyword, context)
			}
		}
	}
}

func findNextPosition(positions []int, startPos int) int {
	for _, pos := range positions {
		if pos >= startPos {
			return pos
		}
	}
	return -1
}
