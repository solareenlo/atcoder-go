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

	var p [501]int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &p[i])
	}

	const mod = 998244353
	var f [501][501]int
	for i := n; i >= 1; i-- {
		f[i][i] = 1
		for j := i + 1; j <= n; j++ {
			f[i][j] = f[i+1][j]
			for k := i + 1; k < j; k++ {
				if p[i] < p[k] {
					f[i][j] = (f[i][j] + f[i+1][k]*f[k][j]) % mod
				}
			}
		}
	}
	fmt.Println(f[1][n])
}
