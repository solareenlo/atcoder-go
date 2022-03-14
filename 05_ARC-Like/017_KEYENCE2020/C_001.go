package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k, s int
	fmt.Scan(&n, &k, &s)

	for i := 0; i < n; i++ {
		if i < k {
			fmt.Fprint(out, s, " ")
		} else if s == 1 {
			fmt.Fprint(out, s+1, " ")
		} else {
			fmt.Fprint(out, s-1, " ")
		}
	}
}
