package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	var c float64
	fmt.Fscan(in, &n, &c)

	x := make([]float64, n)
	y := make([]float64, n)
	sum := 0.0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i], &y[i])
		sum += x[i]
	}

	x0 := sum / float64(n)
	ans := 0.0
	for i := 0; i < n; i++ {
		ans += (x0-x[i])*(x0-x[i]) + (c-y[i])*(c-y[i])
	}
	fmt.Println(ans)
}
