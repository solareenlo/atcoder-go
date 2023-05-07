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

	var f [17]int
	f[0] = 1
	f[1] = 1
	for i := 1; i < 16; i++ {
		f[i+1] = f[i] + f[i-1]
	}
	var t int
	fmt.Fscan(in, &t)
	for t > 0 {
		t--
		var n int
		fmt.Fscan(in, &n)
		i := 0
		for f[i+1] <= n {
			i++
		}
		m := 0
		M := 0
		e := 1
		for i > 0 {
			i--
			a := m + f[i]*e
			var A int
			if a <= n {
				fmt.Fprintln(out, "?", a)
				out.Flush()
				fmt.Fscan(in, &A)
			}
			if A > M {
				m = a
				M = A
			} else {
				e = -e
			}
		}
		fmt.Fprintln(out, "!", M)
		out.Flush()
	}
}
