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

	type P struct{ x, y int }
	l := make([]P, n)
	r := make([]P, n)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		if a < b {
			l = append(l, P{a, b})
		} else {
			r = append(r, P{b, a})
		}
	}

	sort.Slice(l, func(i, j int) bool {
		return l[i].x < l[j].x

	})
	sort.Slice(r, func(i, j int) bool {
		return r[i].x > r[j].x
	})

	s, x := 0, 0
	for _, p := range l {
		x = max(x, s+p.x)
		s += p.x - p.y
	}
	for _, p := range r {
		x = max(x, s+p.y)
		s += p.y - p.x
	}
	fmt.Println(x)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
