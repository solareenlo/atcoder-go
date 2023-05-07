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

	var T int
	fmt.Fscan(in, &T)
	for i := 0; i < T; i++ {
		var N int
		fmt.Fscan(in, &N)
		m := N
		prod := 1
		for d := 2; d*d <= m; d++ {
			if m%d == 0 {
				p := 1
				for m%d == 0 {
					p *= d
					m /= d
				}
				if d == 2 {
					prod *= min(p/2, 4)
				} else {
					prod *= 2
				}
			}
		}
		if m > 1 {
			prod *= min(m-1, 2)
		}
		fmt.Fprintln(out, prod)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
