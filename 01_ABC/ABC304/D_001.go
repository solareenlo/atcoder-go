package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 200200

	type pair struct {
		x, y int
	}

	var p, q [N]int

	var h, w, n int
	fmt.Fscan(in, &h, &w, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &p[i], &q[i])
	}
	var a int
	fmt.Fscan(in, &a)
	x := make([]int, a+1)
	for i := 1; i <= a; i++ {
		fmt.Fscan(in, &x[i])
	}
	var b int
	fmt.Fscan(in, &b)
	y := make([]int, b+1)
	for i := 1; i <= b; i++ {
		fmt.Fscan(in, &y[i])
	}
	mp := make(map[pair]int)
	for i := 1; i <= n; i++ {
		mp[pair{lowerBound(x[1:], p[i]) + 1, lowerBound(y[1:], q[i]) + 1}]++
	}
	mn := int(1e9)
	mx := 0
	for _, val := range mp {
		mn = min(mn, val)
		mx = max(mx, val)
	}
	if len(mp) == (a+1)*(b+1) {
		fmt.Println(mn, mx)
	} else {
		fmt.Println(0, mx)
	}
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
