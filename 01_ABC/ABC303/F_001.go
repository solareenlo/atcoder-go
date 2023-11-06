package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func fx(l, r int) int {
	return (l + r) * (r - l + 1) / 2
}

func main() {
	in := bufio.NewReader(os.Stdin)

	type date struct {
		t, d int
	}

	var n, H int
	fmt.Fscan(in, &n, &H)
	a := make([]date, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i].t, &a[i].d)
	}
	tmp := a[1:]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].t < tmp[j].t
	})
	m := 0
	for i := 1; i <= n; i++ {
		for m != 0 && a[m].d <= a[i].d {
			m--
		}
		m++
		a[m] = a[i]
	}
	n = m
	cur := 0
	mx := 0
	for i := 1; i <= n; i++ {
		x := min(mx/a[i].d, a[i].t)
		y := x - cur
		if mx != 0 {
			y = min(y, (H+mx-1)/mx)
		}
		H -= mx * y
		cur += y
		if H <= 0 {
			fmt.Println(cur)
			return
		}
		if cur != a[i].t {
			L := cur
			R := a[i].t
			r := (H + a[i].d - 1) / a[i].d
			for L+1 < R {
				mid := (L + R) / 2
				if fx(cur+1, mid) >= r {
					R = mid
				} else {
					L = mid
				}
			}
			H -= fx(cur+1, R) * a[i].d
			cur = R
			if H <= 0 {
				fmt.Println(cur)
				return
			}
		}
		mx = max(mx, a[i].t*a[i].d)
	}
	fmt.Println(cur + (H+mx-1)/mx)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
