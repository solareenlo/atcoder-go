package main

import "fmt"

var (
	jc = [1000001]int{}
	ny = [1000001]int{}
)

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)
	z := n + m + k

	jc[0] = 1
	for i := 1; i <= z; i++ {
		jc[i] = i * jc[i-1] % mod
	}
	ny[z] = powMod(jc[z], mod-2)
	for i := z; i > 0; i-- {
		ny[i-1] = i * ny[i] % mod
	}

	ans := 0
	o := 1
	for i := n; i <= z; i++ {
		ans = (ans + C(n-1, i-1)*o%mod*powMod(3, z-i)) % mod
		o = (o + o) % mod
		if i >= n+m {
			o = (o - C(m, i-n) + mod) % mod
		}
		if i >= n+k {
			o = (o - C(i-n-k, i-n) + mod) % mod
		}
	}
	fmt.Println(ans)
}

func C(x, y int) int {
	return jc[y] * ny[x] % mod * ny[(y)-(x)] % mod
}

const mod = 1000000007

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
