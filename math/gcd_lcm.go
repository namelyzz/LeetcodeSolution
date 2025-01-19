package math

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func SteinGCD(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	k := 0
	for (a&1) == 0 && (b&1) == 0 {
		a >>= 1
		b >>= 1
		k++
	}

	for a != b {
		if a&1 == 0 {
			a >>= 1
		} else if b&1 == 0 {
			b >>= 1
		} else if a > b {
			a -= b
			a >>= 1
		} else {
			b -= a
			b >>= 1
		}
	}

	return a << k
}

func LCM(a, b int) int {
	return a * b / GCD(a, b)
}
