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

	var a [2][100010]int
	for i := 1; i <= n; i++ {
		var c, v int
		fmt.Fscan(in, &c, &v)
		a[1][i] = a[1][i-1]
		a[0][i] = a[0][i-1]
		a[c-1][i] += v
	}

	var q int
	fmt.Fscan(in, &q)
	for q > 0 {
		q--
		var l, r int
		fmt.Fscan(in, &l, &r)
		fmt.Fprintln(out, a[0][r]-a[0][l-1], a[1][r]-a[1][l-1])
	}
}
