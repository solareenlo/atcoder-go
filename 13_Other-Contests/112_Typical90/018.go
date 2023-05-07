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

	const pi = 3.14159265358979
	const rad = 57.2957795131

	var t, l, x, y, q float64
	fmt.Fscan(in, &t, &l, &x, &y, &q)
	for q > 0 {
		q--
		var e int
		fmt.Fscan(in, &e)
		Z := l * (1 - math.Cos(2*pi*float64(e)/t)) / 2
		Y := -1 * math.Sin(2*pi*float64(e)/t) * l / 2
		fmt.Fprintln(out, rad*math.Atan(Z/math.Sqrt((Y-y)*(Y-y)+x*x)))
	}
}
