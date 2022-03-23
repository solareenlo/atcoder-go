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

	a := make([]int, m+1)
	b := make([]int, m+1)
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &a[i], &b[i])
	}

	f := [404][404]int{}
	g := [404]int{}
	for i := 1; i <= n; i++ {
		f[i][i] = 1
		for j := m; j >= 1; j-- {
			if f[i][a[j]] != 0 && f[i][b[j]] != 0 {
				g[i] = 1
				break
			}
			if f[i][a[j]] != 0 {
				f[i][b[j]] = 1
			} else if f[i][b[j]] != 0 {
				f[i][a[j]] = 1
			}
		}
	}

	ans := 0
	for i := 1; i < n; i++ {
		if g[i] != 0 {
			continue
		}
		for j := i + 1; j <= n; j++ {
			if g[j] != 0 {
				continue
			}
			ans++
			for k := 1; k <= n; k++ {
				if f[i][k] != 0 && f[j][k] != 0 {
					ans--
					break
				}
			}
		}
	}
	fmt.Println(ans)
}
