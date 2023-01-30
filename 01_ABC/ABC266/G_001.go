package main

import "fmt"

const N = 2000002

var pre [N]int

func main() {
	var r, g, b, k int
	fmt.Scan(&r, &g, &b, &k)

	pre[0] = 1
	for i := 1; i <= max(g+b, r+b); i++ {
		pre[i] = (pre[i-1] * i) % mod
	}

	ans := calc(g+b, g) * calc(g, k) % mod * calc(r+b, r-k) % mod
	fmt.Println(ans)

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func calc(x, y int) int {
	return pre[x] * powMod(pre[y], mod-2) % mod * powMod(pre[x-y], mod-2) % mod
}

const mod = 998244353

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n /= 2
	}
	return res
}
