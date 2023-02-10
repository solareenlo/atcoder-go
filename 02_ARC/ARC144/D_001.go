package main

import "fmt"

const mod = 998244353

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	ans := 0
	lst := (mod + 1) / 2
	for i := 0; i <= n; i++ {
		lst = 2 * lst % mod
		lst = lst * ((k + 1 - i) % mod) % mod
		lst = lst * powMod(i+1, mod-2) % mod
		ans = (ans + lst) % mod
		lst = lst * (n - i) % mod
		lst = lst * powMod(i+1, mod-2) % mod
	}
	fmt.Println(ans)
}

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
