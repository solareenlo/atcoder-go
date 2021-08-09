package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, x, s, f, px, res int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x, &s)
		f = max(f+px-x+s, s)
		res = max(res, f)
		px = x
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
