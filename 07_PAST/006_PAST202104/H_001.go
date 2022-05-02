package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M, K, Q int
	fmt.Fscan(in, &N, &M, &K, &Q)

	P := make([][]int, 2)
	for i := 0; i < N; i++ {
		var p, t int
		fmt.Fscan(in, &p, &t)
		P[t] = append(P[t], p)
	}

	S := make([][]int, 2)
	for i := 0; i < 2; i++ {
		sort.Ints(P[i])
		S[i] = make([]int, len(P[i])+1)
		for j := 0; j < len(P[i]); j++ {
			S[i][j+1] = S[i][j] + P[i][j]
		}
	}

	ans := 1 << 60
	for L := 0; L <= M; L++ {
		R := M - L
		if L > len(P[0]) || R > len(P[1]) {
			continue
		}
		now := S[0][L] + S[1][R]
		now += (R + K - 1) / K * Q
		ans = min(ans, now)
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
