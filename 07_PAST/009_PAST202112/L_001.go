package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, p int
	fmt.Fscan(in, &n, &p)

	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	a = reverseOrderInt(a)

	for i := 0; i < n; i++ {
		a[i] -= i
	}

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1 << 60
	}
	for i := 0; i < n; i++ {
		if a[i]+n-1 <= p && a[i] >= 0 {
			idx := upperBound(dp, a[i])
			dp[idx] = a[i]
		}
	}

	idx := lowerBound(dp, 1<<60)
	fmt.Println(n - idx)
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}

func reverseOrderInt(a []int) []int {
	n := len(a)
	res := make([]int, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
