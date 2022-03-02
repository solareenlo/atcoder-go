package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)

	r := make([]pair, 0)
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		r = append(r, pair{x, y})
	}
	sort.Slice(r, func(i, j int) bool {
		return r[i].x < r[j].x
	})
	r = reverseOrder(r)

	v := make([]pair, 0)
	vx := make([]int, 300000)
	vy := make([]int, 300000)
	y := -1
	xl, yl := 0, 0
	for _, u := range r {
		if u.y > y {
			y = u.y
			v = append(v, u)
			xl++
			vx[xl] = u.x
			yl++
			vy[yl] = u.y
		}
	}
	r = v
	r = reverseOrder(r)

	b := make([]pair, 0)
	for i := 0; i < m; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		b = append(b, pair{x, y})
		xl++
		vx[xl] = x
		yl++
		vy[yl] = y
	}

	tmp := vx[1 : xl+1]
	sort.Ints(tmp)
	uni := unique(vx[1 : xl+1])
	xl = len(uni)
	for i := 0; i < xl; i++ {
		vx[i+1] = uni[i]
	}
	tmp = vy[1 : yl+1]
	sort.Ints(tmp)
	uni = unique(vy[1 : yl+1])
	yl = len(uni)
	for i := 0; i < yl; i++ {
		vy[i+1] = uni[i]
	}

	g := newMinCostFlow(xl + yl)
	for i := 0; i < len(r)-1; i++ {
		x := lowerBound(vx[1:xl+1], r[i].x)
		y := lowerBound(vy[1:yl+1], r[i+1].y) + xl
		g.AddEdge(x, y, k, 0)
	}
	for _, u := range b {
		xx := lowerBound(vy[1:yl+1], u.y) + xl
		yy := lowerBound(vx[1:xl+1], u.x)
		g.AddEdge(xx, yy, 1, 0)
	}
	for i := 1; i < xl; i++ {
		g.AddEdge(i-1, i+1-1, k, vx[i+1]-vx[i])
		g.AddEdge(i+1-1, i-1, k, 0)
	}
	for i := 1; i < yl; i++ {
		g.AddEdge(i+1+xl-1, i+xl-1, k, vy[i+1]-vy[i])
		g.AddEdge(i+xl-1, i+1+xl-1, k, 0)
	}

	x := lowerBound(vy[1:yl+1], r[0].y) + xl
	y = lowerBound(vx[1:xl+1], r[len(r)-1].x)
	tm := g.SlopeL(x, y, k)
	leng := len(tm)
	fmt.Println(tm[leng-1][1])
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

type MinCostFlow struct {
	n   int
	pos [][2]int
	g   [][]_Edge
}
type _Edge struct{ to, rev, capa, cost int }
type Edge struct{ from, to, capa, flow, cost int }

func newMinCostFlow(n int) *MinCostFlow {
	return &MinCostFlow{n: n, g: make([][]_Edge, n)}
}
func (mcf *MinCostFlow) AddEdge(from, to, capa, cost int) int {
	m := len(mcf.pos)
	mcf.pos = append(mcf.pos, [2]int{from, len(mcf.g[from])})
	mcf.g[from] = append(mcf.g[from], _Edge{to, len(mcf.g[to]), capa, cost})
	mcf.g[to] = append(mcf.g[to], _Edge{from, len(mcf.g[from]) - 1, 0, -cost})
	return m
}
func (mcf *MinCostFlow) GetEdge(i int) Edge {
	e := mcf.g[mcf.pos[i][0]][mcf.pos[i][1]]
	re := mcf.g[e.to][e.rev]
	return Edge{mcf.pos[i][0], e.to, e.capa + re.capa, re.capa, e.cost}
}
func (mcf *MinCostFlow) Edges() []Edge {
	m := len(mcf.pos)
	res := make([]Edge, m)
	for i := 0; i < m; i++ {
		res[i] = mcf.GetEdge(i)
	}
	return res
}
func (mcf *MinCostFlow) Flow(s, t int) [2]int {
	res := mcf.Slope(s, t)
	return res[len(res)-1]
}
func (mcf *MinCostFlow) FlowL(s, t, flowLim int) [2]int {
	res := mcf.SlopeL(s, t, flowLim)
	return res[len(res)-1]
}
func (mcf *MinCostFlow) Slope(s, t int) [][2]int {
	return mcf.SlopeL(s, t, int(1e+18))
}
func (mcf *MinCostFlow) SlopeL(s, t, flowLim int) [][2]int {
	dual, dist := make([]int, mcf.n), make([]int, mcf.n)
	pv, pe := make([]int, mcf.n), make([]int, mcf.n)
	vis := make([]bool, mcf.n)
	dualRef := func() bool {
		for i := 0; i < mcf.n; i++ {
			dist[i], pv[i], pe[i] = int(1e+18), -1, -1
			vis[i] = false
		}
		pq := make(PriorityQueue, 0)
		heap.Init(&pq)
		item := &Item{value: s, priority: 0}
		dist[s] = 0
		heap.Push(&pq, item)
		for pq.Len() != 0 {
			v := heap.Pop(&pq).(*Item).value
			if vis[v] {
				continue
			}
			vis[v] = true
			if v == t {
				break
			}
			for i := 0; i < len(mcf.g[v]); i++ {
				e := mcf.g[v][i]
				if vis[e.to] || e.capa == 0 {
					continue
				}
				cost := e.cost - dual[e.to] + dual[v]
				if dist[e.to]-dist[v] > cost {
					dist[e.to] = dist[v] + cost
					pv[e.to] = v
					pe[e.to] = i
					item := &Item{value: e.to, priority: dist[e.to]}
					heap.Push(&pq, item)
				}
			}
		}
		if !vis[t] {
			return false
		}
		for v := 0; v < mcf.n; v++ {
			if !vis[v] {
				continue
			}
			dual[v] -= dist[t] - dist[v]
		}
		return true
	}
	flow, cost, prevCost := 0, 0, -1
	res := make([][2]int, 0, mcf.n)
	res = append(res, [2]int{flow, cost})
	for flow < flowLim {
		if !dualRef() {
			break
		}
		c := flowLim - flow
		for v := t; v != s; v = pv[v] {
			c = mcf.Min(c, mcf.g[pv[v]][pe[v]].capa)
		}
		for v := t; v != s; v = pv[v] {
			mcf.g[pv[v]][pe[v]].capa -= c
			mcf.g[v][mcf.g[pv[v]][pe[v]].rev].capa += c
		}
		d := -dual[s]
		flow += c
		cost += c * d
		if prevCost == d {
			res = res[:len(res)-1]
		}
		res = append(res, [2]int{flow, cost})
		prevCost = cost
	}
	return res
}
func (mcf *MinCostFlow) Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Item struct{ value, priority, index int }
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
func (pq *PriorityQueue) update(item *Item, value int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
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
	sort.Ints(result)
	return result
}

type pair struct{ x, y int }

func reverseOrder(a []pair) []pair {
	n := len(a)
	res := make([]pair, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
