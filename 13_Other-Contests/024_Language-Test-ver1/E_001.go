package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var a, b int
	fmt.Scanf("%d/%d\n", &a, &b)
	flag := true
	n := a * 2 / b
	for i := 0; i < 2; i++ {
		if ((b*(n+1)-2*a)*n)%(2*b) == 0 {
			m := ((b*(n+1) - 2*a) * n) / (2 * b)
			if m >= 1 && m <= n {
				fmt.Fprintf(out, "%d %d\n", n, m)
				flag = false
			}
		}
		n++
	}
	if flag {
		fmt.Fprintln(out, "Impossible")
	}
}
