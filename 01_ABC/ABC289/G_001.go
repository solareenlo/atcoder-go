package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 200200

var n, m int
var a []pair
var b, ans []int

func solve(l, r, tl, tr int) {
	if tl > tr {
		return
	}
	mid := (tl + tr) / 2
	mx := 0
	p := 0
	for i := l; i <= r; i++ {
		t := (n - i + 1) * (b[i] + a[mid].x)
		if t > mx {
			mx = t
			p = i
		}
	}
	ans[a[mid].y] = mx
	solve(l, p, mid+1, tr)
	solve(p, r, tl, mid-1)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	ans = make([]int, N)
	a = make([]pair, N)

	fmt.Fscan(in, &n, &m)
	b = make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &b[i])
	}
	sort.Ints(b[1 : n+1])

	c := make([]int, m+1)
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &c[i])
		a[i] = pair{c[i], i}
	}
	sortPair(a[1 : m+1])

	solve(1, n, 1, m)
	for i := 1; i <= m; i++ {
		fmt.Fprintf(out, "%d ", ans[i])
	}
	fmt.Fprintln(out)
}

type pair struct {
	x, y int
}

func sortPair(tmp []pair) {
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].x == tmp[j].x {
			return tmp[i].y < tmp[j].y
		}
		return tmp[i].x < tmp[j].x
	})
}
