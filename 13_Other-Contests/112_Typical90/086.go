package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 1000000007

	var n, Q int
	fmt.Fscan(in, &n, &Q)

	var d, s [50]int
	for i := 0; i < Q; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c, &d[i])
		a--
		b--
		c--
		s[i] = (1 << a) | (1 << b) | (1 << c)
	}

	ans := 1
	for i := 0; i < 60; i++ {
		tmp := 0
		for j := 0; j < 1<<n; j++ {
			ok := true
			for k := 0; k < Q; k++ {
				if (!((s[k] & j) > 0) && ((d[k]>>i)&1) != 0) || (((s[k] & j) > 0) && !(((d[k] >> i) & 1) != 0)) {
					ok = false
				}
			}
			if ok {
				tmp++
			}
		}
		ans *= tmp
		ans %= MOD
	}
	fmt.Println(ans)
}
