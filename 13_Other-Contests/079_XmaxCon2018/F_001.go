package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Scan(&n)
	r := 5000000
	for i := 0; i < n; i++ {
		a := (i + 1) * r
		b := (n - i) * r
		fmt.Fprintf(out, "%d %d %d %d %d %d %d %d %d %d\n",
			0, a,
			b, 0,
			0, a+1,
			b, 1,
			0, a+2)
	}
}
