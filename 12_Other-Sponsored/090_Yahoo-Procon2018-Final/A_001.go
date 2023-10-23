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

	var n, m int
	fmt.Fscan(in, &n, &m)
	var r [100005]int
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		r[a]++
	}
	var c [100005]int
	for i := 2; i <= m; i++ {
		for j := i; j <= 100000; j += i {
			c[i] += r[j]
		}
	}
	for i := 1; i <= m; i++ {
		a := i
		s := n
		p := make([]int, 0)
		for j := 2; j*j <= a; j++ {
			if a%j == 0 {
				p = append(p, j)
				for a%j == 0 {
					a /= j
				}
			}
		}
		if a != 1 {
			p = append(p, a)
		}
		o := len(p)
		for k := 1; k < 1<<o; k++ {
			q := 1
			w := 0
			for j := 0; j < o; j++ {
				if (k >> j & 1) != 0 {
					q *= p[j]
					w++
				}
			}
			if (w % 2) != 0 {
				s -= c[q]
			} else {
				s += c[q]
			}
		}
		fmt.Fprintln(out, s)
	}
}
