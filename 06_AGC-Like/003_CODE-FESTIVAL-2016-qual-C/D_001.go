package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	const N = 333
	a := [N]string{}
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	s := [N][N]int{}
	f := [N][N]int{}
	ans := 0
	for p := 1; p < m; p++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				s[i][j] = s[i-1][j-1]
				if a[i][p-1] == a[j][p] {
					s[i][j]++
				}
				f[i][j] = min(f[i-1][j], f[i][j-1]) + s[i][j]
			}
		}
		ans += f[n][n]
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
