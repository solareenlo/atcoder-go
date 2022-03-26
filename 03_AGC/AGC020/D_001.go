package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var q int
	fmt.Fscan(in, &q)

	for j := 0; j < q; j++ {
		var a, b, c, d int
		fmt.Fscan(in, &a, &b, &c, &d)
		k := (a + b) / (min(a, b) + 1)
		l := 0
		r := a + b
		for l < r {
			m := (l + r) >> 1
			x := a - m/(k+1)*k - m%(k+1)
			y := b - m/(k+1)
			if y <= x*k {
				l = m + 1
			} else {
				r = m
			}
		}
		x := a - l/(k+1)*k - l%(k+1)
		y := b - l/(k+1)
		r = l + 1 + y - x*k
		for i := c; i <= d; i++ {
			if i <= l {
				if i%(k+1) != 0 {
					fmt.Fprint(out, "A")
				} else {
					fmt.Fprint(out, "B")
				}
			} else {
				if (i-r)%(k+1) != 0 {
					fmt.Fprint(out, "B")
				} else {
					fmt.Fprint(out, "A")
				}
			}
		}
		fmt.Fprintln(out)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
