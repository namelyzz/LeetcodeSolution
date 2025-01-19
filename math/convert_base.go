package math

import "fmt"

func ConvertBaseDemo1(n, base int64) string {
	if base < 2 || base > 36 {
		return ""
	}
	if n == 0 {
		return "0"
	}

	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}

	digits := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var result []byte
	for n > 0 {
		rem := n % base
		result = append(result, digits[rem])
		n /= base
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return sign + string(result)
}

func ConvertBaseDemo2(numStr string, fromBase, toBase int) (string, error) {
	if fromBase < 2 || fromBase > 36 || toBase < 2 || toBase > 36 {
		return "", fmt.Errorf("进制范围需在 2 - 36 之间")
	}

	// 处理符号
	isNegative := false
	if numStr[0] == '-' {
		isNegative = true
		numStr = numStr[1:]
	}

	// 转化为十进制
	var decimalVal int64
	for _, r := range numStr {
		var val int
		if r >= '0' && r <= '9' {
			val = int(r - '0')
		} else if r >= 'A' && r <= 'Z' {
			val = 10 + int(r-'A')
		} else if r >= 'a' && r <= 'z' {
			val = 10 + int(r-'a')
		} else {
			return "", fmt.Errorf("非法字符：%c", r)
		}

		if val >= fromBase {
			return "", fmt.Errorf("字符 %c 超出进制 %d 的范围", r, fromBase)
		}

		decimalVal = decimalVal*int64(fromBase) + int64(val)
	}

	result := ConvertBaseDemo1(decimalVal, int64(toBase))
	if isNegative {
		result = "-" + result
	}
	return result, nil
}

func CountOne(n int) (count int) {
	for n > 0 {
		count++
		n &= n - 1 // 用于清除最右侧的 1
	}
	return count
}
