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

	var n, q int
	fmt.Fscan(in, &n, &q)

	a := make([]int, n+1)
	b := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = i
		b[i] = i
	}

	for i := 0; i < q; i++ {
		var x int
		fmt.Fscan(in, &x)
		o := b[x]
		if o == n {
			o = n - 1
		}
		a[o], a[o+1] = a[o+1], a[o]
		b[a[o]], b[a[o+1]] = b[a[o+1]], b[a[o]]
	}

	for i := 1; i <= n; i++ {
		fmt.Fprint(out, a[i], " ")
	}
}
