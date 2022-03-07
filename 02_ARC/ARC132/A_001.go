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

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	b := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &b[i])
	}

	var q int
	fmt.Fscan(in, &q)
	for i := 0; i < q; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		if a[x]+b[y] > n {
			fmt.Fprint(out, "#")
		} else {
			fmt.Fprint(out, ".")
		}
	}
}
