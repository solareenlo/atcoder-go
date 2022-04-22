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
	a := make([]float64, n)
	b := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i], &b[i])
	}
	c := make([]float64, m)
	d := make([]float64, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &c[i], &d[i])
	}

	const EPS = 1e-9
	l := 0.0
	r := 1000000.0
	for l+EPS < r {
		mid := (l + r) / 2.0
		e := make([]float64, n)
		for i := 0; i < n; i++ {
			e[i] = b[i] - a[i]*mid
		}
		f := make([]float64, m)
		for i := 0; i < m; i++ {
			f[i] = d[i] - c[i]*mid
		}
		sort.Slice(e, func(i, j int) bool {
			return e[i] > e[j]
		})
		sort.Slice(f, func(i, j int) bool {
			return f[i] > f[j]
		})
		tmp := 0.0
		for i := 0; i < 4; i++ {
			tmp += e[i]
		}
		tmp += max(e[4], f[0])
		if tmp > 0 {
			l = mid
		} else {
			r = mid
		}
	}
	fmt.Println(l)
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
