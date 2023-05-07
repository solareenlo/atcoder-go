package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M, Q int
	fmt.Fscan(in, &N, &M, &Q)

	edges := make([]pair, M)
	for i := range edges {
		fmt.Fscan(in, &edges[i].x, &edges[i].y)
	}
	sortPair(edges)

	A := make([]int, Q)
	B := make([]int, Q)
	for i := 0; i < Q; i++ {
		fmt.Fscan(in, &A[i], &B[i])
	}

	for i := 0; i < Q; i += 64 {
		dp := make([]uint, N+1)
		for j := 0; j < 64 && i+j < Q; j++ {
			dp[A[i+j]] |= uint(1 << j)
		}
		for _, tmp := range edges {
			dp[tmp.y] |= dp[tmp.x]
		}
		for j := 0; j < 64 && i+j < Q; j++ {
			if (dp[B[i+j]] & uint(1<<j)) != 0 {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	}
}

type pair struct {
	x, y int
}

func sortPair(tmp []pair) {
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].x == tmp[j].x {
			return tmp[i].y < tmp[j].y
		}
		return tmp[i].x < tmp[j].x
	})
}
