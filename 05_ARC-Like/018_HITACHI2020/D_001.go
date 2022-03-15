package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, t int
	fmt.Fscan(in, &n, &t)

	t++
	f := make([]int, 32)
	f[0] = 1
	for i := 1; i < 32; i++ {
		f[i] = t + 1
	}

	type pair struct{ x, y int }
	a := make([]pair, n+1)
	c := make([]int, n+2)
	c1, c2 := 0, 0
	for i := 1; i < n+1; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		y++
		if x > 0 {
			c1++
			a[c1] = pair{x, y}
		} else {
			c2++
			c[c2] = y
		}
	}

	a = a[1 : c1+1]
	sort.Slice(a, func(i, j int) bool {
		return a[i].x*a[j].y > a[j].x*a[i].y
	})
	tmp := c[1 : c2+1]
	sort.Ints(tmp)

	for i := 1; i < c1+1; i++ {
		for j := 31; j >= 1; j-- {
			if f[j-1] < t {
				f[j] = min(f[j-1]*(a[i-1].x+1)+a[i-1].y, f[j])
			}
		}
	}

	ans := 0
	for i := 31; i >= 0; i-- {
		if f[i] <= t {
			nw := 1
			for f[i]+c[nw] <= t && nw <= c2 {
				f[i] += c[nw]
				nw++
			}
			ans = max(ans, i+nw-1)
		}
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
