package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	const INF = 1 << 60
	u, U, l, L := INF, INF, INF, INF
	d, D, r, R := 0, 0, 0, 0
	cnt0, cnt1 := 0, 0

	s := make([]string, n+2)
	s[0] = strings.Repeat(" ", n+2)
	s[n+1] = strings.Repeat(" ", n+2)
	for i := 1; i < n+1; i++ {
		fmt.Scan(&s[i])
		s[i] = " " + s[i]
		for j := 1; j < n+1; j++ {
			if s[i][j] == '#' {
				u = min(u, i)
				d = max(d, i)
				l = min(l, j)
				r = max(r, j)
				cnt0++
			}
		}
		s[i] += " "
	}

	t := make([]string, n+2)
	t[0] = strings.Repeat(" ", n+2)
	t[n+1] = strings.Repeat(" ", n+2)
	for i := 1; i < n+1; i++ {
		fmt.Scan(&t[i])
		t[i] = " " + t[i]
		for j := 1; j < n+1; j++ {
			if t[i][j] == '#' {
				U = min(U, i)
				D = max(D, i)
				L = min(L, j)
				R = max(R, j)
				cnt1++
			}
		}
		t[i] += " "
	}

	if cnt0 != cnt1 {
		fmt.Println("No")
		return
	}

	if d-u == R-L && r-l != D-U && r-l == R-L && d-u != D-U {
		fmt.Println("No")
		return
	}

	ok0, ok1, ok2, ok3 := true, true, true, true
	for i := u; i < d+1; i++ {
		for j := l; j < r+1; j++ {
			if D-j+l < 0 || L+i-u >= n+2 || s[i][j] != t[D-j+l][L+i-u] {
				ok0 = false
			}
			if U+j-l >= n+2 || R-i+u < 0 || s[i][j] != t[U+j-l][R-i+u] {
				ok1 = false
			}
			if s[i][j] != t[U+i-u][L+j-l] {
				ok2 = false
			}
			if s[i][j] != t[D-i+u][R-j+l] {
				ok3 = false
			}
		}
	}

	if ok0 || ok1 || ok2 || ok3 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
