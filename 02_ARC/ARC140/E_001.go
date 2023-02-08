package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Scan(&n, &m)

	p := 23
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Fprintf(out, "%d ", ((i/p)*(j/p)+i+j)%p+1)
		}
		fmt.Fprintln(out)
	}
}
