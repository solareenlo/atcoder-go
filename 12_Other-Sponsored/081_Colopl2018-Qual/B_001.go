package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, x int
	var s string
	fmt.Fscan(in, &n, &x)
	fmt.Fscan(in, &s)
	c := 0
	for i := 0; i < n; i++ {
		var t int
		fmt.Fscan(in, &t)
		if s[i] == '0' {
			c += t
		} else {
			c += min(x, t)
		}
	}
	fmt.Println(c)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
