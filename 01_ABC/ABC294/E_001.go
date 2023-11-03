package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var a, b [200200]int

	var L, n, m int
	fmt.Fscan(in, &L, &n, &m)
	res := 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i], &b[i])
	}
	tp := 1
	for m > 0 {
		m--
		var v, l int
		fmt.Fscan(in, &v, &l)
		for l > 0 {
			Len := min(l, b[tp])
			if v == a[tp] {
				res += Len
			}
			b[tp] -= Len
			l -= Len
			if b[tp] == 0 {
				tp++
			}
		}
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
