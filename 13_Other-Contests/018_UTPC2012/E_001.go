package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	s := 0
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i], &b[i])
		s += a[i]
	}
	type pair struct {
		x, y int
	}
	ans := int(1e18)
	for i := 0; i < s; i++ {
		mod := make([]pair, n)
		c := make([]int, n)
		left := i
		for j := 0; j < n; j++ {
			mod[j] = pair{-(a[j] * i) % s, j}
			left -= (a[j] * i) / s
			c[j] = (a[j] * i) / s
		}
		sort.Slice(mod, func(i, j int) bool {
			if mod[i].x == mod[j].x {
				return mod[i].y < mod[j].y
			}
			return mod[i].x < mod[j].x
		})
		for j := 0; j < left; j++ {
			c[mod[j].y]++
		}
		dif := -1
		for j := 0; j < n; j++ {
			dif = max(dif, (b[j]-c[j]+a[j]-1)/a[j])
		}
		ans = min(ans, dif*s+i)
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
