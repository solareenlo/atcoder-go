package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M, D int
	fmt.Fscan(in, &N, &M, &D)
	a := make([]int, N)
	b := make([]int, M)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < M; i++ {
		fmt.Fscan(in, &b[i])
	}
	sort.Ints(a)
	sort.Ints(b)
	ans := -1
	for i := 0; i < N; i++ {
		A := lowerBound(b, a[i]-D)
		B := upperBound(b, a[i]+D)
		if B-A > 0 {
			ans = max(ans, a[i]+b[B-1])
		}
	}
	fmt.Println(ans)
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
