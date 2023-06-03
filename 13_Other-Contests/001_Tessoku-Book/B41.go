package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var X, Y int
	fmt.Fscan(in, &X, &Y)
	x := make([]int, 0)
	y := make([]int, 0)
	for X > 1 || Y > 1 {
		x = append(x, X)
		y = append(y, Y)
		if X > Y {
			X = X - Y
		} else {
			Y = Y - X
		}
	}
	n := len(x)
	fmt.Println(n)
	for i := n - 1; i >= 0; i-- {
		fmt.Println(x[len(x)-1], y[len(y)-1])
		x = x[:len(x)-1]
		y = y[:len(y)-1]
	}
}
