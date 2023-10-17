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
	var a [55]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	var x int
	fmt.Fscan(in, &x)
	l := 1
	r := n + 1
	for l < r {
		mid := (l + r) >> 1
		if a[mid] < x {
			l = mid + 1
		} else {
			r = mid
		}
	}
	fmt.Println(l)
}
