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
	a := make([]int, n+3)
	for i := 3; i <= n+2; i++ {
		fmt.Fscan(in, &a[i])
	}

	s := make([]int, n+3)
	b := make([]int, 4)
	for i := 4; i <= n+2; i++ {
		s[i] = s[i-3] + (a[i] - a[i-1])
		if s[i] < 0 {
			b[i%3] = max(b[i%3], -s[i])
		}
	}

	if b[0]+b[1]+b[2] > a[3] {
		fmt.Fprintln(out, "No")
		return
	}

	fmt.Fprintln(out, "Yes")
	b[0] += a[3] - (b[0] + b[1] + b[2])
	for i := 1; i <= n+2; i++ {
		fmt.Fprint(out, b[i%3]+s[i], " ")
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
