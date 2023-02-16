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
	x := make([]int, n)
	y := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i], &y[i])
	}

	for i := 0; i < n; i++ {
		if (x[i]+y[i]+x[0]+y[0])%2 != 0 {
			fmt.Println(-1)
			return
		}
	}

	d := 0
	for i := 0; i < n; i++ {
		d = max(d, abs(x[i])+abs(y[i]))
	}

	t := 0
	for t = 0; (t*(t+1)/2+d)%2 != 0 || t*(t+1)/2 < d; t++ {

	}
	fmt.Println(t)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
