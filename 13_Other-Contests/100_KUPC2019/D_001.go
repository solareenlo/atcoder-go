package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const mod = 998244353

	var n int
	var t string
	fmt.Fscan(in, &n, &t)
	s := " " + t + " "

	var inv, h [1000007]int
	inv[0] = 1
	inv[1] = 1
	for i := 2; i <= n+1; i++ {
		inv[i] = (mod - mod/i) * inv[mod%i] % mod
	}
	h[0] = 1
	h[1] = 1
	for i := 2; i <= n; i++ {
		h[i] = h[i-1] * (4*i - 2) % mod * inv[i+1] % mod
	}

	cnt := 0
	ans := 1
	for i := 1; i <= n+1; i++ {
		if s[i] == s[i-1] {
			cnt++
		} else {
			ans = ans * h[cnt] % mod
			cnt = 1
		}
	}
	fmt.Println(ans)
}
