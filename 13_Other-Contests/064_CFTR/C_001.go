package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Scan(&n, &k)

	w := make([]int64, 1000005)
	d := make([]int64, 1000005)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &w[i], &d[i])
	}

	l, r := int64(0), int64(2e18)
	for r-l > 1 {
		m := (l + r) / 2
		cnt := 0
		for i := 0; i < n; i++ {
			if m >= w[i] {
				cnt += int((m-w[i])/d[i]) + 1
			}
		}
		if cnt < k {
			l = m
		} else {
			r = m
		}
	}

	fmt.Println(r)
}
