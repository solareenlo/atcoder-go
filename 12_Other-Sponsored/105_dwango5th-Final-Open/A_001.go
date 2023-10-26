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

	const MX = 100100

	var a, b [MX]int
	var eb, er [MX]bool

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	var s string
	fmt.Fscan(in, &s)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &a[i], &b[i])
		if s[b[i]-1] == 'R' {
			er[a[i]] = true
		} else if (k & 1) != 0 {
			eb[a[i]] = true
		}
		if s[a[i]-1] == 'R' {
			er[b[i]] = true
		} else if (k & 1) != 0 {
			eb[b[i]] = true
		}
	}
	if (k & 1) == 0 {
		for i := 0; i < m; i++ {
			if !er[a[i]] {
				eb[b[i]] = true
			}
			if !er[b[i]] {
				eb[a[i]] = true
			}
		}
	}
	for i := 1; i <= n; i++ {
		if eb[i] {
			fmt.Fprintln(out, "First")
		} else {
			fmt.Fprintln(out, "Second")
		}
	}
}
