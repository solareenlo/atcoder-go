package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const MOD = 998244353

	var q, k int
	fmt.Fscan(in, &q, &k)

	var F [5001]int
	F[0] = 1
	for q > 0 {
		q--
		var c string
		var x int
		fmt.Fscan(in, &c, &x)
		if c == "-" {
			for i := 0; i+x <= k; i++ {
				F[i+x] = (F[i+x] - F[i] + MOD) % MOD
			}
		} else {
			for i := k; i-x >= 0; i-- {
				F[i] = (F[i] + F[i-x]) % MOD
			}
		}
		fmt.Fprintln(out, (F[k]+MOD)%MOD)
	}
}
