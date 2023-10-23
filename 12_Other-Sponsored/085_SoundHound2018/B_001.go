package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, l, r int
	fmt.Fscan(in, &n, &l, &r)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		fmt.Fprint(out, min(r, max(l, a)), " ")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
