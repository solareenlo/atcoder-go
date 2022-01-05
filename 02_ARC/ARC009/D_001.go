package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAXN = 1000 + 10
const MAXV = 50000 + 10
const MAX = 10000000

var (
	n    = [MAXN]int{}
	stk  = [MAXN]int{}
	cnt  int
	fa   = [MAXN]int{}
	C    = [MAXN][MAXN]int{}
	side = [MAXV]node{}
	M    = make(map[int]int)
	edge = make([]int, 0)
)

type node struct{ u, v, w int }

func found(x int) int {
	if fa[x] != x {
		fa[x] = found(fa[x])
	}
	return fa[x]
}

func dfs(c, x, pr, ed int) {
	if x > ed {
		was := 0
		for i := 1; i < cnt+1; i++ {
			u := side[stk[i]].u
			v := side[stk[i]].v
			u = found(u)
			v = found(v)
			was += side[stk[i]].w
			if u != v {
				fa[u] = v
			}
		}
		mk := 1
		rt := found(C[c][1])
		for i := 2; i < n[c]+1; i++ {
			if found(C[c][i]) != rt {
				mk = 0
				break
			}
		}
		if mk != 0 {
			M[was]++
			M[was] = min(M[was], MAX)
		}
		for i := 1; i < n[c]+1; i++ {
			fa[C[c][i]] = C[c][i]
		}
		return
	}
	for i := pr + 1; i < len(edge); i++ {
		cnt++
		stk[cnt] = edge[i]
		dfs(c, x+1, i, ed)
		cnt--
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var A, T, K int
	fmt.Fscan(in, &A, &T, &K)

	icd := [MAXN][MAXN]int{}
	for i := 1; i < A+1; i++ {
		fmt.Fscan(in, &n[i])
		for j := 1; j < n[i]+1; j++ {
			fmt.Fscan(in, &C[i][j])
			icd[i][C[i][j]] = 1
		}
	}

	var m int
	fmt.Fscan(in, &m)

	var total int
	for i := 1; i < m+1; i++ {
		var u, v, w int
		fmt.Fscan(in, &u, &v, &w)
		side[i] = node{u, v, w}
		total += w
	}

	f := [MAXV]int{}
	f[0] = 1
	for i := 1; i < A+1; i++ {
		for k := range M {
			delete(M, k)
		}
		edge = edge[:0]
		for j := 1; j < m+1; j++ {
			if icd[i][side[j].u] != 0 && icd[i][side[j].v] != 0 {
				edge = append(edge, j)
			}
		}
		for j := 1; j < n[i]+1; j++ {
			fa[C[i][j]] = C[i][j]
		}
		dfs(i, 1, -1, n[i]-1)
		for j := MAXV - 1; j >= 0; j-- {
			if f[j] == 0 {
				continue
			} else {
				for k, v := range M {
					f[j+k] += min(MAX, f[j]*v)
					f[j+k] = min(f[j+k], MAX)
				}
			}
			f[j] = 0
		}
	}

	for i := MAXV - 1; i >= 0; i-- {
		if f[i] < K {
			K -= f[i]
		} else {
			fmt.Println(total - i)
			return
		}
	}
	fmt.Println(-1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
