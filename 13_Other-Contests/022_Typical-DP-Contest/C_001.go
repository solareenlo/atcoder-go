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

	var k int
	fmt.Fscan(in, &k)

	n := 1 << k
	r := make([]float64, 1050)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &r[i])
	}

	d := [15][1050]float64{}
	for i := 0; i < n; i++ {
		d[0][i] = 1
	}

	for i := 1; i <= k; i++ {
		for j := 0; j < n; j++ {
			for l := 0; l < (1 << (i - 1)); l++ {
				t := 1 << (i - 1)
				q := l + (j ^ t) - (j & (t - 1))
				d[i][j] += d[i-1][j] * d[i-1][q] / (1 + math.Pow(10.0, (r[q]-r[j])/400.0))
			}
		}
	}

	for i := 0; i < n; i++ {
		fmt.Fprintln(out, d[k][i])
	}
}
