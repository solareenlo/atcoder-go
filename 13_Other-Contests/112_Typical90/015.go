package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int
	fmt.Scan(&N)

	var p, m [100001]int
	p[0] = 1
	m[0] = 1
	for i := 1; i <= N; i++ {
		p[i] = p[i-1] * i % mod
		m[i] = divMod(m[i-1], i)
	}
	for i := 1; i <= N; i++ {
		var ans int
		rem := N - 1
		for j := 1; j <= N && rem >= 0; j++ {
			ans = (ans + (p[rem+j]*m[j]%mod)*m[rem]%mod) % mod
			rem -= i
		}
		fmt.Fprintln(out, ans)
	}
}

const mod = 1000000007

func divMod(a, b int) int {
	ret := a * modInv(b)
	ret %= mod
	return ret
}

func modInv(a int) int {
	b, u, v := mod, 1, 0
	for b != 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	u %= mod
	if u < 0 {
		u += mod
	}
	return u
}
