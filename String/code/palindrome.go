package code

// longestPalindromeBruteForce 暴力枚举最长回文子串
func longestPalindromeBruteForce(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}

	maxLen, start := 1, 0
	runes := []rune(s)

	// 第一重循环：子串起始位置
	for i := 0; i < n; i++ {
		// 第二重循环：子串结束位置
		for j := i; j < n; j++ {
			// 第三重循环：检查回文
			if isPalindrome(runes, i, j) {
				curLen := j - i + 1
				if curLen > maxLen {
					maxLen = curLen
					start = i
				}
			}
		}
	}

	return string(runes[start : start+maxLen])
}

// isPalindrome 双指针回文检测
func isPalindrome(runes []rune, left, right int) bool {
	for left < right {
		if runes[left] != runes[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// longestPalindromeCenter 中心扩散法
func longestPalindromeCenter(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}

	maxLen, start := 1, 0
	runes := []rune(s)

	// 遍历所有可能的中心点（包括间隙）
	for i := 0; i < n; i++ {
		// 处理奇数长度回文（单中心）
		len1 := expandAroundCenter(runes, i, i)
		// 处理偶数长度回文（双中心）
		len2 := expandAroundCenter(runes, i, i+1)

		curLen := max(len1, len2)
		if curLen > maxLen {
			maxLen = curLen
			// 计算起始位置：i - (长度-1)/2
			start = i - (curLen-1)/2
		}
	}

	return string(runes[start : start+maxLen])
}

func expandAroundCenter(runes []rune, left, right int) int {
	for left >= 0 && right < len(runes) && runes[left] == runes[right] {
		left--
		right++
	}

	// 返回有效回文长度：right - left - 1
	return right - left - 1
}
