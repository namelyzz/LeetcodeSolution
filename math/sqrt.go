package math

func Sqrt(n float64) float64 {
	if n < 0 {
		return -1
	}

	x := n
	for i := 0; i < 10; i++ { // 迭代 10 次，确保足够精准
		x = (x + n/x) / 2
	}
	return x
}
