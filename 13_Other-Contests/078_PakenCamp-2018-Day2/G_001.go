package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const MOD = 1000000007
const MAX = 1 << 17

func main() {
	in := bufio.NewReader(os.Stdin)

	type pair struct {
		x, y int
	}

	var N, M, K int
	fmt.Fscan(in, &N, &M, &K)
	GG := make([]pair, 2*M)
	for i := 0; i < M; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		GG[i] = pair{a, b}
		GG[i+M] = pair{b, a}
	}
	sort.Slice(GG, func(i, j int) bool {
		if GG[i].x == GG[j].x {
			return GG[i].y < GG[j].y
		}
		return GG[i].x < GG[j].x
	})

	G := make([][]pair, N)
	for _, g := range GG {
		G[g.x] = append(G[g.x], pair{g.y, 1})
	}

	deg := make([]int, N)
	in_V := make([]int, N)
	for i := range in_V {
		in_V[i] = 1
	}
	for i := 0; i < N; i++ {
		deg[i] = len(G[i])
	}

	dp_00 := make([]int, MAX)
	dp_01 := make([]int, MAX)
	dp_00[1] = 0
	dp_01[1] = 1
	for n := 2; n < MAX; n++ {
		dp_00[n] = (K - 1) * dp_01[n-1] % MOD
		dp_01[n] = (dp_00[n-1] + (K-2)*dp_01[n-1]) % MOD
	}

	ANS := 1
	que := make([]int, 0)
	for v := 0; v < N; v++ {
		if deg[v] <= 2 {
			que = append(que, v)
		}
	}
	for i := 0; i < len(que); i++ {
		v := que[i]
		var to, n int
		if deg[v] == 1 {
			for _, tmp := range G[v] {
				to = tmp.x
				n = tmp.y
				if in_V[to] != 0 {
					break
				}
			}
			ANS = (ANS * powMod(K-1, n)) % MOD
			deg[v] = 0
			in_V[v] = 0
			deg[to] -= 1
			if deg[to] == 2 {
				que = append(que, to)
			}
		} else if deg[v] == 2 {
			H := make([]pair, 0)
			for _, tmp := range G[v] {
				to := tmp.x
				n := tmp.y
				if in_V[to] != 0 {
					H = append(H, pair{to, n})
				}
			}
			a := H[0].x
			na := H[0].y
			b := H[1].x
			nb := H[1].y
			if a == b && b == v {
				ANS = (ANS * dp_00[na]) % MOD
				ANS = (ANS * K) % MOD
				in_V[v] = 0
				deg[v] = 0
				G[v] = make([]pair, 0)
				continue
			}
			G[a] = append(G[a], pair{b, na + nb})
			G[b] = append(G[b], pair{a, na + nb})
			G[v] = make([]pair, 0)
			deg[v] = 0
			in_V[v] = 0
		} else if deg[v] == 0 {
			ANS = (ANS * K) % MOD
			in_V[v] = 0
		}
	}
	flag := false
	for _, x := range in_V {
		if x != 0 {
			flag = true
			break
		}
	}
	if !flag {
		fmt.Println(ANS)
		return
	}

	V := make([]int, 0)
	for v := 0; v < N; v++ {
		if in_V[v] != 0 {
			V = append(V, v)
		}
	}
	N = len(V)
	A := make([][]int, 100)
	for i := range A {
		A[i] = make([]int, 3)
	}
	g := 0
	for i := 0; i < N; i++ {
		v := V[i]
		for _, tmp := range G[v] {
			to := tmp.x
			n := tmp.y
			if in_V[to] != 0 {
				to = lowerBound(V, to)
				A[g][0] = min(i, to)
				A[g][1] = max(i, to)
				A[g][2] = n
				g++
			}
		}
	}
	// A.resize(g)
	sort_key := make([]int, g)
	for i := 0; i < g; i++ {
		sort_key[i] = A[i][0]*1000000000000 + A[i][1]*1000000 + A[i][2]
	}
	idx := make([]int, g)
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return sort_key[idx[i]] < sort_key[idx[j]]
	})
	B := make([][]int, 0)
	for i := 0; i < g; i += 2 {
		B = append(B, A[idx[i]])
	}
	A = B

	dp := make([]int, 1<<N)
	dp[0] = 1
	full := (1 << N) - 1
	ans := 0
	for k := 0; k < 100; k++ {
		sum := 0
		for _, x := range dp {
			sum += x
		}
		if sum == 0 {
			break
		}
		ans += dp[full]
		newdp := make([]int, 1<<N)
		for s := 0; s < (1 << N); s++ {
			if dp[s] == 0 {
				continue
			}
			rest := full - s
			if rest == 0 {
				continue
			}
			min_bit := rest & -rest
			t := rest
			for t > 0 {
				if t&min_bit != 0 {
					coef := K - k
					for e := 0; e < len(A); e++ {
						a := A[e][0]
						b := A[e][1]
						n := A[e][2]
						if (t&(1<<a)) != 0 && (t&(1<<b)) != 0 {
							coef = coef * dp_00[n] % MOD
						}
						if (t&(1<<a)) != 0 && (s&(1<<b)) != 0 {
							coef = coef * dp_01[n] % MOD
						}
						if (s&(1<<a)) != 0 && (t&(1<<b)) != 0 {
							coef = coef * dp_01[n] % MOD
						}
					}
					newdp[s|t] += coef * dp[s] % MOD
					newdp[s|t] %= MOD
				}
				t = (t - 1) & rest
			}
		}
		dp = newdp
	}
	ans %= MOD
	fmt.Println(ANS * ans % MOD)
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
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

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		n /= 2
	}
	return res
}
