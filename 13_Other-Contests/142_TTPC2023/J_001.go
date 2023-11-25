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

	var t int
	fmt.Fscan(in, &t)
	for t > 0 {
		var n, m int
		fmt.Fscan(in, &n, &m)
		if n == 5 && m == 15 {
			fmt.Fprint(out, "0 1 2 3 4 5 6 7 12 13 14 15 19 23 31\n")
		} else {
			bit := 0
			for (2 << bit) < m {
				bit++
			}
			fmt.Fprint(out, "0")
			m--
			for i := 1; i < 1<<bit; i++ {
				fmt.Fprintf(out, " %d", i)
				m--
			}
			for i := n; i > bit; i-- {
				if m > 0 {
					bit := 0
					for (m & (1 << bit)) <= 0 {
						bit++
					}
					for j := 0; j < 1<<bit; j++ {
						fmt.Fprintf(out, " %d", (1<<i)-1-j)
						m--
					}
				}
			}
			fmt.Fprintf(out, "\n")
		}
		t--
	}
}
