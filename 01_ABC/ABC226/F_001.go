package main

import "fmt"

const mod = 998244353

func P(x, T int) int {
	if T != 0 {
		if T&1 != 0 {
			return x * P(x, T-1) % mod
		}
		return P(x*x%mod, T/2)
	}
	return 1
}

func main() {
	var n, p int
	fmt.Scan(&n, &p)

	D := make([]map[int]int, 53)
	for i := range D {
		D[i] = map[int]int{}
	}
	F := map[int]int{}
	I := map[int]int{}
	D[0][1] = 1
	F[0] = 1
	I[0] = 1

	for i := 1; i < 53; i++ {
		F[i] = F[i-1] * i % mod
		I[i] = P(F[i], mod-2)
	}

	for i := 1; i <= n; i++ {
		for j := i; j >= 0; j-- {
			for k, x := range D[j] {
				l := lcm(k, i-j)
				D[i][l] += x * F[n-j-1] % mod * I[n-i]
				D[i][l] %= mod
			}
		}
	}

	res := 0
	for k, x := range D[n] {
		res += x * P(k, p)
		res %= mod
	}
	fmt.Println(res)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
