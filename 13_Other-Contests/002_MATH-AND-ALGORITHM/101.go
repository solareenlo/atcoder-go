package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	initMod()

	ans := make([]int, N)
	for i := 1; i <= N; i++ {
		for j := 1; j <= (N+i-1)/i; j++ {
			ans[i-1] += fact[N-(i-1)*(j-1)] * invf[j] % mod * invf[N-(i-1)*(j-1)-j]
			ans[i-1] %= mod
		}
	}
	for i := 0; i < N; i++ {
		fmt.Println(ans[i])
	}
}

const mod = 1000000007
const size = 100010

var fact, invf [size]int

func initMod() {
	fact[0] = 1
	invf[0] = 1
	for i := int(1); i < size; i++ {
		fact[i] = (fact[i-1] * i) % mod
		invf[i] = invMod(fact[i])
	}
}

func powMod(a, n int) int {
	res := int(1)
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n /= 2
	}
	return res
}

func invMod(a int) int {
	return powMod(a, mod-2)
}
