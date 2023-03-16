package main

import (
	"fmt"
)

var (
	n   int
	l   int
	s   [5010]string
	x   [5010]int
	y   [5010]int
	k   [5010]int
	m   map[string]int
	ans string
	po  [20010]int
)

func det(x1 int, y1 int, x2 int, y2 int) int {
	d := int64(x1)*int64(y2) - int64(x2)*int64(y1)
	if d == 0 {
		return 0
	}
	if d > 0 {
		return 1
	}
	return -1
}

func main() {
	m = make(map[string]int)
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i], &x[i], &y[i])
		m[s[i]] = i
		po[x[i]+10000] = i + 1
	}
	fmt.Scan(&l)
	sx := 0
	sy := 0
	for i := 0; i < l; i++ {
		var t string
		fmt.Scan(&t)
		k[i] = m[t]
		sx += x[k[i]]
		sy += y[k[i]]
	}
	dir := det(x[k[1]]*l-x[k[0]]*l, y[k[1]]*l-y[k[0]]*l, sx-x[k[0]]*l, sy-y[k[0]]*l)
	k[l] = k[0]
	for i := 0; i < 20010; i++ {
		if po[i] == 0 {
			continue
		}
		t := po[i] - 1
		f := 1
		for j := 0; j < l; j++ {
			tmp := det(x[k[j+1]]-x[k[j]], y[k[j+1]]-y[k[j]], x[t]-x[k[j]], y[t]-y[k[j]])
			if tmp == dir {
				f &= 1
			} else {
				f &= 0
			}
		}
		if f == 1 {
			fmt.Println(s[t])
		}
	}
}
