package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	d := make([]int, 99999)
	e := make([]int, 99999)
	d[0] = 1
	e[0] = 1

	const mod = 998244353
	s := 0
	tot := 1
	for i := 1; i <= n; i++ {
		var x int
		fmt.Fscan(in, &x)
		s += x
		tot = tot * 3 % mod
		for j := i * 301; j >= 0; j-- {
			d[j] = (d[j]*2 + d[max(0, j-x)]) % mod
			e[j] = (e[j] + e[max(0, j-x)]) % mod
		}
	}
	fmt.Println(((tot+3*(e[(s+1)/2]-e[s/2+1]-d[(s+1)/2]))%mod + mod) % mod)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
