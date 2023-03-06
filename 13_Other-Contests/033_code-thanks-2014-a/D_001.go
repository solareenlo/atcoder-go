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

	var n, q int
	fmt.Fscan(in, &n, &q)
	for i := 0; i < q; i++ {
		var a, b, l, r int
		fmt.Fscan(in, &a, &b, &l, &r)
		fmt.Fprintf(out, "%d\n", (r-l)*100-100*max(0, min(b, r)-max(a, l)))
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
