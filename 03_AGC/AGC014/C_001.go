package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)

	const N = 2222
	var sx, sy int
	mp := [N][N]bool{}
	for i := 1; i <= n; i++ {
		var s string
		fmt.Fscan(in, &s)
		for j := 1; j <= m; j++ {
			switch s[j-1] {
			case '#':
				mp[i][j] = true
			case '.':
				mp[i][j] = false
			case 'S':
				mp[i][j] = false
				sx = i
				sy = j
			}
		}
	}

	dx := []int{-1, 0, 0, 1}
	dy := []int{0, -1, 1, 0}
	type node struct{ x, y, d int }
	q := make([]node, 0)
	q = append(q, node{sx, sy, k})
	vis := [N][N]bool{}
	vis[sx][sy] = true
	ans := 1 << 60
	for len(q) > 0 {
		t := q[0]
		q = q[1:]
		ans = min(ans, min(t.x-1, min(t.y-1, min(n-t.x, m-t.y))))
		for i := 0; i < 4; i++ {
			tx := t.x + dx[i]
			ty := t.y + dy[i]
			if tx < 1 || ty < 1 || tx > n || ty > m || mp[tx][ty] || vis[tx][ty] || t.d == 0 {
				continue
			}
			q = append(q, node{tx, ty, t.d - 1})
			vis[tx][ty] = true
		}
	}
	fmt.Println((ans+k-1)/k + 1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
