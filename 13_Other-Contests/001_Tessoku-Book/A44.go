package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, q int
	fmt.Fscan(in, &n, &q)
	A := make([]int, n+1)
	for i := 1; i <= n; i++ {
		A[i] = i
	}
	b := 0
	for i := 1; i <= q; i++ {
		var c int
		fmt.Fscan(in, &c)
		if c == 1 {
			var x, y int
			fmt.Fscan(in, &x, &y)
			A[abs(b-x)] = y
		} else if c == 2 {
			if b == 0 {
				b = n + 1
			} else {
				b = 0
			}
		} else {
			var x int
			fmt.Fscan(in, &x)
			fmt.Println(A[abs(b-x)])
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
