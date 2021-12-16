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

	var h, w, n int
	fmt.Fscan(in, &h, &w, &n)

	r := make([]int, n+1)
	c := make([]int, n+1)
	a := make([]int, n+1)
	mp := map[int][]int{}
	for i := 1; i < n+1; i++ {
		fmt.Fscan(in, &r[i], &c[i], &a[i])
		mp[a[i]] = append(mp[a[i]], i)
	}

	keys := make([]int, 0, len(mp))
	for k, _ := range mp {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	dp := make([]int, 200002)
	rr := make([]int, 200002)
	cc := make([]int, 200002)
	for _, j := range keys {
		for _, k := range mp[j] {
			dp[k] = max(rr[r[k]], cc[c[k]])
		}
		for _, k := range mp[j] {
			rr[r[k]] = max(rr[r[k]], dp[k]+1)
			cc[c[k]] = max(cc[c[k]], dp[k]+1)
		}
	}

	for i := 1; i < n+1; i++ {
		fmt.Fprintln(out, dp[i])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
