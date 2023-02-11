package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 100100

	var n int
	fmt.Fscan(in, &n)

	t := make([]int, N)
	x := make([]int, N)
	y := make([]int, N)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &t[i+1], &x[i+1], &y[i+1])
	}

	ok := true
	for i := 0; i < n; i++ {
		diff_t := t[i+1] - t[i]
		dist := abs(x[i+1]-x[i]) + abs(y[i+1]-y[i])
		if diff_t < dist {
			ok = false
		}
		if diff_t%2 != dist%2 {
			ok = false
		}
	}

	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
