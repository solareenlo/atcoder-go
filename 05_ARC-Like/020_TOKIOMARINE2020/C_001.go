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

	var n, k int
	fmt.Fscan(in, &n, &k)

	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	for j := 0; j < k; j++ {
		b := make([]int, n)
		m := n * 2
		for i := 0; i < n; i++ {
			m = min(m, a[i])
			b[max(0, i-a[i])]++
			if i+a[i]+1 < n {
				b[i+a[i]+1]--
			}
		}
		if m >= n {
			break
		}
		for i := 0; i < n-1; i++ {
			b[i+1] += b[i]
		}
		copy(a, b)
	}

	for i := 0; i < n; i++ {
		fmt.Fprint(out, a[i], " ")
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
