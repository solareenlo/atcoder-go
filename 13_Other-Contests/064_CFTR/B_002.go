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
	fmt.Scan(&n, &q)

	for i := 0; i < q; i++ {
		var v, w int
		fmt.Fscan(in, &v, &w)
		if n > 1 {
			for v != w {
				if v > w {
					v, w = w, v
				}
				w = (w - 1 + n - 1) / n
			}
		}
		fmt.Fprintln(out, v)
	}
}
