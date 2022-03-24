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

	ans := n
	a := [303][303]int{}
	pos := [303]int{}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Fscan(in, &a[i][j])
		}
		pos[i] = 1
	}

	cnt := [303]int{}
	vis := [303]bool{}
	mx := 0
	for t := 1; t < m; t++ {
		for i := 1; i <= m; i++ {
			cnt[i] = 0
		}
		for i := 1; i <= n; i++ {
			for vis[a[i][pos[i]]] {
				pos[i]++
			}
			if pos[i] <= m {
				cnt[a[i][pos[i]]]++
				if (cnt[a[i][pos[i]]]) > cnt[mx] {
					mx = a[i][pos[i]]
				}
			}
		}
		ans = min(ans, cnt[mx])
		vis[mx] = true
	}

	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
