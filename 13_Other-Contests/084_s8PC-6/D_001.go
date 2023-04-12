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

	type P struct {
		x, y float64
	}

	var n int
	fmt.Fscan(in, &n)
	p := make([]P, n)
	for i := 0; i < n; i++ {
		var x, r float64
		fmt.Fscan(in, &x, &r)
		p[i] = P{x, r}
	}
	sort.Slice(p, func(i, j int) bool {
		if p[i].x == p[j].x {
			return p[i].y < p[j].y
		}
		return p[i].x < p[j].x
	})
	var f [100005]float64
	f[0] = p[0].y
	for i := 0; i < n-1; i++ {
		f[i+1] = calc(f[i], p[i+1].y, p[i+1].x-p[i].x)
	}
	b := p[n-1].y
	ans := f[n-1]
	for i := n - 2; i >= 0; i-- {
		ans = math.Max(ans, calc(b, f[i], p[i+1].x-p[i].x))
		b = calc(b, p[i].y, p[i+1].x-p[i].x)
	}
	fmt.Println(ans)
}

func calc(a, b, d float64) float64 {
	if a < d {
		return b
	}
	return math.Pow(math.Pow(a-d, 3)+math.Pow(b, 3), 1.0/3)
}
