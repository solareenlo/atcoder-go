package main

import (
	"bufio"
	"fmt"
	"os"
)

func sum(a, b int) int {
	n := a / b
	return n*(n-1)/2*b + (a-b*n)*n
}

func calc(l, r, a, b int) int {
	return a*(r-l) + sum(r, b) - sum(l, b)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var n, a, b, c, d int
		fmt.Fscan(in, &n, &a, &b, &c, &d)

		if b == d {
			if a == c {
				fmt.Fprintln(out, n)
				continue
			}
			fmt.Fprintln(out, 0)
			continue
		}
		if b > d {
			a, c = c, a
			b, d = d, b
		}
		l := max(0, b*d*(c-a-1)/(d-b)) + 1
		r := min(n, b*d*(c-a)/(d-b)) + 1
		ans := 0
		if l < r {
			ans = r - l - calc(l, r, c, d) + calc(l, r, a, b)
		}
		l = max(0, b*d*(c-a)/(d-b)) + 1
		r = min(n, b*d*(c-a+1)/(d-b)) + 1
		if l < r {
			ans += r - l - calc(l, r, a, b) + calc(l, r, c, d)
		}
		fmt.Fprintln(out, ans)
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
