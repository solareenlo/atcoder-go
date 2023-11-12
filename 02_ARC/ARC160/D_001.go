package main

import "fmt"

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)
	if (m % k) != 0 {
		fmt.Println(0)
		return
	}
	m /= k
	var C [2004]int
	C[0] = 1
	for i := 1; i <= n-k+1; i++ {
		for j := i; j > 0; j-- {
			C[j] = (C[j] + C[j-1]) % MOD
		}
	}
	iv := 1
	for i := 1; i <= n*2-k; i++ {
		iv = iv * i % MOD
	}
	iv = invMod(iv)
	ans := 0
	for i := 0; i <= n-k+1; i++ {
		if i*k <= m {
			mul := 1
			for j := 1; j <= n*2-k; j++ {
				mul = mul * (((m - i*k%MOD + MOD) + j) % MOD) % MOD
			}
			if (i & 1) != 0 {
				mul = (MOD - mul) % MOD
			}
			ans = (ans + ((mul * iv % MOD) * C[i] % MOD)) % MOD
		}
	}
	fmt.Println(ans)
}

const MOD = 998244353

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		n /= 2
	}
	return res
}

func invMod(a int) int {
	return powMod(a, MOD-2)
}
