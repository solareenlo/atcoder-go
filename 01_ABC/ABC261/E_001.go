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

	var n, x int
	fmt.Fscan(in, &n, &x)

	p := (1 << 30) - 1
	q := 0
	for i := 1; i <= n; i++ {
		var t, a int
		fmt.Fscan(in, &t, &a)
		if t == 1 {
			p &= a
			q &= a
		} else if t == 2 {
			p |= a
			q |= a
		} else {
			p ^= a
			q ^= a
		}
		x = (x & p) | ((^x) & q)
		fmt.Fprintln(out, x)
	}
}
