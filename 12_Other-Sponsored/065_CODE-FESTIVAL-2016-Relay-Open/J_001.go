package main

import (
	"bufio"
	"fmt"
	"os"
)

var M int
var x, y [170010]int

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Scan(&n)
	for i := 0; 6*i+1 < n; i++ {
		if 6*i+1 < n {
			for j := 0; j < n; j += 2 {
				put(6*i+1, j)
			}
		}
		if 6*i+3 < n {
			put(6*i+3, 0)
		}
		if 6*i+4 < n {
			for j := 1; j < n; j += 2 {
				put(6*i+4, j)
			}
		}
		if 6*i+5 < n {
			put(6*i+5, 0)
		}
	}
	if n%6 == 4 {
		for i := 2; i < n; i += 2 {
			put(n-1, i)
		}
	}
	if n%6 == 1 {
		for i := 2; i < n; i += 2 {
			put(n-2, i)
		}
	}
	fmt.Fprintln(out, M)
	for i := 1; i <= M; i++ {
		fmt.Fprintln(out, x[i], y[i])
	}
}

func put(a, b int) {
	M++
	x[M] = a
	y[M] = b
}
