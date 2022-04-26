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

	var n, qq int
	fmt.Fscan(in, &n, &qq)
	r := make([]int, n)
	c := make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = n * i
		c[i] = i
	}

	for ; qq > 0; qq-- {
		var t int
		fmt.Fscan(in, &t)
		if t == 3 {
			r, c = c, r
		} else {
			var a, b int
			fmt.Fscan(in, &a, &b)
			a--
			b--
			if t == 1 {
				r[a], r[b] = r[b], r[a]
			} else if t == 2 {
				c[a], c[b] = c[b], c[a]
			} else if t == 4 {
				fmt.Fprintln(out, r[a]+c[b])
			}
		}
	}
}
