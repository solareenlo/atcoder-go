package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	tmp := a[1:]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] < tmp[j]
	})

	var dfs func(int, int, int) int
	dfs = func(k, l, r int) int {
		if k < 0 {
			return 0
		}
		mid := r
		for i := l; i < r; i++ {
			if ((a[i] >> k) & 1) != 0 {
				mid = i
				break
			}
		}
		if mid == l || mid == r {
			return dfs(k-1, l, r)
		} else {
			return min(dfs(k-1, l, mid), dfs(k-1, mid, r)) + (1 << k)
		}
	}
	fmt.Println(dfs(30, 1, n+1))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
