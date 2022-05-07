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
	b := make([]int, n)
	c := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i], &b[i], &c[i])
	}

	l := 0
	r := 1 << 60
	for r-l > 1 {
		mid := (l + r) / 2
		cnt := 0
		for i := 0; i < n; i++ {
			d := mid - b[i]
			if d < 0 {
				continue
			}
			d = d/c[i] + 1
			if d > a[i] {
				d = a[i]
			}
			cnt += d
		}
		if cnt >= k {
			r = mid
		} else {
			l = mid
		}
	}
	fmt.Println(r)
}
