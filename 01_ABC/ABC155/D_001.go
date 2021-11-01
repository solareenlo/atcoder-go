package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)

	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a)

	start, end := -1<<60, 1<<60
	for end-start > 1 {
		mid := (start + end) / 2
		cnt := 0
		for i := 0; i < n; i++ {
			l, r := i, n
			for r-l > 1 {
				m := (l + r) / 2
				if a[i]*a[m] <= mid {
					if a[i] < 0 {
						r = m
					} else {
						l = m
					}
				} else {
					if a[i] < 0 {
						l = m
					} else {
						r = m
					}
				}
			}
			if a[i] < 0 {
				cnt += n - r
			} else {
				cnt += l - i
			}
		}
		if cnt < k {
			start = mid
		} else {
			end = mid
		}
	}

	fmt.Println(end)
}
