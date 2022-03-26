package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var xa, ya, xb, yb, n int
	fmt.Fscan(in, &xa, &ya, &xb, &yb, &n)

	type pair struct{ x, y int }
	p := make([]pair, n+1)
	w := int(1e9)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &p[i].x, &p[i].y)
		if xa > xb {
			p[i].x = w - p[i].x
		}
		if ya > yb {
			p[i].y = w - p[i].y
		}
	}

	if xa > xb {
		xa = w - xa
		xb = w - xb
	}
	if ya > yb {
		ya = w - ya
		yb = w - yb
	}
	sort.Slice(p, func(i, j int) bool {
		return p[i].x < p[j].x
	})

	t := 0
	a := make([]int, 200002)
	for i := 1; i <= n; i++ {
		if xa <= p[i].x && p[i].x <= xb && ya <= p[i].y && p[i].y <= yb {
			t++
			a[t] = p[i].y
		}
	}

	f := make([]int, 200002)
	for i := range f {
		f[i] = 1 << 60
	}
	for i := 1; i <= t; i++ {
		idx := lowerBound(f[1:t+1], a[i]) + 1
		f[idx] = a[i]
	}

	var i int
	for i = 1; i <= t && f[i] <= w; i++ {
	}

	ans := 100.0*float64(xb-xa+yb-ya) + (math.Pi*5.0-20.0)*float64(i-1)
	if i == min(xb-xa, yb-ya)+2 {
		ans += math.Pi * 5.0
	}
	fmt.Println(ans)
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
