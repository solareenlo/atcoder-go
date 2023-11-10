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

	var A [1 << 18]int
	var C [18][2][2]int

	var n, q int
	fmt.Fscan(in, &n, &q)
	for i := 0; i < 1<<n; i++ {
		fmt.Fscan(in, &A[i])
	}
	for i := 0; i < n; i++ {
		C[i][0][0] = 1
		C[i][1][1] = 1
	}
	for q > 0 {
		q--
		var x, y int
		fmt.Fscan(in, &x, &y)
		for k := 0; k < 2; k++ {
			C[x][y^1][k] = (C[x][y^1][k] + C[x][y][k]) % MOD
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < 1<<n; j++ {
			if ((j >> i) & 1) != 0 {
				p := A[j^(1<<i)]
				q := A[j]
				A[j^(1<<i)] = (C[i][0][0]*p%MOD + C[i][0][1]*q%MOD) % MOD
				A[j] = (C[i][1][0]*p%MOD + C[i][1][1]*q%MOD) % MOD
			}
		}
	}
	for i := 0; i < 1<<n; i++ {
		fmt.Fprintf(out, "%d ", A[i])
	}
}
