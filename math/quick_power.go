package math

func QuickPowerIterative(a, n int) int {
    if n < 0 {
        return 0 // 处理负指数情况
    }

    result := 1
    for n > 0 {
        if n&1 != 0 {
            result *= a
        }
        a *= a
        n >>= 1
    }
    return result
}

func QuickPowerRecursive(a, n int) int {
    if n < 0 {
        return 0
    }
    if n == 0 {
        return 1
    }
    if n == 1 {
        return a
    }

    half := QuickPowerRecursive(a, n>>1)
    if n&1 == 1 {
        return half * half * a
    }
    return half * half
}
