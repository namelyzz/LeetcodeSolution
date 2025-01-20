package code

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func Test_longestPalindromeBruteForce(t *testing.T) {
	s := "aaabbbccccbbbaaaaaabbbaaaaaa"
	res := longestPalindromeBruteForce(s)
	assert.Equal(t, res, "aaabbbccccbbbaaa")
}

func Test_longestPalindromeCenter(t *testing.T) {
	s := "aaabbbccccbbbaaaaaabbbaaaaaa"
	res := longestPalindromeCenter(s)
	assert.Equal(t, res, "aaabbbccccbbbaaa")
}

func TestPalindromeExecutionTime(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	letters := []rune("abcdefghijklmnopqrstuvwxyz")
	s := make([]rune, 100000)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	str := string(s)

	// 测试中心扩散法
	start := time.Now()
	longestPalindromeCenter(str)
	t.Logf("中心扩散法耗时: %v\n", time.Since(start))

	// 测试暴力解法
	start = time.Now()
	longestPalindromeBruteForce(str)
	t.Logf("暴力解法耗时: %v\n", time.Since(start))
}
