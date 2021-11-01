package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	s := make([]int, m)
	c := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&s[i], &c[i])
	}

	for i := 0; i < 1000; i++ {
		x := i / 10
		d := make([]int, 0)
		d = append(d, i%10)
		for x > 0 {
			d = append(d, x%10)
			x /= 10
		}
		d = reverseOrderInt(d)
		if len(d) != n {
			continue
		}

		ok := true
		for j := 0; j < m; j++ {
			if d[s[j]-1] != c[j] {
				ok = false
				break
			}
		}
		if ok {
			fmt.Println(i)
			return
		}
	}

	fmt.Println(-1)
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
