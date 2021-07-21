package main

import (
	"bufio"
	"fmt"
	"os"
)

const inf int = 1000000007

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Scan(&n, &m)

	f := [401][401]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			f[i][j] = inf
		}
		f[i][i] = 0
	}

	var a, b, c int
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &a, &b, &c)
		f[a-1][b-1] = c
	}

	res := 0
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				f[i][j] = min(f[i][j], f[i][k]+f[k][j])
				if f[i][j] < inf {
					res += f[i][j]
				}
			}
		}
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
