package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 300005

type node struct{ x, y int }

var (
	w, h, n int
	ans     int
	Y       int
	p       = make([]node, N)
	a       = make([]int, N)
	b       = make([]int, N)
	q       = make([]int, N)
)

func solve(l, r int) {
	if l == r {
		return
	}
	mid := (l + r) >> 1
	for i, y1, y2 := mid, 0, h; i >= l; i-- {
		a[i] = y1
		b[i] = y2
		if p[i].y >= Y {
			y2 = min(y2, p[i].y)
		}
		if p[i].y <= Y {
			y1 = max(y1, p[i].y)
		}
	}
	for i, y1, y2 := mid+1, 0, h; i <= r; i++ {
		a[i] = y1
		b[i] = y2
		if p[i].y >= Y {
			y2 = min(y2, p[i].y)
		}
		if p[i].y <= Y {
			y1 = max(y1, p[i].y)
		}
	}
	for i, j, h, t := mid, mid+1, 1, 0; i >= l; i-- {
		for ; j <= r && a[j] <= a[i]; j++ {
			for ; h <= t && b[q[t]]+p[q[t]].x <= b[j]+p[j].x; t-- {

			}
			t++
			q[t] = j
		}
		for ; h < t && b[i]+p[q[h]].x <= b[q[h+1]]+p[q[h+1]].x; h++ {

		}
		ans = max(ans, (p[q[h]].x-p[i].x+min(b[i], b[q[h]])-a[i])*2)
	}
	for i, j, h, t := mid+1, mid, 1, 0; i <= r; i++ {
		for ; j >= l && a[j] <= a[i]; j-- {
			for ; h <= t && b[q[t]]-p[q[t]].x <= b[j]-p[j].x; t-- {

			}
			t++
			q[t] = j
		}
		for ; h < t && b[i]-p[q[h]].x <= b[q[h+1]]-p[q[h+1]].x; h++ {

		}
		ans = max(ans, (p[i].x-p[q[h]].x+min(b[i], b[q[h]])-a[i])*2)
	}
	solve(l, mid)
	solve(mid+1, r)
}

func work() {
	sort.Slice(p[:n+1], func(i, j int) bool {
		return p[i].x < p[j].x
	})
	Y = h / 2
	p[0].x = 0
	p[n+1].x = w
	p[n+1].y = Y
	p[0].y = Y
	solve(0, n+1)
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &w, &h, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &p[i].x, &p[i].y)
	}
	ans = max(h, w)*2 + 2
	work()
	w, h = h, w
	for i := 1; i <= n; i++ {
		p[i].x, p[i].y = p[i].y, p[i].x
	}
	work()
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
