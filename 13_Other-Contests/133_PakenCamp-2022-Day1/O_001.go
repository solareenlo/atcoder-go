package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"sort"
)

type pair struct {
	x, y int
}

type tuple struct {
	x, y, z int
}

type point struct {
	x, y int
}

func (l point) minus(r point) point {
	return point{l.x - r.x, l.y - r.y}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int
	fmt.Fscan(in, &N)
	E := make([][]pair, N)
	for i := 0; i < N-1; i++ {
		var A, B, C int
		fmt.Fscan(in, &A, &B, &C)
		A--
		B--
		E[A] = append(E[A], pair{C, B})
		E[B] = append(E[B], pair{C, A})
	}
	ans := make([]int, N)
	used := make([]bool, N)
	sz := make([]int, N)
	var dfs func(int)
	dfs = func(s int) {
		id := make([]int, 0)
		var dfs_sz func(int, int)
		dfs_sz = func(v, u int) {
			id = append(id, v)
			sz[v] = 1
			for _, P := range E[v] {
				w := P.y
				if w != u && !used[w] {
					dfs_sz(w, v)
					sz[v] += sz[w]
				}
			}
		}
		dfs_sz(s, -1)
		g := -1
		for _, v := range id {
			if sz[v]*2 >= sz[s] {
				ok := true
				for _, P := range E[v] {
					w := P.y
					if !used[w] {
						if sz[w] < sz[v] && sz[w]*2 >= sz[s] {
							ok = false
						}
					}
				}
				if ok {
					g = v
				}
			}
		}
		C := make([]tuple, 0)
		S := make([]int, 1)
		var dfs2 func(int, int, int, int)
		dfs2 = func(v, cd, cs, u int) {
			if v != g {
				ans[g] = max(ans[g], cs/cd)
				ans[v] = max(ans[v], cs/cd)
				C = append(C, tuple{v, cd, cs})
			}
			for _, P := range E[v] {
				w := P.y
				if w != u && !used[w] {
					dfs2(w, cd+1, cs+P.x, v)
					if v == g {
						S = append(S, len(C))
					}
				}
			}
		}
		dfs2(g, 0, 0, -1)
		if len(S) >= 3 {
			cnt := len(S) - 1
			LOG := 32 - clz(uint32(cnt-1))
			for i := 0; i < LOG; i++ {
				P := make([][]point, 2)
				for j := 0; j < cnt; j++ {
					for k := S[j]; k < S[j+1]; k++ {
						P[(j>>i)&1] = append(P[(j>>i)&1], point{C[k].y, C[k].z})
					}
				}
				for j := 0; j < 2; j++ {
					P[j] = upper_hull(P[j])
				}
				for j := 0; j < cnt; j++ {
					for k := S[j]; k < S[j+1]; k++ {
						v := C[k].x
						ans[v] = max(ans[v], get_max_gradient(P[((j>>i)&1)^1], point{-C[k].y, -C[k].z}))
					}
				}
			}
		}
		used[g] = true
		for _, P := range E[g] {
			w := P.y
			if !used[w] {
				dfs(w)
			}
		}
	}
	dfs(0)
	for i := 0; i < N; i++ {
		fmt.Fprintln(out, ans[i])
	}
}

func upper_hull(P []point) []point {
	N := len(P)
	sort.Slice(P, func(i, j int) bool {
		return P[i].x < P[j].x
	})
	ans := make([]point, 0)
	ans = append(ans, P[0])
	for i := 1; i < N; i++ {
		if P[i].x == ans[len(ans)-1].x && P[i].y > ans[len(ans)-1].y {
			ans = ans[:len(ans)-1]
		} else if P[i].x == ans[len(ans)-1].x && P[i].y <= ans[len(ans)-1].y {
			continue
		}
		for len(ans) >= 2 {
			A := ans[len(ans)-2]
			B := ans[len(ans)-1]
			if cross(B.minus(A), P[i].minus(A)) >= 0 {
				ans = ans[:len(ans)-1]
			} else {
				break
			}
		}
		ans = append(ans, P[i])
	}
	return ans
}

func cross(P, Q point) int {
	return P.x*Q.y - P.y*Q.x
}

func get_max_gradient(P []point, A point) int {
	N := len(P)
	tv := 0
	fv := N
	for fv-tv > 1 {
		mid := (tv + fv) / 2
		if cross(P[mid-1].minus(A), P[mid].minus(A)) > 0 {
			tv = mid
		} else {
			fv = mid
		}
	}
	return (P[tv].y - A.y) / (P[tv].x - A.x)
}

func clz(x uint32) int {
	return bits.LeadingZeros32(x)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
