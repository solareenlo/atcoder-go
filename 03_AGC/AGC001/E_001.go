package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int32
	fmt.Fscan(in, &n)

	const M = 2010
	const MOD = 1_000_000_007
	a := make([]int32, n+2)
	b := make([]int32, n+2)
	f := [2 * M][2 * M]int32{}
	for i := int32(1); i <= n; i++ {
		fmt.Fscan(in, &a[i], &b[i])
		f[M-a[i]][M-b[i]]++
	}

	g := [2 * M][2 * M]int32{}
	g[1][1] = 1
	for i := 1; i <= 4015; i++ {
		for j := 1; j <= 4015; j++ {
			f[i][j] = int32((int(f[i][j]) + int(f[i-1][j])%MOD + int(f[i][j-1])) % MOD)
			g[i][j] = int32((int(g[i][j]) + int(g[i-1][j])%MOD + int(g[i][j-1])) % MOD)
		}
	}

	ans := 0
	for i := int32(1); i <= n; i++ {
		ans = (ans + int(f[a[i]+M][b[i]+M])) % MOD
		ans = (ans + MOD - int(g[a[i]*2+1][b[i]*2+1])) % MOD
	}
	fmt.Println(500000004 * ans % MOD)
}
