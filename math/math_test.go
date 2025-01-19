package math

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestGcd(t *testing.T) {
	assert.Equal(t, 7, GCD(21, 7))
	assert.Equal(t, 12, GCD(36, 24))
	assert.Equal(t, 1, GCD(39, 29))
}

func TestSteinGCD(t *testing.T) {
	assert.Equal(t, 7, SteinGCD(21, 7))
	assert.Equal(t, 12, SteinGCD(36, 24))
	assert.Equal(t, 1, SteinGCD(39, 29))
	assert.Equal(t, 9, SteinGCD(18, 81))
}

func TestLCM(t *testing.T) {
	assert.Equal(t, 21, LCM(21, 7))
	assert.Equal(t, 72, LCM(36, 24))
	assert.Equal(t, 39*29, LCM(39, 29))
}

func TestPrime(t *testing.T) {
	assert.False(t, IsPrime(1))
	assert.True(t, IsPrime(3))
	assert.True(t, IsPrime(17))
	assert.False(t, IsPrime(88))

	t.Log(SieveOfEratosthenes(39))
	t.Log(SieveOfEratosthenes(17))
	t.Log(SieveOfEratosthenes(56))

	t.Log(EulerSieve(39))
	t.Log(EulerSieve(17))
	t.Log(EulerSieve(56))

	t.Log(PrimeFactors(81))
	t.Log(PrimeFactors(29))
	t.Log(PrimeFactors(99))
}

func TestQuickPower(t *testing.T) {
	assert.Equal(t, 8, QuickPowerIterative(2, 3))
	assert.Equal(t, 8, QuickPowerRecursive(2, 3))
	assert.Equal(t, 387420489, QuickPowerIterative(9, 9))
	assert.Equal(t, 387420489, QuickPowerRecursive(9, 9))
	assert.Equal(t, 0, QuickPowerIterative(15, -3))
	assert.Equal(t, 0, QuickPowerRecursive(15, -3))
}

func TestFibonacci(t *testing.T) {
	assert.Equal(t, 21, FibonacciIterative(8))
	assert.Equal(t, 21, FibonacciRecursive(8))
	assert.Equal(t, 514229, FibonacciRecursive(29))
	assert.Equal(t, 514229, FibonacciRecursive(29))
}

func TestConvertBase(t *testing.T) {
	assert.Equal(t, "11101", ConvertBaseDemo1(29, 2))
	res1, err1 := ConvertBaseDemo2("29", 10, 2)
	assert.Equal(t, "11101", res1)
	assert.Nil(t, err1)

	assert.Equal(t, "", ConvertBaseDemo1(29, 37))
	_, err2 := ConvertBaseDemo2("11101", 2, 37)
	assert.Errorf(t, err2, "非法")

	assert.Equal(t, "35", ConvertBaseDemo1(29, 8))
	res3, err3 := ConvertBaseDemo2("29", 16, 8)
	assert.Equal(t, "51", res3)
	assert.Nil(t, err3)

	assert.Equal(t, "-35", ConvertBaseDemo1(-29, 8))
	res4, err4 := ConvertBaseDemo2("-29", 16, 8)
	assert.Equal(t, "-51", res4)
	assert.Nil(t, err4)

	assert.Equal(t, 2, CountOne(3))
	assert.Equal(t, 4, CountOne(29))
	assert.Equal(t, 7, CountOne(7777))
	t.Logf("7777 的二进制：%s", ConvertBaseDemo1(7777, 2))
}

func TestSqrt(t *testing.T) {
	assert.InDelta(t, float64(43.8064), Sqrt(1919), 1e-4, "超出误差范围")
	assert.InDelta(t, float64(54.1202), Sqrt(2929), 1e-4, "超出误差范围")
	assert.InDelta(t, float64(60), Sqrt(3600), 1e-4, "超出误差范围")
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string // 使用字符串表示预期值
	}{
		{"n=0", 0, "1"},
		{"n=5", 5, "120"},
		{"n=10", 10, "3628800"},
		{"n=20", 20, "2432902008176640000"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Factorial(tt.input)
			expected, ok := new(big.Int).SetString(tt.expected, 10)
			if !ok {
				t.Fatalf("无法解析预期值 %q 为 big.Int", tt.expected)
			}
			assert.Equal(t, expected.String(), actual.String(), "阶乘结果不匹配")
		})
	}
}

func TestCombination(t *testing.T) {
	tests := []struct {
		name     string
		n, k     int
		expected int
	}{
		// 边界条件
		{"k > n", 5, 6, 0},
		{"k == 0", 10, 0, 1},
		{"k == n", 10, 10, 1},
		{"n == 0", 0, 0, 1},
		{"n == 0, k > 0", 0, 1, 0},

		// 典型用例
		{"C(5,2)", 5, 2, 10},
		{"C(6,3)", 6, 3, 20},
		{"C(10,4)", 10, 4, 210},
		{"C(20,10)", 20, 10, 184756},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Combination(tt.n, tt.k)
			assert.Equal(t, tt.expected, result,
				"Combination(%d, %d) 应该返回 %d，但得到 %d",
				tt.n, tt.k, tt.expected, result)
		})
	}
}
