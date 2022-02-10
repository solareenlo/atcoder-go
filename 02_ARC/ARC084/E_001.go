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

	var k, n int
	fmt.Fscan(in, &k, &n)

	a := make([]int, n+1)
	if k&1 != 0 {
		for i := 1; i <= n; i++ {
			a[i] = (k + 1) >> 1
		}
		t := n
		x := n / 2
		for j := 0; j < x; j++ {
			if a[t] == 1 {
				a[t] = k
				t--
			} else {
				a[t]--
				t = n
			}
		}
		for i := 1; i <= t; i++ {
			fmt.Fprint(out, a[i], " ")
		}
	} else {
		fmt.Fprint(out, k/2, " ")
		for i := 1; i < n; i++ {
			fmt.Fprint(out, k, " ")
		}
	}
	fmt.Fprintln(out)
}
