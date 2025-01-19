package math

func Combination(n, k int) int {
	if k > n {
		return 0
	}
	if k == 0 || k == n {
		return 1
	}

	return Combination(n-1, k-1) + Combination(n-1, k)
}
