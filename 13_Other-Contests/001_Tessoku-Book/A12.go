package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	l := 0
	r := 1000000001
	for l < r {
		m := (l + r) / 2
		preans := 0
		for i := 0; i < n; i++ {
			preans += m / a[i]
		}
		if k <= preans {
			r = m
		} else {
			l = m + 1
		}
	}
	fmt.Println(l)
}
