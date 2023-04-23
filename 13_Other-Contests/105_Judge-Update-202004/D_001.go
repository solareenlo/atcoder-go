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
	fmt.Fscan(in, &n, &q)
	var a [100005]int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		if i != 0 {
			a[i] = gcd(a[i], a[i-1])
		}
	}
	for i := 0; i < q; i++ {
		var s int
		fmt.Fscan(in, &s)
		l := -1
		r := n
		for l+1 < r {
			m := (l + r) / 2
			if gcd(s, a[m]) == 1 {
				r = m
			} else {
				l = m
			}
		}
		if r == n {
			fmt.Fprintln(out, gcd(s, a[n-1]))
		} else {
			fmt.Fprintln(out, r+1)
		}
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
