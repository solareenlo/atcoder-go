package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const sz = 1 << 18

	var d [1 << 19]int

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 0; i < m; i++ {
		var t, l, r int
		fmt.Fscan(in, &t, &l, &r)
		for l, r = l+sz, r+sz; l <= r; l, r = l/2, r/2 {
			d[l] = t
			d[r] = t
			l++
			r--
		}
	}
	a := 0
	for i := 2; i < sz*2; i++ {
		if d[i/2] > d[i] {
			d[i] = d[i/2]
		}
		if i >= sz {
			a += d[i]
		}
	}
	fmt.Println(a)
}
