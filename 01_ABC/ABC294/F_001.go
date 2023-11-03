package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 50009

	var a, b, c, d [N]float64

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i], &b[i])
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &c[i], &d[i])
	}
	l, r := 0.0, 1.0
	v := make([]float64, m)
	for r-l > 1e-12 {
		x := (l + r) / 2
		z := x / (1 - x)
		sum := 0
		for i := 0; i < m; i++ {
			v[i] = c[i] - d[i]*z
		}
		sort.Slice(v, func(i, j int) bool {
			return v[i] < v[j]
		})
		for i := 0; i < n; i++ {
			w := a[i] - b[i]*z
			lb := lowerBound(v, -w)
			sum += m - lb
		}
		if sum < k {
			r = x
		} else {
			l = x
		}
	}
	fmt.Printf("%.12f\n", r*100)
}

func lowerBound(a []float64, x float64) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
