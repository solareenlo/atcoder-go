package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var edge, redge [220000][]int
var best int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	diffs := make([]int, 0)
	for i := 0; i < n-1; i++ {
		diffs = append(diffs, abs(a[i]-a[i+1]))
		diffs = append(diffs, abs(a[i]+a[i+1]))
	}
	sort.Ints(diffs)
	diffs = unique(diffs)

	uf := NewDsu(len(diffs))
	same := make([]int, len(diffs))
	for i := range same {
		same[i] = -1
	}
	for i := 0; i < n-1; i++ {
		sub := abs(a[i] - a[i+1])
		add := abs(a[i] + a[i+1])
		to := lowerBound(diffs, sub)
		from := lowerBound(diffs, add)
		if uf.Same(from, to) {
			same[from] = to
			same[to] = from
		}
		uf.Merge(from, to)
		edge[from] = append(edge[from], to)
		redge[to] = append(redge[to], from)
	}

	ans := 0
	for _, g := range uf.Groups() {
		edges := 0
		for _, v := range g {
			edges += len(edge[v])
		}
		if edges > len(g) {
			fmt.Println(-1)
			return
		}
		if edges == len(g) {
			var root int
			for _, v := range g {
				if same[v] != -1 {
					root = v
				}
			}
			other := same[root]
			it := find(edge[root], other)
			cost := 0
			total := 0
			if it != len(edge[root]) { // root -> other
				edge[root] = erase(edge[root], it)
				rit := find(redge[other], root)
				redge[other] = erase(redge[other], rit)
				cost++
				total++
			} else { // other -> root
				it = find(redge[root], other)
				redge[root] = erase(redge[root], it)
				rit := find(edge[other], root)
				edge[other] = erase(edge[other], rit)
			}
			total += dfscost(root, -1)
			length := deps[other] + 1
			for now := other; now != root; now = p[now] {
				cost += cc[now]
			}
			othercost := length - cost
			ans += total - cost + min(cost, othercost)
		} else {
			root := g[0]
			total := dfscost(root, -1)
			best = total
			dfs(root, -1, best)
			ans += best
		}
	}
	fmt.Println((ans + 1) / 2)
}

var p, cc, deps [220000]int

func dfscost(x, last int) int {
	ans := 0
	for _, to := range edge[x] {
		if to == last {
			continue
		}
		p[to] = x
		cc[to] = 0
		deps[to] = deps[x] + 1
		ans += dfscost(to, x)
	}
	for _, to := range redge[x] {
		if to == last {
			continue
		}
		p[to] = x
		cc[to] = 1
		deps[to] = deps[x] + 1
		ans += dfscost(to, x) + 1
	}
	return ans
}

func dfs(x, last, cost int) {
	best = min(best, cost)
	for _, to := range edge[x] {
		if to == last {
			continue
		}
		dfs(to, x, cost+1)
	}
	for _, to := range redge[x] {
		if to == last {
			continue
		}
		dfs(to, x, cost-1)
	}
}

func unique(a []int) []int {
	occurred := map[int]bool{}
	result := []int{}
	for i := range a {
		if occurred[a[i]] != true {
			occurred[a[i]] = true
			result = append(result, a[i])
		}
	}
	// sort.Ints(result)
	n := len(result)
	for i := 0; i < n; i++ {
		a[i] = result[i]
	}
	return result
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

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func find(s []int, x int) int {
	pos := len(s)
	for i := range s {
		if s[i] == x {
			pos = i
			break
		}
	}
	return pos
}

func erase(a []int, pos int) []int {
	return append(a[:pos], a[pos+1:]...)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
