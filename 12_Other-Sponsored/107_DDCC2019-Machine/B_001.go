package main

import (
	"fmt"
)

func main() {
	var xs, ys, xa, ya, xb, yb, xe, ye int
	fmt.Scan(&xs, &ys, &xa, &ya, &xb, &yb, &xe, &ye)

	for i := 0; i < 10; i++ {
		x1, y1, x2, y2 := xa, ya, xb, yb

		if x1-3*(2*i) >= 0 {
			x1 -= 3 * (2 * i)
		} else if x1+3*(2*i) <= 900 {
			x1 += 3 * (2 * i)
		} else if y1-3*(2*i) >= 0 {
			y1 -= 3 * (2 * i)
		} else if y1+3*(2*i) <= 900 {
			y1 += 3 * (2 * i)
		}

		if x2-3*(2*i+1) >= 0 {
			x2 -= 3 * (2*i + 1)
		} else if x2+3*(2*i+1) <= 900 {
			x2 += 3 * (2*i + 1)
		} else if y2-3*(2*i+1) >= 0 {
			y2 -= 3 * (2*i + 1)
		} else if y2+3*(2*i+1) <= 900 {
			y2 += 3 * (2*i + 1)
		}

		fmt.Printf("%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d\n", x1, 500, 60, 60, y1, 500, 60, 60, 2500, 0, 0)
		fmt.Printf("%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d\n", x2, 500, 60, 60, y2, 500, 60, 60, 0, 2500, 0)
		fmt.Printf("%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d\n", xe, 500, 60, 60, ye, 500, 60, 60, 0, 0, 5000)
	}
}
