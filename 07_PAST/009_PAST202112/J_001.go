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

	var N, Q int
	fmt.Fscan(in, &N, &Q)

	A := [300][300]int{}
	xa := 1
	var xb, xc, ya int
	yb := 1
	var yc int
	for i := 0; i < Q; i++ {
		var t int
		fmt.Fscan(in, &t)
		if t == 1 {
			var x, y int
			fmt.Fscan(in, &x, &y)
			x--
			y--
			A[xa*x+xb*y+xc][ya*x+yb*y+yc] ^= 1
		} else if t == 2 {
			var c string
			fmt.Fscan(in, &c)
			if c == "A" {
				za := xa
				zb := xb
				zc := xc
				xa = zb
				xb = -za
				xc = (N-1)*za + zc
				za = ya
				zb = yb
				zc = yc
				ya = zb
				yb = -za
				yc = (N-1)*za + zc
			} else {
				za := ya
				zb := yb
				zc := yc
				ya = -zb
				yb = za
				yc = (N-1)*zb + zc
				za = xa
				zb = xb
				zc = xc
				xa = -zb
				xb = za
				xc = (N-1)*zb + zc
			}
		} else {
			var c string
			fmt.Fscan(in, &c)
			if c == "A" {
				xc = (N-1)*xa + xc
				xa = -xa
				yc = (N-1)*ya + yc
				ya = -ya
			} else {
				xc = (N-1)*xb + xc
				xb = -xb
				yc = (N-1)*yb + yc
				yb = -yb
			}
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Fprint(out, A[xa*i+xb*j+xc][ya*i+yb*j+yc])
		}
		fmt.Fprintln(out)
	}
}
