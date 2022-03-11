package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, q int
	var s string
	fmt.Fscan(in, &n, &q, &s)

	t := make([]string, q+1)
	d := make([]string, q+1)
	for i := 1; i <= q; i++ {
		fmt.Fscan(in, &t[i], &d[i])
	}

	l, r := -1, n
	for i := q; i >= 1; i-- {
		if d[i] == "L" {
			if l < n-1 && s[l+1] == t[i][0] {
				l++
			}
			if r < n && s[r] == t[i][0] {
				r++
			}
		} else {
			if l >= 0 && s[l] == t[i][0] {
				l--
			}
			if r > 0 && s[r-1] == t[i][0] {
				r--
			}
		}
	}

	fmt.Println(max(0, r-l-1))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
