package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const B = 100

	initMod()

	var q int
	fmt.Fscan(in, &q)
	res := make([]int, q)
	a := make([]int, q)
	b := make([]int, q)
	c := make([]int, q)
	var buf [B][B][]int
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &a[i], &b[i], &c[i])
		if a[i] < b[i] {
			a[i], b[i] = b[i], a[i]
		}
		if a[i] < B {
			buf[a[i]][b[i]] = append(buf[a[i]][b[i]], i)
		} else {
			y := 0
			for y*a[i] <= c[i] {
				x := (c[i] - y*a[i]) / b[i]
				res[i] = (res[i] + nCrMod(x+y+1, x)) % mod
				y++
			}
		}
	}
	for s := 0; s < B; s++ {
		for t := 0; t < B; t++ {
			dp := make([]int, 10001)
			for k := 0; k <= 10000; k++ {
				dp[k] = 1
				if k >= s {
					dp[k] = (dp[k] + dp[k-s]) % mod
				}
				if k >= t {
					dp[k] = (dp[k] + dp[k-t]) % mod
				}
			}
			for _, i := range buf[s][t] {
				res[i] = dp[c[i]]
			}
		}
	}
	for _, x := range res {
		fmt.Println(x)
	}
}

const mod = 1000000007
const size = 100001

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

func nCrMod(n, r int) int {
	if n < r || n < 0 || r < 0 {
		return 0
	}
	return fact[n] * invf[r] % mod * invf[n-r] % mod
}
