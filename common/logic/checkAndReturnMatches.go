package logic

import (
	"errors"
	"fmt"
	"regexp"
)

// CheckAndReturnMatches 检查字符串中是否存在匹配 \w+-[\w-]+ 的项，并返回匹配结果或提示信息
func CheckAndReturnMatches(input string) (string, error) {
	// 编译正则表达式
	pattern := `\w+-[\w-]+`
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", fmt.Errorf("正则表达式编译错误: %s", err.Error())
	}

	// 查找匹配项
	match := re.FindString(input)
	if match == "" {
		return "", errors.New("没有找到匹配项")
	}

	return match, nil
}
