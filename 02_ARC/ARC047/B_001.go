package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	x := make([]int, n+1)
	y := make([]int, n+1)
	const inf = 1 << 60
	lx, rx, ly, ry := inf, -inf, inf, -inf
	for i := 1; i < n+1; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		x[i] = a - b
		lx = min(lx, x[i])
		rx = max(rx, x[i])
		y[i] = a + b
		ly = min(ly, y[i])
		ry = max(ry, y[i])
	}

	k := rx - lx - ry + ly
	if k != 0 {
		flag := 0
		if k < 0 {
			for i := 1; i <= n; i++ {
				if x[i] == lx && y[i] != ly && y[i] != ry {
					flag |= 1
				} else {
					flag |= 0
				}
			}
			if flag != 0 {
				rx -= k
			} else {
				lx += k
			}
		} else {
			for i := 1; i <= n; i++ {
				if y[i] == ly && x[i] != lx && x[i] != rx {
					flag |= 1
				} else {
					flag |= 0
				}
			}
			if flag != 0 {
				ly += k
			} else {
				ry -= k
			}
		}
	}
	ax := (lx + rx) / 2
	ay := (ly + ry) / 2
	fmt.Println((ax+ay)/2, (ay-ax)/2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
