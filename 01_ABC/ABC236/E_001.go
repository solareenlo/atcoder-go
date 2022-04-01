package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const N = 100005

var (
	n int
	f = make([]float64, N)
	g = make([]float64, N)
	b = make([]float64, N)
)

func check() bool {
	f[0] = 0
	g[0] = 0
	for i := 1; i <= n; i++ {
		f[i] = math.Max(f[i-1], g[i-1]) + b[i]
		g[i] = f[i-1]
	}
	return math.Max(f[n], g[n]) > 0
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	a := make([]float64, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	r := 1e9
	l := 0.0
	ans1 := 0.0
	for r-l >= 5e-5 {
		mid := (l + r) / 2.0
		for i := 1; i <= n; i++ {
			b[i] = a[i] - mid
		}
		if check() {
			l = mid
			ans1 = mid
		} else {
			r = mid
		}
	}

	r = 1e9
	l = 0.0
	ans2 := 0.0
	for r-l >= 5e-5 {
		mid := (l + r) / 2.0
		for i := 1; i <= n; i++ {
			if a[i] >= mid {
				b[i] = 1.0
			} else {
				b[i] = -1.0
			}
		}
		if check() {
			l = mid
			ans2 = mid
		} else {
			r = mid
		}
	}

	fmt.Println(ans1)
	fmt.Println(int(math.Floor(ans2 + 0.5)))
}
