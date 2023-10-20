package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var x, y [15]float64

func main() {
	in := bufio.NewReader(os.Stdin)

	var a [15]float64
	var sum, f, g, pc [33333]float64

	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i], &y[i], &a[i])
	}
	for i := 1; i < (1 << n); i++ {
		pc[i] = pc[i-(i&(-i))] + 1
	}
	for i := 0; i < 1<<n; i++ {
		for j := 0; j < n; j++ {
			if (i & (1 << j)) != 0 {
				sum[i] += a[j]
			}
		}
	}
	for i := 1; i < (1 << n); i++ {
		f[i] = 1e18
		for j := 0; j < n; j++ {
			if (i & (1 << j)) != 0 {
				if i == (1 << j) {
					f[i] = 0
					continue
				}
				s := 1e18
				for k := 0; k < n; k++ {
					if j != k && (i&(1<<k)) != 0 {
						s = math.Min(s, dis(j, k))
					}
				}
				f[i] = math.Min(f[i], f[i^(1<<j)]+s)
			}
		}
	}
	g[0] = 1e18
	for i := 1; i < (1 << n); i++ {
		g[i] = -1e18
		for k := i; k > 0; k = ((k - 1) & i) {
			g[i] = math.Max(g[i], math.Min(g[i^k], (sum[k]-f[k])/pc[k]))
		}
	}
	fmt.Println(g[(1<<n)-1])
}

func dis(i, j int) float64 {
	return math.Sqrt((x[i]-x[j])*(x[i]-x[j]) + (y[i]-y[j])*(y[i]-y[j]))
}
