package math

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return false
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func SieveOfEratosthenes(n int) []int {
	if n < 2 {
		return []int{}
	}
	isPrime := make([]bool, n+1)
	for i := range isPrime {
		isPrime[i] = true
	}
	isPrime[0], isPrime[1] = false, false

	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}

	primes := make([]int, 0)
	for i, v := range isPrime {
		if v {
			primes = append(primes, i)
		}
	}
	return primes
}

func EulerSieve(n int) []int {
	if n < 2 {
		return []int{}
	}

	isPrime := make([]bool, n+1)
	primes := make([]int, 0)
	for i := range isPrime {
		isPrime[i] = true
	}
	isPrime[0], isPrime[1] = false, false

	for i := 2; i <= n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
		for j := 0; j < len(primes) && i*primes[j] <= n; j++ {
			isPrime[i*primes[j]] = false
			if i%primes[j] == 0 {
				break
			}
		}
	}
	return primes
}

func PrimeFactors(n int) (factors []int) {
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}

	if n > 1 {
		factors = append(factors, n)
	}
	return factors
}
