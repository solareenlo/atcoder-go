package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	l, r := 1, n
	for l < r {
		mid := (l + r) >> 1
		fmt.Printf("? 1 %d %d %d\n", n, l, mid)
		var x int
		fmt.Fscan(in, &x)
		if x == mid-l+1 {
			l = mid + 1
		} else {
			r = mid
		}
	}
	y := l
	l, r = 1, n
	for l < r {
		mid := (l + r) >> 1
		fmt.Printf("? %d %d 1 %d\n", l, mid, n)
		var x int
		fmt.Fscan(in, &x)
		if x == mid-l+1 {
			l = mid + 1
		} else {
			r = mid
		}
	}
	fmt.Printf("! %d %d\n", l, y)
}
