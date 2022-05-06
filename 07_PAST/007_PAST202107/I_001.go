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

	var n int
	fmt.Fscan(in, &n)
	var x1, y1, x2, y2 float64
	fmt.Fscan(in, &x1, &y1, &x2, &y2)

	mx := (x1 + x2) / 2
	my := (y1 + y2) / 2
	a := make([]float64, n)
	b := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i], &b[i])
		a[i] -= mx
		b[i] -= my
	}

	x1 -= mx
	x2 -= mx
	y1 -= my
	y2 -= my
	p := -math.Atan2(y2, x2)
	c := math.Cos(p)
	s := math.Sin(p)
	for i := 0; i < n; i++ {
		fmt.Fprintln(out, c*a[i]-s*b[i], s*a[i]+c*b[i])
	}
}
