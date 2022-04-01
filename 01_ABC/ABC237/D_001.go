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
	var s string
	fmt.Fscan(in, &n, &s)

	for i := 0; i < n; i++ {
		if s[i] == 'R' {
			fmt.Fprint(out, i, " ")
		}
	}

	fmt.Fprint(out, n, " ")
	for i := n - 1; i >= 0; i-- {
		if s[i] == 'L' {
			fmt.Fprint(out, i, " ")
		}
	}
}
