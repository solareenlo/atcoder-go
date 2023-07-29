package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	type pair struct {
		x, y int
	}

	var n, m int
	fmt.Fscan(in, &n, &m)
	edge := make([]pair, 2000)
	var g [2000][2000]bool
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		edge[i] = pair{a, b}
		g[a][b] = true
		g[b][a] = true
	}
	y := make([]int, n)
	used := make([]bool, n)
	sum := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &y[i])
		y[i]--
	}
	d := NewDsu(n)
	for i := 0; i < m; i++ {
		a := edge[i].x
		b := edge[i].y
		for j := 0; j < n; j++ {
			if g[a][j] && g[b][j] {
				d.Merge(a, b)
				d.Merge(a, j)
			}
		}
	}
	for i := 0; i < n; i++ {
		if !used[i] {
			cur := i
			cnt := 0
			for !used[cur] {
				cnt++
				used[cur] = true
				if !d.Same(cur, y[cur]) {
					fmt.Println("NO")
					return
				}
				cur = y[cur]
			}
			sum[d.Leader(i)] += cnt - 1
		}
	}
	for _, x := range sum {
		if x%2 != 0 {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}

type dsu struct {
	n            int
	parentOrSize []int
}

func NewDsu(n int) *dsu {
	d := new(dsu)
	d.n = n
	d.parentOrSize = make([]int, d.n)
	for i := range d.parentOrSize {
		d.parentOrSize[i] = -1
	}
	return d
}

func (d *dsu) Merge(a, b int) int {
	if !(0 <= a && a < d.n) {
		panic("")
	}
	if !(0 <= b && b < d.n) {
		panic("")
	}
	x := d.Leader(a)
	y := d.Leader(b)
	if x == y {
		return x
	}
	if -d.parentOrSize[x] < -d.parentOrSize[y] {
		x, y = y, x
	}
	d.parentOrSize[x] += d.parentOrSize[y]
	d.parentOrSize[y] = x
	return x
}

func (d *dsu) Same(a, b int) bool {
	if !(0 <= a && a < d.n) {
		panic("")
	}
	if !(0 <= b && b < d.n) {
		panic("")
	}
	return d.Leader(a) == d.Leader(b)
}

func (d *dsu) Leader(a int) int {
	if !(0 <= a && a < d.n) {
		panic("")
	}
	if d.parentOrSize[a] < 0 {
		return a
	}
	d.parentOrSize[a] = d.Leader(d.parentOrSize[a])
	return d.parentOrSize[a]
}

func (d *dsu) Size(a int) int {
	if !(0 <= a && a < d.n) {
		panic("")
	}
	return -d.parentOrSize[d.Leader(a)]
}

func (d *dsu) Groups() [][]int {
	leaderBuf := make([]int, d.n)
	groupSize := make([]int, d.n)
	for i := 0; i < d.n; i++ {
		leaderBuf[i] = d.Leader(i)
		groupSize[leaderBuf[i]]++
	}
	result := make([][]int, d.n)
	for i := 0; i < d.n; i++ {
		result[i] = make([]int, 0, groupSize[i])
	}
	for i := 0; i < d.n; i++ {
		result[leaderBuf[i]] = append(result[leaderBuf[i]], i)
	}
	eraseEmpty := func(a [][]int) [][]int {
		result := make([][]int, 0, len(a))
		for i := range a {
			if len(a[i]) != 0 {
				result = append(result, a[i])
			}
		}
		return result
	}
	return eraseEmpty(result)
}
