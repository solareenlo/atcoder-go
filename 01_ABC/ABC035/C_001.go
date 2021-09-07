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
	fmt.Scan(&n, &q)
	o := make([]int, n+2)
	var l, r int
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &l, &r)
		o[l-1]++
		o[r]++
	}
	for i := 0; i < n; i++ {
		o[i+1] += o[i]
	}
	for i := 0; i < n; i++ {
		fmt.Fprint(out, o[i]%2)
	}
	fmt.Fprintln(out)
}
