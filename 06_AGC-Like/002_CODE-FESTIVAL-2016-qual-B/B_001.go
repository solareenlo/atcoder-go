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

	var n, a, b int
	var s string
	fmt.Fscan(in, &n, &a, &b, &s)
	s = " " + s

	c1, c2 := 0, 0
	for i := 1; i <= n; i++ {
		if s[i] == 'a' {
			if c1 < a+b {
				c1++
				fmt.Fprintln(out, "Yes")
			} else {
				fmt.Fprintln(out, "No")
			}
		} else if s[i] == 'b' {
			if c1 < a+b && c2 < b {
				c1++
				c2++
				fmt.Fprintln(out, "Yes")
			} else {
				fmt.Fprintln(out, "No")
			}
		} else {
			fmt.Fprintln(out, "No")
		}
	}
}
