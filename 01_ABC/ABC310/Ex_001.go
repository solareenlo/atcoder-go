package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MINF = -9187201950435737472
	const maxn = 606

	var f [maxn][maxn]int
	var g, c, d [maxn]int
	var h [maxn * maxn]int

	var n, H int
	fmt.Fscan(in, &n, &H)
	L := 1
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &c[i], &d[i])
		L = max(L, abs(c[i]))
	}

	for i := 0; i < maxn; i++ {
		for j := 0; j < maxn; j++ {
			f[i][j] = MINF
		}
	}
	f[0][0] = 0
	for t := 1; t <= L+L; t++ {
		for i := 1; i <= n; i++ {
			for j := max(c[i], 0); max(j, j-c[i]) < L+L; j++ {
				f[t][j-c[i]] = max(f[t][j-c[i]], f[t-1][j]+d[i])
			}
		}
	}
	for i := range h {
		h[i] = MINF
	}
	h[0] = 0
	mx := 0
	for i := 1; i <= L+L; i++ {
		for j := 0; j < L+L; j++ {
			g[i] = max(g[i], f[i][j])
		}
		if mx == 0 || g[i]*mx > g[mx]*i {
			mx = i
		}
		for j := i; j <= ((L * L) << 2); j++ {
			h[j] = max(h[j], h[j-i]+g[i])
		}
	}
	ans := int(1e18)
	for i := 0; i <= ((L * L) << 2); i++ {
		ans = min(ans, (max(H-h[i], 0)+g[mx]-1)/g[mx]*mx+i)
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
