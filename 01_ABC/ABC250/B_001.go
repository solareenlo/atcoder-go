package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, a, b int
	fmt.Scan(&n, &a, &b)

	for i := 1; i <= n*a; i++ {
		for j := 1; j <= n*b; j++ {
			if ((i+a-1)/a+(j+b-1)/b)&1 != 0 {
				fmt.Fprint(out, "#")
			} else {
				fmt.Fprint(out, ".")
			}
		}
		fmt.Fprintln(out)
	}
}
