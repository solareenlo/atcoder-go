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
	A := make([][]int, n)
	for i := range A {
		A[i] = make([]int, n)
	}
	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			fmt.Fscan(in, &A[y][x])
		}
	}
	var Q int
	fmt.Fscan(in, &Q)
	for Q > 0 {
		Q--
		var cmd, y, x int
		fmt.Fscan(in, &cmd, &y, &x)
		x--
		y--
		if cmd&1 != 0 {
			A[y], A[x] = A[x], A[y]
		} else {
			fmt.Println(A[y][x])
		}
	}
}
