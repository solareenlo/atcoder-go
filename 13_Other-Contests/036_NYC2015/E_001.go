package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const mod = 1000000007
	const MAX = 100010

	var n int
	fmt.Fscan(in, &n)

	d := make([]int, n)
	s := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &d[i])
		s += d[i]
	}

	var fact, inv, invfact [MAX]int
	fact[0] = 1
	inv[0] = 1
	inv[1] = 1
	invfact[0] = 1
	for i := 1; i < MAX; i++ {
		fact[i] = i * fact[i-1] % mod
		if i >= 2 {
			inv[i] = mod - inv[mod%i]*(mod/i)%mod
		}
		invfact[i] = invfact[i-1] * inv[i] % mod
	}

	if s != 2*n-2 {
		fmt.Println(0)
		return
	}

	ans := fact[n-2]
	for i := 0; i < n; i++ {
		ans = ans * invfact[d[i]-1] % mod
	}
	fmt.Println(ans)
}
