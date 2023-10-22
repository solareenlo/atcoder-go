package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type node struct {
	mi, tag int
}

var a [1000000]node
var y []int

func pushdown(now int) {
	a[now*2].mi += a[now].tag
	a[now*2+1].mi += a[now].tag
	a[now*2].tag += a[now].tag
	a[now*2+1].tag += a[now].tag
	a[now].tag = 0
}

func pushup(now int) {
	a[now].mi = min(a[now*2].mi, a[now*2+1].mi)
}

func build(now, l, r int) {
	if l == r {
		a[now].mi = y[l]
		return
	}
	mid := (l + r) / 2
	build(now*2, l, mid)
	build(now*2+1, mid+1, r)
	pushup(now)
}

func query(now, l, r, L, R int) int {
	if l == L && r == R {
		return a[now].mi
	}
	pushdown(now)
	mid := (l + r) / 2
	if R <= mid {
		return query(now*2, l, mid, L, R)
	} else if L > mid {
		return query(now*2+1, mid+1, r, L, R)
	} else {
		return min(query(now*2, l, mid, L, mid), query(now*2+1, mid+1, r, mid+1, R))
	}
}

func add(now, l, r, L, R, v int) {
	if l == L && r == R {
		a[now].mi += v
		a[now].tag += v
		return
	}
	pushdown(now)
	mid := (l + r) / 2
	if R <= mid {
		add(now*2, l, mid, L, R, v)
	} else if L > mid {
		add(now*2+1, mid+1, r, L, R, v)
	} else {
		add(now*2, l, mid, L, mid, v)
		add(now*2+1, mid+1, r, mid+1, R, v)
	}
	pushup(now)
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	x := make([]int, n+1)
	y = make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &x[i])
		y[i] = x[i]
	}
	sort.Ints(y[1:])
	build(1, 1, n)
	ans := 0
	for i := 1; i <= n; i++ {
		pos := lowerBound(y[1:], x[i]) + 1
		ans += query(1, 1, n, pos, n) - x[i]
		if pos-1 != 0 {
			add(1, 1, n, 1, pos-1, 1)
		}
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
