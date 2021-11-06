package main

import (
	"fmt"
	"sort"
)

type pair struct{ x, y int }

func check(s []pair) bool {
	h := 0
	for _, p := range s {
		b := h + p.x
		if b < 0 {
			return false
		}
		h += p.y
	}
	return true
}

func main() {
	var n int
	fmt.Scan(&n)

	up := make([]pair, n)
	down := make([]pair, n)
	total := 0
	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		h, b := 0, 0
		for i := range s {
			if s[i] == '(' {
				h++
			} else {
				h--
			}
			b = min(b, h)
		}
		if h > 0 {
			up = append(up, pair{b, h})
		} else {
			down = append(down, pair{b - h, -h})
		}
		total += h
	}

	sort.Slice(up, func(i, j int) bool {
		return up[i].x > up[j].x
	})
	sort.Slice(down, func(i, j int) bool {
		return down[i].x > down[j].x
	})

	if check(up) && check(down) && total == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
