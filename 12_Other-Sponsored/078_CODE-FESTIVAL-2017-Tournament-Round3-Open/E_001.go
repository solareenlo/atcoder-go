package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	ts := make([]tuple, n-1)
	for i := 0; i < n-1; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		ts[i] = tuple{c, a - 1, b - 1}
	}
	sortTuple(ts)

	p := make([]int, n)
	ch := make([][]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
		ch[i] = append(ch[i], p[i])
	}

	laz := make([]int, n)
	ans := make([]int, n)
	for i := 0; i < n-1; i++ {
		c, a, b := ts[i].x, ts[i].y, ts[i].z
		a = p[a]
		b = p[b]
		if len(ch[a]) < len(ch[b]) {
			a, b = b, a
		}
		sd := len(ch[a]) - len(ch[b])
		for _, d := range ch[b] {
			p[d] = a
			ans[d] += laz[b] + sd*c - laz[a]
			ch[a] = append(ch[a], d)
		}
		laz[a] += len(ch[b]) * c
	}
	for i := 0; i < n; i++ {
		fmt.Fprintln(out, ans[i]+laz[p[0]])
	}
}

type tuple struct {
	x, y, z int
}

func sortTuple(tup []tuple) {
	sort.Slice(tup, func(i, j int) bool {
		if tup[i].x == tup[j].x {
			if tup[i].y == tup[j].y {
				return tup[i].z > tup[j].z
			}
			return tup[i].y > tup[j].y
		}
		return tup[i].x > tup[j].x
	})
}
