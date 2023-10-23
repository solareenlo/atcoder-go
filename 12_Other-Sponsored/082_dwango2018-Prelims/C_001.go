package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 1000000007

	var n, m int
	fmt.Fscan(in, &n, &m)
	var a, b [101]int
	a[0] = -1
	b[0] = -1
	var cta, ctb [101]int
	sa, sb := 0, 0
	var ct int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		sa += a[i]
		if a[i] != a[i-1] {
			ct = 1
		}
		cta[i] = ct
		ct++
	}
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &b[i])
		sb += b[i]
		if b[i] != b[i-1] {
			ct = 1
		}
		ctb[i] = ct
		ct++
	}
	var dpa, dpb [1001]int
	dpa[0] = 1
	dpb[0] = 1
	for i := 1; i <= n; i++ {
		for j := cta[i]; j <= sb; j++ {
			dpa[j] += dpa[j-cta[i]]
			dpa[j] %= MOD
		}
	}
	for i := 1; i <= m; i++ {
		for j := ctb[i]; j <= sa; j++ {
			dpb[j] += dpb[j-ctb[i]]
			dpb[j] %= MOD
		}
	}
	fmt.Println(dpa[sb] * dpb[sa] % MOD)
}
