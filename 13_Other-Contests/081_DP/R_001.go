package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const mod = 1000000007

	var n, k int
	fmt.Fscan(in, &n, &k)
	var to [64][51][51]int
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Fscan(in, &to[0][i][j])
		}
	}
	for i := 1; i <= 63; i++ {
		for j := 1; j <= n; j++ {
			for l := 1; l <= n; l++ {
				for k := 1; k <= n; k++ {
					to[i][j][l] = (to[i][j][l] + to[i-1][j][k]*to[i-1][k][l]) % mod
				}
			}
		}
	}
	var a [51]int
	for i := 1; i <= n; i++ {
		a[i] = 1
	}
	var b [51]int
	for i := 63; i >= 0; i-- {
		if (k & (1 << i)) != 0 {
			for j := 1; j <= n; j++ {
				b[j] = a[j]
				a[j] = 0
			}
			for j := 1; j <= n; j++ {
				for k := 1; k <= n; k++ {
					a[j] = (a[j] + to[i][j][k]*b[k]) % mod
				}
			}
		}
	}
	ans := 0
	for i := 1; i <= n; i++ {
		ans = (ans + a[i]) % mod
	}
	fmt.Println(ans)
}
