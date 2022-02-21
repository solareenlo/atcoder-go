package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	w := make([]int, n)
	p := make([]int, n)
	s := make([]int, 1<<8)
	maxi := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &w[i])
		maxi = max(maxi, w[i])
		p[i] = i
		for mask := 0; mask < 1<<n; mask++ {
			if mask>>i&1 != 0 {
				s[mask] += w[i]
			}
		}
	}

	e := make([]int, 1<<8)
	for i := 0; i < m; i++ {
		var l, v int
		fmt.Fscan(in, &l, &v)
		if v < maxi {
			fmt.Println(-1)
			return
		}
		for mask := 0; mask < 1<<n; mask++ {
			if s[mask] > v {
				e[mask] = max(e[mask], l)
			}
		}
	}

	ans := 1 << 60
	for {
		dp := make([]int, 8)
		for i := 0; i < n; i++ {
			mask := 1 << p[i]
			for j := i + 1; j < n; j++ {
				mask += 1 << p[j]
				dp[j] = max(dp[j], dp[i]+e[mask])
			}
		}
		ans = min(ans, dp[n-1])
		if !nextPermutation(sort.IntSlice(p)) {
			break
		}
	}

	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func nextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}
