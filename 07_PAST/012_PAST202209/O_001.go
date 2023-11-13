package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const MOD = 998244353

	var F, G [1 << 17][18]int
	var H [1 << 17][35]int

	var n, q int
	fmt.Fscan(in, &n, &q)
	for i := 0; i < 1<<n; i++ {
		fmt.Fscan(in, &F[i][popcount(uint32(i))])
	}
	for i := 0; i < 1<<n; i++ {
		fmt.Fscan(in, &G[i][popcount(uint32(i))])
	}
	for i := 0; i < n; i++ {
		for j := 0; j < 1<<n; j++ {
			if ((j >> i) & 1) != 0 {
				for p := 0; p <= n; p++ {
					F[j][p] = (F[j][p] + F[j^(1<<i)][p]) % MOD
					G[j][p] = (G[j][p] + G[j^(1<<i)][p]) % MOD
				}
			}
		}
	}
	for i := 0; i < 1<<n; i++ {
		for j := 0; j < n+1; j++ {
			for k := 0; k < n+1; k++ {
				H[i][j+k] = (H[i][j+k] + F[i][j]*G[i][k]%MOD) % MOD
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < 1<<n; j++ {
			if ((j >> i) & 1) != 0 {
				for p := 0; p <= 2*n; p++ {
					H[j][p] = (H[j][p] - H[j^(1<<i)][p] + MOD) % MOD
				}
			}
		}
	}
	for q > 0 {
		q--
		var x, y int
		fmt.Fscan(in, &x, &y)
		fmt.Fprintln(out, H[x][y])
	}
}

func popcount(n uint32) int {
	return bits.OnesCount32(n)
}
