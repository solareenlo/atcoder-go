package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const N = 600005
	const INF = int(1e9)

	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, N)
	vis := make([]bool, N)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		vis[a[i]] = true
	}
	R := make([]int, N)
	L := make([]int, N)
	for i := 1; i < m+m; i += 2 {
		R[i] = INF
		L[i] = -1
	}
	p := 0
	for i := 1; i < m+m; i += 2 {
		p = -1
		for t, j := 0, i; j <= m+m && t <= R[i]; j, t = 2*j, t+1 {
			if vis[j] {
				p = t
			}
		}
		R[i] = p
		for j := 3 * i; j < m+m; j += 2 * i {
			R[j] = min(R[j], R[i]-1)
		}
	}
	for i := m + m - 1; i >= 1; i -= 2 {
		p = INF
		for j := 3 * i; j < m+m; j += 2 * i {
			L[i] = max(L[i], L[j]+1)
		}
		for t, j := 0, i; j <= m+m; j, t = 2*j, t+1 {
			if t >= L[i] && vis[j] {
				p = t
				break
			}
		}
		L[i] = p
	}
	op := true
	for i := 1; i < m+m; i++ {
		if L[i] > R[i] {
			op = false
		}
	}
	for i := 1; i <= n; i++ {
		t := 0
		for a[i]&1 == 0 {
			a[i] >>= 1
			t++
		}
		if op && t >= L[a[i]] && t <= R[a[i]] {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
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
