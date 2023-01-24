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

	var n, q int
	var s string
	fmt.Fscan(in, &n, &q, &s)

	var op, x int
	idx := 0
	for q > 0 {
		fmt.Fscan(in, &op, &x)
		if op == 2 {
			fmt.Fprintf(out, "%c\n", s[(x-1+idx+n)%n])
		} else {
			idx = (idx - x + n) % n
		}
		q--
	}
}
