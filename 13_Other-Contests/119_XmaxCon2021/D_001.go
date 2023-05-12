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

	const Maxk = 8
	const Maxn = Maxk*2 - 1
	const Mod = 998244353

	var n, k int
	fmt.Fscan(in, &n, &k)

	if n >= 2*k {
		for i := 0; i < k; i++ {
			for j := 0; j < k; j++ {
				fmt.Fprint(out, "0 ")
			}
			fmt.Fprintln(out)
		}
		return
	}

	var a [Maxn][Maxk][Maxk]int
	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			for t := 0; t < k; t++ {
				fmt.Fscan(in, &a[i][j][t])
			}
		}
	}
	var f [1 << Maxn][Maxk][Maxk]int
	for i := 0; i < k; i++ {
		f[0][i][i] = 1
	}
	for mask := 0; mask < (1 << n); mask++ {
		for j := 0; j < k; j++ {
			for t := 0; t < k; t++ {
				if f[mask][j][t] == 0 {
					continue
				}
				for i := 0; i < n; i++ {
					if ((mask >> i) & 1) != 0 {
						continue
					}
					if popcount(uint32(mask>>i))&1 != 0 {
						for l := 0; l < k; l++ {
							f[mask|(1<<i)][j][l] = (f[mask|(1<<i)][j][l] + f[mask][j][t]*(Mod-a[i][t][l])) % Mod
						}
					} else {
						for l := 0; l < k; l++ {
							f[mask|(1<<i)][j][l] = (f[mask|(1<<i)][j][l] + f[mask][j][t]*a[i][t][l]) % Mod
						}
					}
				}
			}
		}
	}
	for i := 0; i < k; i++ {
		for j := 0; j < k; j++ {
			fmt.Fprintf(out, "%d ", f[(1<<n)-1][i][j])
		}
		fmt.Fprintln(out)
	}
}

func popcount(n uint32) int {
	return bits.OnesCount32(n)
}
