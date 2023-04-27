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

	var k int
	fmt.Fscan(in, &k)

	n := make([]int, k)
	m := make([]int, k)
	good := make([]bool, k)
	val := make([]int, k)
	for id := 0; id < k; id++ {
		fmt.Fscan(in, &n[id], &m[id])
		g := make([][]int, n[id])
		d := NewDsu(n[id])
		for i := 0; i < m[id]; i++ {
			var x, y int
			fmt.Fscan(in, &x, &y)
			x--
			y--
			g[x] = append(g[x], y)
			g[y] = append(g[y], x)
			d.merge(x, y)
		}
		good[id] = true
		for i := 0; i < n[id]; i++ {
			if len(g[i]) != 2 {
				good[id] = false
				break
			}
		}
		if !good[id] {
			continue
		}
		sizes := make([]int, n[id])
		for i := 0; i < n[id]; i++ {
			sizes[d.find(i)] += 1
		}
		sort.Ints(sizes)
		sizes = reverseOrderInt(sizes)
		for sizes[len(sizes)-1] == 0 {
			sizes = sizes[:len(sizes)-1]
		}
		for _, x := range sizes {
			if x%2 == 0 {
				good[id] = false
			}
		}
		if sizes[0] != sizes[len(sizes)-1] {
			good[id] = false
		}
		val[id] = sizes[0]
	}
	for i := 0; i < k; i++ {
		for j := 0; j < k; j++ {
			if m[i]+m[j] == 0 || (good[i] && good[j] && val[i] == val[j]) {
				fmt.Fprint(out, 1)
			} else {
				fmt.Fprint(out, 0)
			}
		}
		fmt.Fprintln(out)
	}
}

type dsu struct {
	n int
	p []int
}

func NewDsu(n int) *dsu {
	d := new(dsu)
	d.n = n
	d.p = make([]int, n)
	for i := range d.p {
		d.p[i] = i
	}
	return d
}

func (d *dsu) find(x int) int {
	if x == d.p[x] {
		return x
	}
	d.p[x] = d.find(d.p[x])
	return d.p[x]
}

func (d *dsu) merge(x, y int) bool {
	x = d.find(x)
	y = d.find(y)
	if x != y {
		d.p[x] = y
		return true
	}
	return false
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
