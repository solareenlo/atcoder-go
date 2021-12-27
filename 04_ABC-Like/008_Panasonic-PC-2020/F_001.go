package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	M := 205891132094649
	var C func(a, b, c, d int) int
	C = func(a, b, c, d int) int {
		u, l, p := b, b, M
		for p > 0 {
			if (a/p+(4-a/p%3)%3)*p <= c {
				if u/p%3 == 1 {
					u = u/(3*p)*3*p + 2*p
				}
				if l/p%3 == 1 {
					l = l/(3*p)*3*p + p - 1
				}
			}
			p /= 3
		}
		return c - a + min(u-b+max(u, d)-min(u, d), b+d-2*l)
	}

	in := bufio.NewReader(os.Stdin)

	var Q int
	fmt.Fscan(in, &Q)

	for i := 0; i < Q; i++ {
		var a, b, c, d int
		fmt.Fscan(in, &a, &b, &c, &d)
		a--
		c--
		if a > c {
			a, c = c, a
			b, d = d, b
		}
		b--
		d--
		if b > d {
			b = M - b - 1
			d = M - d - 1
		}
		fmt.Println(max(C(a, b, c, d), C(b, a, d, c)))
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
