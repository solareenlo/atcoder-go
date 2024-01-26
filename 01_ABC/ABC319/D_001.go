package main

import (
	"bufio"
	"fmt"
	"os"
)

var n, m int
var a [500005]int

func p(z int) bool {
	x, y := 0, a[0]
	for i := 1; i < n; i++ {
		if y+1+a[i] > z {
			x++
			y = a[i]
		} else {
			y += a[i] + 1
		}
	}
	return x < m
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Scan(&n, &m)
	var l, r int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		l = max(l, a[i])
		r += a[i] + 1
	}
	for l < r {
		mid := (l + r) >> 1
		if p(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	fmt.Println(l)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
