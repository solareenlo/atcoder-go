package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)

	E := make([][]int, m)

	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		a--
		b--
		E[i] = []int{c, a, b}
	}
	loop := make([][]int, 0)
	unloop := make([][]int, 0)
	for i := 0; i < m; i++ {
		D := NewDsu(n)
		for j := 0; j < m; j++ {
			D.Merge(E[j][1], E[j][2])
		}
		D2 := NewDsu(n)
		for j := 0; j < m; j++ {
			if j == i {
				continue
			}
			D2.Merge(E[j][1], E[j][2])
		}
		if len(D.Groups()) != len(D2.Groups()) {
			unloop = append(unloop, E[i])
		} else {
			loop = append(loop, E[i])
		}
	}

	D := NewDsu(n)
	for i := 0; i < m; i++ {
		D.Merge(E[i][1], E[i][2])
	}
	now := len(D.Groups())

	sort.Slice(loop, func(i, j int) bool {
		for k := range loop[i] {
			return loop[i][k] < loop[j][k]
		}
		return false
	})
	sort.Slice(unloop, func(i, j int) bool {
		for k := range unloop[i] {
			return unloop[i][k] < unloop[j][k]
		}
		return false
	})

	ans := 1000000000
	temp := 0
	for i := 0; i < k-now; i++ {
		if i == len(unloop) {
			temp = ans
			break
		}
		temp += unloop[i][0]
	}
	ans = min(ans, temp)
	if len(loop) != 0 {
		temp = loop[0][0]
		for i := 0; i < len(loop); i++ {
			if i == 0 {
				continue
			}
			unloop = append(unloop, loop[i])
		}
		sort.Slice(unloop, func(i, j int) bool {
			for k := range unloop[i] {
				return unloop[i][k] < unloop[j][k]
			}
			return false
		})
		for i := 0; i < k-now; i++ {
			temp += unloop[i][0]
		}
		ans = min(ans, temp)
	}

	fmt.Println(ans)
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
