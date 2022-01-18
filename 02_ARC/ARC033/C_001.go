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

	var q int
	fmt.Fscan(in, &q)

	n := 1 << 18
	s := make([]int, 1<<19)
	for i := 0; i < q; i++ {
		var t, x int
		fmt.Fscan(in, &t, &x)
		t &= 1
		if t == 0 {
			k := 1
			l := 0
			for k < n {
				k *= 2
				m := l + s[k]
				if m < x {
					l = m
					k++
				}
			}
			x = k - n
			fmt.Fprintln(out, x)
		}
		x += n
		s[x] = t
		x /= 2
		for x > 0 {
			s[x] = s[x*2] + s[x*2+1]
			x /= 2
		}
	}
}
