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
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i], &y[i])
	}
	x[n] = x[0]
	y[n] = y[0]

	var a, b int
	fmt.Fscan(in, &a, &b)
	ans := false
	for i := 0; i < n; i++ {
		px := x[i] - a
		py := y[i] - b
		qx := x[i+1] - a
		qy := y[i+1] - b
		if qy < py {
			px, qx = qx, px
			py, qy = qy, py
		}
		if px*qy-py*qx > 0 && py <= 0 && qy > 0 {
			ans = !ans
		}
	}

	if ans {
		fmt.Println("INSIDE")
	} else {
		fmt.Println("OUTSIDE")
	}
}
