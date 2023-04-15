package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = 1 << 30

	type pair struct {
		x, y int
	}

	var N int
	fmt.Fscan(in, &N)
	data := make([]pair, 0)
	for i := 0; i < N; i++ {
		var A, B int
		fmt.Fscan(in, &A, &B)
		data = append(data, pair{A, -B})
		data = append(data, pair{B, -A})
	}
	sort.Slice(data, func(i, j int) bool {
		if data[i].x == data[j].x {
			return data[i].y < data[j].y
		}
		return data[i].x < data[j].x
	})

	dp := make([]int, 2*N)
	for i := range dp {
		dp[i] = INF
	}
	for _, d := range data {
		idx := LowerBound(dp, -d.y)
		dp[idx] = -d.y
	}
	fmt.Println(LowerBound(dp, INF))
}

func LowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
