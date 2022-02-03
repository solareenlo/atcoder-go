package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, r int
	fmt.Fscan(in, &n, &r)

	t, s := -1<<60, -1<<60
	for i := 0; i < n-1; i++ {
		var c string
		var a int
		fmt.Fscan(in, &c, &a)
		if c == "-" {
			s = t + a
			t = max(r-a, s)
			r = t
		} else {
			r += a
			s += a
			t = max(t-a, s)
		}
	}
	fmt.Println(r)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
