package practice

import "strings"

// SameRegroupString 字符串长度都小于等于5000
func SameRegroupString(s1, s2 string) bool {
	length1, length2 := len([]rune(s1)), len([]rune(s2))
	if length1 > 5000 || length2 > 5000 || length1 != length2 {
		return false
	}
	for _, v := range s1 {
		if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) {
			return false
		}
	}
	return true
}
