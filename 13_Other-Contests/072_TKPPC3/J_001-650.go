package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

var out = bufio.NewWriter(os.Stdout)

const EPS = 1e-9

type circle struct {
	x, y, r float64
}

var ans []circle

func main() {
	defer out.Flush()

	var n int
	fmt.Scan(&n)

	x0 := -1e7
	y0 := -1e7
	for i := 0; i < n; i++ {
		ok := false
		for y := -2000; y <= 2000; y += 250 {
			if !ok {
				for x := -2000; x <= 2000; x += 250 {
					if !ok {
						if float64(y) < y0 || (float64(y) == y0 && float64(x) < x0) {
							continue
						}
						res := ask(float64(x), float64(y))
						if res > 0 {
							ok = true
							x0 = float64(x)
							y0 = float64(y)
						}
					}
				}
			}
		}

		var f func(float64, float64) float64
		f = func(x, y float64) float64 {
			lo := x - 250.0
			hi := x
			for i := 0; i < 40; i++ {
				mi := (lo + hi) / 2
				if ask(mi, y) > 0 {
					hi = mi
				} else {
					lo = mi
				}
			}
			return hi
		}

		y1 := y0 - 0.1
		y2 := y0
		y3 := y0 + 0.1
		x1 := f(x0, y1)
		x2 := f(x0, y2)
		x3 := f(x0, y3)

		xc := (y1-y2)*(x3*x3-x1*x1+y3*y3-y1*y1) - (y1-y3)*(x2*x2-x1*x1+y2*y2-y1*y1)
		xc /= 2*(x1-x2)*(y1-y3) - 2*(x1-x3)*(y1-y2)
		yc := (x1-x3)*(x2*x2-x1*x1+y2*y2-y1*y1) - (x1-x2)*(x3*x3-x1*x1+y3*y3-y1*y1)
		yc /= 2*(x1-x2)*(y1-y3) - 2*(x1-x3)*(y1-y2)
		ans = append(ans, circle{xc, yc, math.Hypot(x1-xc, y1-yc)})
	}

	sort.Slice(ans, func(C1, C2 int) bool {
		if math.Abs(ans[C1].x-ans[C2].x) > EPS {
			return ans[C1].x < ans[C2].x
		}
		if math.Abs(ans[C1].y-ans[C2].y) > EPS {
			return ans[C1].y < ans[C2].y
		}
		return false
	})

	fmt.Fprintln(out, "!")
	out.Flush()
	for i := 0; i < n; i++ {
		fmt.Fprintf(out, "%.f %.f %.f\n", ans[i].x, ans[i].y, ans[i].r)
	}
	out.Flush()
}

func ask(x, y float64) int {
	fmt.Fprintf(out, "? %.9f %.9f\n", x, y)
	out.Flush()
	var res int
	fmt.Scan(&res)
	for _, C := range ans {
		if (x-C.x)*(x-C.x)+(y-C.y)*(y-C.y) < C.r*C.r+EPS {
			res--
		}
	}
	return res
}
