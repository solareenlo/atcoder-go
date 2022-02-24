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

	const N = 3030
	a := make([]int, N)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	const mod = 998244353
	f := [N][N]int{}
	c := [N][N]int{}
	f[0][0] = 1
	c[0][0] = 1
	for i := 1; i <= n; i++ {
		c[i][0] = 1
		c[i][i] = 1
		for j := 1; j < i; j++ {
			c[i][j] = (c[i-1][j] + c[i-1][j-1]) % mod
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			f[i][j] = (2*f[i-1][j]*j + f[i-1][j-1]) % mod
		}
	}

	ans := 0
	for i := 0; i <= n; i++ {
		ans = (ans + 1*c[n][i]*f[m][n]) % mod
	}

	for l := 1; l <= n; l++ {
		for r := l; r <= n; r++ {
			x := l - 1
			y := n - r
			ans = (ans + 1*c[x+y][x]*f[m][x+y]) % mod
			if a[r+1] < a[r] {
				break
			}
		}
	}

	fmt.Println(ans)
}
