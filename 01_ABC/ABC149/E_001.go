package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	sort.Ints(a)

	s := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		s[i] = s[i+1] + a[i]
	}

	l, r := 0, 1<<60
	for r-l > 1 {
		mid := (l + r) / 2
		cnt := 0
		for i := 0; i < n; i++ {
			pos := lowerBound(a, mid-a[i])
			cnt += n - pos
		}
		if cnt >= m {
			l = mid
		} else {
			r = mid
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		pos := lowerBound(a, r-a[i])
		if pos < n {
			res += a[i]*(n-pos) + s[pos]
			m -= n - pos
		}
	}
	fmt.Println(res + l*m)
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
