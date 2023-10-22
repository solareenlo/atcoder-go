package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const MAXN = 300111
	const LL_INF = 1 << 60
	const INF = 1061109567

	var N, K int
	fmt.Fscan(in, &N, &K)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}

	lis := make([]int, N+1)
	for i := range lis {
		lis[i] = INF
	}
	pos := make([]int, N)
	for i := 0; i < N; i++ {
		k := lowerBound(lis, A[i])
		pos[i] = k
		lis[k] = A[i]
	}

	Len := lowerBound(lis, lis[N])
	S := make([][]int, Len+1)
	for i := 0; i < Len+1; i++ {
		S[i] = append(S[i], 0)
	}

	var G [MAXN][]pair
	G[Len] = append(G[Len], pair{INF, N})
	S[Len] = append(S[Len], 1)

	way := make([]int, N)
	for i := N - 1; i >= 0; i-- {
		if len(G[pos[i]+1]) == 0 {
			continue
		}
		np := pos[i] + 1
		k := lowerBoundPair(G[np], pair{A[i], INF})
		var tmp int
		way[i] = min(LL_INF, S[np][len(S[np])-1]-S[np][k])
		if len(G[pos[i]]) == 0 || G[pos[i]][len(G[pos[i]])-1].x != A[i] {
			tmp = S[pos[i]][len(S[pos[i]])-1] + way[i]
		} else {
			x := lowerBoundPair(G[pos[i]], pair{A[i], -INF})
			tmp = S[pos[i]][x] + way[i]
		}

		G[pos[i]] = append(G[pos[i]], pair{A[i], i})
		S[pos[i]] = append(S[pos[i]], tmp)
	}

	ans := make([]int, Len)
	if S[0][len(S[0])-1] < K {
		fmt.Fprintln(out, "None")
	} else {
		last := -1
		last_i := -1
		for i := 0; i < Len; i++ {
			for j0, j1 := 0, 0; j0 < len(G[i]); {
				for j1 < len(G[i]) && G[i][j0].x == G[i][j1].x && G[i][j1].y > last_i {
					j1++
				}
				j := j1 - 1
				if last < G[i][j].x && last_i < G[i][j].y {
					if way[G[i][j].y] >= K {
						last = G[i][j].x
						ans[i] = G[i][j].x
						last_i = G[i][j].y
						break
					} else {
						K -= way[G[i][j].y]
					}
				}
				j0 = j1
			}
		}
		for i := 0; i < Len; i++ {
			fmt.Fprintf(out, "%d ", ans[i])
		}
		fmt.Fprintln(out)
	}
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

type pair struct {
	x, y int
}

func lowerBoundPair(a []pair, x pair) int {
	idx := sort.Search(len(a), func(i int) bool {
		if a[i].x == x.x {
			return a[i].y >= x.y
		}
		return a[i].x >= x.x
	})
	return idx
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
