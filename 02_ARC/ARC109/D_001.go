package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var xa, ya, xb, yb, xc, yc int
		fmt.Fscan(in, &xa, &ya, &xb, &yb, &xc, &yc)

		X := -max(xa, xb, xc) + xa + xb + xc
		Y := -max(ya, yb, yc) + ya + yb + yc
		tmp := 0
		if X == Y && X != 0 && X != 1 {
			tmp = 1
		}
		fmt.Println(max(abs(X), abs(Y)) + tmp)
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a ...int) int {
	res := a[0]
	for i := range a {
		if res < a[i] {
			res = a[i]
		}
	}
	return res
}
