package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var a, d [100100]int

	var n, k int
	fmt.Fscan(in, &n, &k)
	s, l, r := 0, 0, 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < n; i++ {
		for l < r && a[d[r-1]] >= a[i] {
			r--
		}
		d[r] = i
		r++
		s += a[d[l]]
		if d[l] == i-k+1 {
			l++
		}
	}
	fmt.Println(s)
}
