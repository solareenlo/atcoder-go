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

	type P struct{ a, b int }
	c := make([]P, 200009)
	x, y, z, w := 1<<60, 0, 1<<60, 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &c[i].a, &c[i].b)
		if c[i].a > c[i].b {
			c[i].a, c[i].b = c[i].b, c[i].a
		}
		x = min(x, c[i].a)
		y = max(y, c[i].b)
		z = min(z, c[i].b)
		w = max(w, c[i].a)
	}

	ans1 := (w - x) * (y - z)
	y -= x
	tmp := c[1 : n+1]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].a < tmp[j].a
	})

	ans2 := 1 << 60
	for i, j := 1, 1<<60; i < n; i++ {
		j = min(c[i].b, j)
		w = max(w, c[i].b)
		x = min(j, c[i+1].a)
		ans2 = min(ans2, w-x)
	}
	fmt.Println(min(ans1, ans2*y))
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
