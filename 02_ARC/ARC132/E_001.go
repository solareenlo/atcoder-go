package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	var s string
	fmt.Fscan(in, &n, &s)

	f := make([]int, 100005)
	f[0] = 1
	for i := 1; i < n+1; i++ {
		f[i] = divMod(f[i-1]*(2*i-1)%mod, (2 * i))
	}
	l := 0
	r := strings.Count(s, ".")
	j := 0
	ans := 0
	for i := 0; i < n+1; i++ {
		if i == n || s[i] == '.' {
			ans += (j + strings.Count(s[j:i], "<")) * f[l] % mod * f[r] % mod
			ans %= mod
			j = i + 1
			l++
			r--
		}
	}
	fmt.Println(ans)
}

const mod = 998244353

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
