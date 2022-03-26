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
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	x := make([]float64, n)
	y := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i], &y[i])
	}

	p := make([]float64, 210)
	for i := 0; i < n; i++ {
		tot := 0
		for j := 0; j < n; j++ {
			if i != j {
				p[tot] = math.Atan2(y[j]-y[i], x[j]-x[i])
				tot++
			}
		}
		tmp := p[:tot]
		sort.Float64s(tmp)
		for j := 0; j < tot; j++ {
			p[j+tot] = p[j] + 2*math.Pi
		}
		ans := 0.0
		for j := 0; j < tot; j++ {
			ans = math.Max(ans, math.Pi-p[j+tot-1]+p[j])
		}
		fmt.Fprintln(out, ans/2/math.Pi)
	}
}
