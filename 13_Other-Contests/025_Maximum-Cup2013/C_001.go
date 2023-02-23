package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	IN := bufio.NewReader(os.Stdin)

	var N, sx, sy, ex, ey int
	fmt.Fscan(IN, &N, &sx, &sy, &ex, &ey)

	var in func(int, int, int) bool
	in = func(a, b, c int) bool {
		var t int
		if a > b {
			t = a
			a = b
			b = t
		}
		return a <= c && c <= b
	}

	var ccw func(int, int, int, int, int, int) int
	ccw = func(a, b, c, d, e, f int) int {
		p := a*d + c*f + e*b
		q := a*f + c*b + e*d
		if p > q {
			return 1
		}
		if p < q {
			return -1
		}
		return 0
	}

	A := 1
	for i := 1; i < N; i++ {
		var px, py, qx, qy int
		fmt.Fscan(IN, &px, &py, &qx, &qy)
		a := ccw(px, py, sx, sy, ex, ey)
		b := ccw(qx, qy, sx, sy, ex, ey)
		c := ccw(sx, sy, px, py, qx, qy)
		d := ccw(ex, ey, px, py, qx, qy)
		e := (c == 0) && in(px, qx, sx) && in(py, qy, sy)
		f := (d == 0) && in(px, qx, ex) && in(py, qy, ey)
		g := (a == 0) && in(sx, ex, px) && in(sy, ey, py)
		h := (b == 0) && in(sx, ex, qx) && in(sy, ey, qy)
		if e || f || g || h || (a*b < 0 && c*d < 0) {
			A++
		}
	}
	fmt.Println(A)
}
