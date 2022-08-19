package practice

import (
	"strings"
	"unicode"
)

// ReplaceBlank 字符串中只包含大小写字母且长度小于等于1000，空格用%20替换
func ReplaceBlank(s string) (string, bool) {
	// 校验长度
	if len([]rune(s)) > 1000 {
		return s, false
	}
	for _, v := range s {
		// 校验是否不为大小写字母
		if string(v) != " " && unicode.IsLetter(v) == false {
			return s, false
		}
	}
	// -1 表示不限制替换数量
	return strings.Replace(s, " ", "%20", -1), true
}

/*
	unicode.IsLetter()
	strings.Replace()
*/
