package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, K, M, R int
	fmt.Fscan(in, &N, &K, &M, &R)
	a := make([]int, N-1)
	for i := 0; i < N-1; i++ {
		fmt.Fscan(in, &a[i])
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	ans := R * K
	for i := 0; i < K-1; i++ {
		ans -= a[i]
	}
	if N != K && a[K-1] >= ans {
		fmt.Println(0)
		return
	}
	if M >= ans {
		fmt.Println(max(0, ans))
		return
	}
	fmt.Println(-1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
