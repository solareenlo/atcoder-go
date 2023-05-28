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
	var s string
	fmt.Fscan(in, &n, &s)
	var p2 [400040]int
	p2[0] = 1
	for i := 0; i < n*2; i++ {
		p2[i+1] = p2[i] * 2
		if p2[i+1] >= mod {
			p2[i+1] -= mod
		}
	}
	ans := 0
	t := [2]int{1, 0}
	for i := 0; i < n; i++ {
		if s[i] == '0' {
			ans = (ans + (t[0]*3+t[1])*p2[(n-i-1)*2]) % mod
			t[0] = (t[0] + t[1]*2) % mod
		} else {
			t[1] = (t[0]*2 + t[1]) % mod
		}
	}
	ans = (ans + t[0] + t[1]) % mod
	fmt.Println(ans)
}
