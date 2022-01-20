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
	a = reverseOrderInt(a)

	b := make([]int, n)
	for i := range b {
		fmt.Fscan(in, &b[i])
	}
	sort.Ints(b)

	l := 0
	r := 1 << 60
	for r-l > 1 {
		m := (l + r) / 2
		c := 0
		R := 0
		for i := 0; i < n; i++ {
			for R < n && a[i]*b[R] < m {
				R++
			}
			c += R
		}
		if c < k {
			l = m
		} else {
			r = m
		}
	}
	fmt.Println(l)
}

func reverseOrderInt(a []int) []int {
	n := len(a)
	res := make([]int, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
