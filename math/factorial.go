package math

import (
	"math/big"
)

// Factorial 计算 n 的阶乘，返回 big.Int 类型避免溢出
func Factorial(n int) *big.Int {
	if n < 0 {
		return big.NewInt(0) // 处理负数输入
	}

	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}
