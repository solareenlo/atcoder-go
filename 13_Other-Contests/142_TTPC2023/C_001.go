package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var solve func() int
	solve = func() int {
		var n int
		fmt.Fscan(in, &n)
		q := int(math.Sqrt(float64(n)))
		for (q+1)*(q+1) <= n {
			q++
		}
		for q*q > n {
			q--
		}
		r := n - q*q
		return min(r, q) + (q-1)*(q-1) + min(r, q-1)
	}

	var t int
	fmt.Fscan(in, &t)

	for t > 0 {
		t--
		fmt.Fprintln(out, solve())
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
