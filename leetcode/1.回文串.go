package leetcode

import "strings"

func IsPalindrome(s string) bool {
	if s == "" {
		return false
	}
	// 1. 将字符串全部转换为小写
	s = strings.ToLower(s)

	// 2. 双指针初始化
	left, right := 0, len(s)-1
	// 3. 双指针遍历
	// 3.1 确定双指针遍历条件
	for left < right {
		// 3.2 左指针向右遍历时，非数字、字母字符则跳过
		if !(s[left] >= 'a' && s[left] <= 'z' || s[left] >= '0' && s[left] <= '9') {
			left++
			continue
		}
		// 3.3 右指针向左遍历时，非数字、字母字符则跳过
		if !(s[right] >= 'a' && s[right] <= 'z' || s[right] >= '0' && s[right] <= '9') {
			right--
			continue
		}
		// 3.4 左指针指向值和右指针指向值不同时，则非回文串
		if s[left] != s[right] {
			return false
		}
		// 3.5 左指针向右
		left++
		// 3.6 右指针向左
		right--
	}
	// 4. 遍历完成，符合回文串特征
	return true
}
