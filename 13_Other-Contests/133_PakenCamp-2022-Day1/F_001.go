package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, d int
	fmt.Fscan(in, &n, &m, &d)
	l := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &l[i])
	}
	ans := -1
	var a [20][20]int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &a[i][j])
		}
	}
	for i := 1; i < 1<<n; i++ {
		now := make([]int, m)
		sum := 0
		for j := 0; j < n; j++ {
			if ((i >> j) & 1) != 0 {
				sum += l[j]
				for k := 0; k < m; k++ {
					now[k] = max(now[k], a[j][k])
				}
			}
		}
		if sum > d {
			continue
		}
		s := 0
		for j := 0; j < m; j++ {
			s += now[j]
		}
		ans = max(ans, s)
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
