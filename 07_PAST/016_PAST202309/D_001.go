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

	a := make([]int, n)
	m := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		m = max(m, a[i])
	}

	for i := 0; i < n; i++ {
		fmt.Fprintf(out, "%d ", (a[i]*1000000000+m/2)/m)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
