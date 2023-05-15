package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const Inf = 1000000001

	var N, M, K, T int
	fmt.Fscan(in, &N, &M, &K, &T)
	u := make([]int, M)
	v := make([]int, M)
	l := make([]int, M)
	d := make([][]int, N)
	for i := range d {
		d[i] = make([]int, N)
		for j := range d[i] {
			d[i][j] = Inf
		}
	}
	for i := 0; i < N; i++ {
		d[i][i] = 0
	}
	for i := 0; i < M; i++ {
		fmt.Fscan(in, &u[i], &v[i], &l[i])
		u[i]--
		v[i]--
		d[u[i]][v[i]] = l[i]
		d[v[i]][u[i]] = l[i]
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			for k := 0; k < N; k++ {
				d[j][k] = min(d[j][k], d[j][i]+d[i][k])
			}
		}
	}

	s := make([]int, K)
	g := make([]int, K)
	m := make([]int, K)
	t := make([]int, 0)
	for i := 0; i < K; i++ {
		fmt.Fscan(in, &s[i], &g[i], &m[i])
		s[i]--
		g[i]--
		t = append(t, i)
	}
	sort.Slice(t, func(a, b int) bool {
		return d[s[t[a]]][g[t[a]]]*m[t[b]] < d[s[t[b]]][g[t[b]]]*m[t[a]]
	})

	ax := make([]int, 0)
	at := make([]int, 0)
	aa := make([][]int, 0)
	type pair struct {
		x, y int
	}
	us := make([][]pair, M)
	for i := 0; i < len(t); i++ {
		ind := t[i]
		ss := s[ind]
		gg := g[ind]
		tmp := make([]int, 1)
		tmp[0] = ss
		es := make([]int, 0)
		cc := ss
		for cc != gg {
			for j := 0; j < M; j++ {
				if u[j] != cc && v[j] != cc {
					continue
				}
				to := u[j] ^ v[j] ^ cc
				if l[j]+d[ss][cc]+d[to][gg] != d[ss][gg] {
					continue
				}
				tmp = append(tmp, to)
				es = append(es, j)
				cc = to
				break
			}
		}
		ok := T
		for tt := 0; tt < T+1; tt++ {
			mid := tt
			mm := mid
			flag := true
			for j := 0; j < len(es); j++ {
				ll := mm
				rr := mm + l[es[j]]
				for k := 0; k < len(us[es[j]]); k++ {
					A := max(ll, us[es[j]][k].x)
					B := min(rr, us[es[j]][k].y)
					if A < B {
						flag = false
						break
					}
				}
				mm = rr
			}
			if flag {
				ok = tt
				break
			}
		}
		if ok+d[ss][gg] > T {
			continue
		}

		aa = append(aa, tmp)
		at = append(at, ok)
		ax = append(ax, ind)
		mm := ok
		for j := 0; j < len(es); j++ {
			ll := mm
			rr := mm + l[es[j]]
			us[es[j]] = append(us[es[j]], pair{ll, rr})
			mm = rr
		}
	}
	fmt.Println(len(aa))
	for i := 0; i < len(aa); i++ {
		fmt.Println(ax[i]+1, at[i], len(aa[i])-1)
		for j := 0; j < len(aa[i]); j++ {
			if j != 0 {
				fmt.Print(" ")
			}
			fmt.Print(aa[i][j] + 1)
		}
		fmt.Println()
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
