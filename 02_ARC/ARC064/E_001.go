package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 1002
	x := make([]float64, N)
	y := make([]float64, N)
	var n int
	fmt.Fscan(in, &x[0], &y[0], &x[1], &y[1], &n)

	r := make([]float64, N)
	for i := 2; i < n+2; i++ {
		fmt.Fscan(in, &x[i], &y[i], &r[i])
	}

	a := [N][N]float64{}
	for i := 1; i < n+2; i++ {
		for j := 0; j < i; j++ {
			a[i][j] = math.Sqrt((x[i]-x[j])*(x[i]-x[j])+(y[i]-y[j])*(y[i]-y[j])) - r[i] - r[j]
			if a[i][j] < 0 {
				a[i][j] = 0
			}
			a[j][i] = a[i][j]
		}
	}

	d := make([]float64, N)
	for i := 0; i < n+2; i++ {
		d[i] = a[0][i]
	}

	k := 0
	f := make([]bool, N)
	f[0] = true
	for i := 0; i < n; i++ {
		for t, j := 2e9, 1; j < n+2; j++ {
			if !f[j] && d[j] < t {
				k = j
				t = d[j]
			}
		}
		f[k] = true
		for j := 1; j < n+2; j++ {
			if d[k]+a[k][j] < d[j] {
				d[j] = d[k] + a[k][j]
			}
		}
	}
	fmt.Println(d[1])
}
