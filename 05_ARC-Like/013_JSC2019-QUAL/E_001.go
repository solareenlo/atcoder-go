package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var fa = [200005]int{}

func find(x int) int {
	if x == fa[x] {
		return x
	}
	fa[x] = find(fa[x])
	return fa[x]
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, h, w int
	fmt.Fscan(in, &n, &h, &w)

	type edge struct{ x, y, d int }
	a := make([]edge, n+1)
	for i := 1; i < n+1; i++ {
		fmt.Fscan(in, &a[i].x, &a[i].y, &a[i].d)
		a[i].y += h
	}
	tmp := a[1:]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].d > tmp[j].d
	})
	for i := 0; i < n; i++ {
		a[i+1] = tmp[i]
	}

	c := make([]int, h+w+1)
	for i := 1; i < h+w+1; i++ {
		fa[i] = i
		c[i] = 1
	}

	ans := 0
	for i := 1; i < n+1; i++ {
		x := find(a[i].x)
		y := find(a[i].y)
		if x != y {
			fa[x] = y
			c[y] += c[x]
		}
		if c[y] != 0 {
			c[y]--
			ans += a[i].d
		}
	}
	fmt.Println(ans)
}
