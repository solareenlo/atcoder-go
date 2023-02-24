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

	var t int
	fmt.Fscan(in, &t)
	for t > 0 {
		t--
		var a, b, c float64
		fmt.Fscan(in, &a, &b, &c)
		if a == 0 {
			if b == 0 {
				if c != 0 {
					fmt.Fprintln(out, 0)
				} else {
					fmt.Fprintln(out, 3)
				}
			} else {
				fmt.Fprintf(out, "1 %.12f\n", -c*1.0/b)
			}
		} else {
			if a < 0 {
				a = -a
				b = -b
				c = -c
			}
			d := b*b - 4.0*a*c
			if math.Abs(d) < 1e-9 {
				fmt.Fprintf(out, "1 %.12f\n", -b/2.0/a)
			} else if d < 0 {
				fmt.Fprintln(out, 0)
			} else {
				var x, y float64
				if b > 0 {
					x = (-b - math.Sqrt(d)) / 2 / a
					y = c * 1.0 / a / x
				} else {
					y = (-b + math.Sqrt(d)) / 2 / a
					x = c * 1.0 / a / y
				}
				fmt.Fprintf(out, "2 %.12f %.12f\n", x, y)
			}
		}
	}
}
