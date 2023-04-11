package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 200001

type pair struct {
	x, y int
}

var n, m int
var a [N][]pair
var s, t [N << 2]int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &m)
	for i := 1; i <= m; i++ {
		var x, y, z int
		fmt.Fscan(in, &x, &y, &z)
		a[y] = append(a[y], pair{x, z})
	}
	for i := 1; i <= n; i++ {
		add(1, 1, n, i, i, s[1])
		for _, it := range a[i] {
			add(1, 1, n, it.x, i, it.y)
		}
	}
	fmt.Println(max(s[1], 0))
}

func add(x, L, R, l, r, c int) {
	if l <= L && R <= r {
		s[x] += c
		t[x] += c
		return
	}
	mid := (L + R) >> 1
	if l <= mid {
		add(x<<1, L, mid, l, r, c)
	}
	if r > mid {
		add(x<<1|1, mid+1, R, l, r, c)
	}
	s[x] = max(s[x<<1], s[x<<1|1]) + t[x]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
