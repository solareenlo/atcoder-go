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
	var q int
	fmt.Fscan(in, &q)
	for Q := 1; Q <= q; Q++ {
		var t, k int
		fmt.Fscan(in, &t, &k)
		if t == 1 {
			fmt.Fscan(in, &a[k])
		} else {
			fmt.Fprintln(out, a[k])
		}
	}
}
