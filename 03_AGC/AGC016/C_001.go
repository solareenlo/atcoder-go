package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, a, b int
	fmt.Scan(&n, &m, &a, &b)

	if n%a == 0 && m%b == 0 {
		fmt.Fprintln(out, "No")
		return
	}

	fmt.Fprintln(out, "Yes")
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if i%a == 0 && j%b == 0 {
				fmt.Fprint(out, -a*b*500+499, " ")
			} else {
				fmt.Fprint(out, 500, " ")
			}
		}
		fmt.Fprintln(out)
	}
}
