package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, x, y int
	fmt.Fscan(in, &n, &x, &y)

	var a [300001]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	l := 0
	r := int(1e10)
	s := 0
	for l <= r {
		m := (l + r) / 2
		b := 0
		for i := 1; i <= n; i++ {
			if a[i] >= m {
				b += (a[i] - m) / y
			} else {
				b -= (m - a[i] + x - 1) / x
			}
		}
		if b >= 0 {
			l = m + 1
			s = m
		} else {
			r = m - 1
		}
	}
	fmt.Println(s)
}
