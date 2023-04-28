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

	const MOD = 1000000007

	dd := []int{0, 1, 0, -1}

	var n, t, q, si, sj, s1, s2, s3 int
	fmt.Fscan(in, &n, &t, &q, &si, &sj, &s1, &s2, &s3)
	board := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &board[i])
	}
	v := make([]int, 0)
	inv := make([]int, n*n)
	for i := range inv {
		inv[i] = -1
	}
	inv[si*n+sj] = 0
	v = append(v, si*n+sj)
	for b := 0; b < q; b++ {
		var t int
		fmt.Fscan(in, &t)
		if t == 1 || t == 2 {
			var i, j int
			fmt.Fscan(in, &i, &j)
			if inv[i*n+j] == -1 {
				inv[i*n+j] = len(v)
				v = append(v, i*n+j)
			}
		} else {
			var k int
			fmt.Fscan(in, &k)
			for a := 0; a < k; a++ {
				var i, j int
				fmt.Fscan(in, &i, &j)
				if inv[i*n+j] == -1 {
					inv[i*n+j] = len(v)
					v = append(v, i*n+j)
				}
			}
		}
	}
	m := len(v)
	cost := make([][]int, m)
	for i := range cost {
		cost[i] = make([]int, m)
	}
	for _, s := range v {
		for _, t := range v {
			if inv[t] != -1 {
				cost[inv[s]][inv[t]] = abs(s/n-t/n) + abs(s%n-t%n)
			}
		}
	}

	res := solve(cost)
	ans := make([]int, 0)
	for i := 1; i < len(res); i++ {
		s := v[res[i]]
		dist := make([]int, n*n)
		for i := range dist {
			dist[i] = MOD
		}
		dist[s] = 0
		que := make([]int, 0)
		que = append(que, s)
		for len(que) > 0 {
			p := que[0]
			que = que[1:]
			if p == v[res[i-1]] {
				break
			}
			for k := 0; k < 4; k++ {
				ti := p/n + dd[k]
				tj := p%n + dd[k^1]
				if 0 <= ti && ti < n && 0 <= tj && tj < n && board[ti][tj] == '.' && dist[ti*n+tj] > dist[p]+1 {
					dist[ti*n+tj] = dist[p] + 1
					que = append(que, ti*n+tj)
				}
			}
		}
		p := v[res[i-1]]
		for p != s {
			for k := 0; k < 4; k++ {
				ti := p/n + dd[k]
				tj := p%n + dd[k^1]
				if 0 <= ti && ti < n && 0 <= tj && tj < n && board[ti][tj] == '.' && dist[ti*n+tj] == dist[p]-1 {
					p = ti*n + tj
					ans = append(ans, p)
					break
				}
			}
		}
	}
	for i := 0; i < t; i++ {
		p := ans[min(i, len(ans)-1)]
		fmt.Fprintln(out, p/n, p%n, p/n, p%n)
	}
}

func solve(cost [][]int) []int {
	s := 0
	n := len(cost)
	vis := make([]bool, n)
	res := make([]int, n)
	for i := range res {
		res[i] = -1
	}
	res[0] = s
	vis[s] = true
	for i := 1; i < n; i++ {
		prev := res[i-1]
		mi := -1
		for j := 0; j < n; j++ {
			if !vis[j] && (res[i] == -1 || mi > cost[prev][j]) {
				mi = cost[prev][j]
				res[i] = j
			}
		}
		vis[res[i]] = true
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
