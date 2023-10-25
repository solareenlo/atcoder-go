package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, x int
	fmt.Fscan(in, &n, &x)
	c, m := 0, 0
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		c += a * b
		m = max(m, b)
	}
	fmt.Println(c + x*m)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
