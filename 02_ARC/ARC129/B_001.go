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

	var n int
	fmt.Fscan(in, &n)

	miniR := 1 << 60
	maxL := 1
	for i := 0; i < n; i++ {
		var l, r int
		fmt.Fscan(in, &l, &r)
		maxL = max(maxL, l)
		miniR = min(miniR, r)
		fmt.Fprintln(out, max(maxL-miniR+1, 0)/2)
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
