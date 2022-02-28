package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	ans int
	a   = make([]int, 300_300)
	n   int
)

func solve() {
	for i := 2; i < n; i++ {
		ans = max(ans, abs(a[i+1]-a[i])-abs(a[i+1]-a[1]))
	}
	type pair struct{ x, y int }
	v := make([]pair, 0)
	for i := 1; i < n; i++ {
		if a[i] < a[i+1] {
			v = append(v, pair{a[i], a[i+1]})
		}
	}
	sort.Slice(v, func(i, j int) bool {
		return v[i].x < v[j].x || (v[i].x == v[j].x && v[i].y < v[j].y)
	})
	if len(v) == 0 {
		return
	}
	mx := v[0].y
	for i := 1; i < len(v); i++ {
		ans = max(ans, 2*(min(v[i].y, mx)-v[i].x))
		mx = max(mx, v[i].y)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	a = make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	s := 0
	for i := 1; i < n; i++ {
		s += abs(a[i] - a[i+1])
	}
	solve()
	a = reverseOrderInt(a)
	solve()
	fmt.Println(s - ans)
}

func reverseOrderInt(a []int) []int {
	n := len(a)
	res := make([]int, n)
	n = copy(res, a)
	for i, j := 1, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
