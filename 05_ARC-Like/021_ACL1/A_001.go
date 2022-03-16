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

	const N = 300000
	a := make([]int, N)
	b := make([]int, N)
	c := make([]int, N)
	var p, q int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &p, &q)
		p--
		q--
		b[p] = q
		c[p] = i
	}

	p = 0
	q = 0
	for x := 0; x < n; x++ {
		r := n - 1 - b[x]
		if r > p {
			p = r
		}
		if p == x {
			for i := q; i < x+1; i++ {
				a[c[i]] = x + 1 - q
			}
			q = x + 1
		}
	}

	for i := 0; i < n; i++ {
		fmt.Fprintln(out, a[i])
	}
}
