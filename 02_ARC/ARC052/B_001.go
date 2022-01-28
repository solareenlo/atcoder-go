package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, q int
	fmt.Fscan(in, &n, &q)

	x := make([]float64, n)
	r := make([]float64, n)
	h := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i], &r[i], &h[i])
	}

	for j := 0; j < q; j++ {
		var a, b float64
		fmt.Fscan(in, &a, &b)
		res := 0.0
		for i := 0; i < n; i++ {
			bottom := math.Max(x[i], a)
			top := math.Min(x[i]+h[i], b)
			if bottom < top {
				R := r[i] * (x[i] + h[i] - bottom) / h[i]
				res += math.Pi * R * R * (x[i] + h[i] - bottom) / 3.0
				R = r[i] * (x[i] + h[i] - top) / h[i]
				res -= math.Pi * R * R * (x[i] + h[i] - top) / 3.0
			}
		}
		fmt.Fprintln(out, res)
	}
}
