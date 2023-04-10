package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var T int
	fmt.Scan(&T)
	for T > 0 {
		T--
		var n, m int
		fmt.Scan(&n, &m)
		if ((n * m) & 1) != 0 {
			fmt.Fprintln(out, "First")
			out.Flush()
			fmt.Fprintf(out, "%d %d\n", 0, 0)
			out.Flush()
			for {
				var x, y int
				fmt.Scan(&x, &y)
				if x < 0 {
					break
				}
				if (x&1) != 0 && y == m-1 {
					fmt.Fprintf(out, "%d %d\n", x+1, 0)
					out.Flush()
				} else if (x&1) == 0 && y == 0 {
					fmt.Fprintf(out, "%d %d\n", x-1, m-1)
					out.Flush()
				} else {
					if ((x + y) & 1) != 0 {
						fmt.Fprintf(out, "%d %d\n", x, y+1)
						out.Flush()
					} else {
						fmt.Fprintf(out, "%d %d\n", x, y-1)
						out.Flush()
					}
				}
			}
		} else {
			fmt.Fprintln(out, "Second")
			out.Flush()
			for {
				var x, y int
				fmt.Scan(&x, &y)
				if x < 0 {
					break
				}
				fmt.Fprintf(out, "%d %d\n", n-1-x, m-1-y)
				out.Flush()
			}
		}
	}
}
