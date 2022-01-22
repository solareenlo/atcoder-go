package main

import (
	"fmt"
	"os"
	"strings"
)

const N = 405

var (
	n       int
	sx      int
	sy      int
	a       = [N][N]int{}
	b       = [N][N]int{}
	visited = [N][N]int{}
	chk     int
	tot     int
	cnt     int
	s       = make([]string, N)
)

func valid(A, B int) bool {
	return (1 <= A && A <= n && 1 <= B && B <= n && s[A][B] != '#')
}

func dfs(A, B, C int) {
	if !valid(A, B) || visited[A][B] == chk {
		return
	}
	visited[A][B] = chk
	cnt++
	if C != 0 {
		if b[A][B] == 0 {
			dfs(A+1, B, C^1)
		}
		if b[A][B] == 1 {
			dfs(A-1, B, C^1)
		}
	} else {
		if a[A][B] == 0 {
			dfs(A, B+1, C^1)
		}
		if a[A][B] == 1 {
			dfs(A, B-1, C^1)
		}
	}
}

func check() {
	for i := 1; i <= n; i++ {
		C := 0
		for j := 1; j <= n; j++ {
			if a[i][j] != 2 {
				a[i][j] = C
				C ^= 1
			}
		}
	}
	for i := 1; i <= n; i++ {
		C := 0
		for j := 1; j <= n; j++ {
			if b[j][i] != 2 {
				b[j][i] = C
				C ^= 1
			}
		}
	}
	chk++
	cnt = 0
	tmp := 0
	if a[sx][sy] == 2 {
		tmp++
	}
	dfs(sx, sy, tmp)
	if cnt == tot {
		fmt.Println("POSSIBLE")
		os.Exit(0)
	}
}

func solve() {
	for i := 1; i <= n; i++ {
		C := 0
		for j := 1; j <= n; j++ {
			if a[i][j] != 2 {
				C++
			}
		}
		if C%2 == 0 {
			continue
		}
		for j := 1; j <= n; j++ {
			if a[i][j] != 2 {
				a[i][j] = 2
				check()
				a[i][j] = 0
			}
		}
		return
	}
	for i := 1; i <= n; i++ {
		C := 0
		for j := 1; j <= n; j++ {
			if b[j][i] != 2 {
				C++
			}
		}
		if C%2 == 0 {
			continue
		}
		for j := 1; j <= n; j++ {
			if b[j][i] != 2 {
				b[j][i] = 2
				check()
				b[j][i] = 0
			}
		}
		return
	}
}

func main() {
	fmt.Scan(&n)
	s[0] = strings.Repeat("#", n+2)
	for i := 1; i <= n; i++ {
		fmt.Scan(&s[i])
		s[i] = "#" + s[i] + "#"
		for j := 1; j <= n; j++ {
			if s[i][j] == '#' {
				a[i][j] = 2
				b[i][j] = 2
			} else {
				tot++
			}
			if s[i][j] == 's' {
				sx = i
				sy = j
			}
		}
	}
	s[n+1] = strings.Repeat("#", n+2)
	a[sx][sy] = 2
	solve()
	a[sx][sy] = 0
	b[sx][sy] = 2
	solve()
	fmt.Println("IMPOSSIBLE")
}
