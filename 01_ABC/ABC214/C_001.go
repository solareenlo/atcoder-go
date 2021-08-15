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
	fmt.Scan(&n)

	s := make([]int, n*2)
	t := make([]int, n*2)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &t[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s[i])
	}
	for i := 1; i < 2*n; i++ {
		s[i%n] = min(s[i%n], s[(i-1)%n]+t[(i-1)%n])
	}
	for i := 0; i < n; i++ {
		fmt.Fprintln(out, s[i])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
