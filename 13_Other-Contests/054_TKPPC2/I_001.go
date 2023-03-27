package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"reflect"
	"sort"
)

const INF = 1061109567
const INFL = 4557430888798830399
const MaxX = 1000000000

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var Q int
	fmt.Fscan(in, &Q)
	points := make([]Points, 0)
	queries := make([]Point, 0)
	for ii := 0; ii < Q; ii++ {
		var ty int
		fmt.Fscan(in, &ty)
		if ty == 1 {
			var L, R int
			fmt.Fscan(in, &L, &R)
			points = append(points, Points{Point{L, MaxX - R}, pair{len(queries), INF}})
		} else if ty == 2 {
			var K int
			fmt.Fscan(in, &K)
			points[K].second.second = len(queries)
		} else if ty == 3 {
			var L, R int
			fmt.Fscan(in, &L, &R)
			queries = append(queries, Point{L, MaxX - R})
		} else {
			os.Exit(0)
		}
	}
	N := 1
	for N < len(queries) {
		N *= 2
	}
	nodePoints := make([][]Point, N*2)
	for _, p := range points {
		x := p.first
		l := p.second.first
		r := p.second.second
		if r == INF {
			r = N
		}
		for l, r = l+N, r+N; l < r; l, r = l>>1, r>>1 {
			if (l & 1) != 0 {
				nodePoints[l] = append(nodePoints[l], x)
				l++
			}
			if (r & 1) != 0 {
				r--
				nodePoints[r] = append(nodePoints[r], x)
			}
		}
	}
	nodeQueries := make([][]Query, N*2)
	for qi := 0; qi < len(queries); qi++ {
		q := Query{qi, queries[qi]}
		for i := qi + N; i > 0; i >>= 1 {
			nodeQueries[i] = append(nodeQueries[i], q)
		}
	}
	ans := make([]int, len(queries))
	for i := 0; i < N*2; i++ {
		if len(nodePoints[i]) != 0 && len(nodeQueries[i]) != 0 {
			solve(nodePoints[i], nodeQueries[i], ans)
		}
	}
	for i := 0; i < len(queries); i++ {
		if ans[i] == 0 {
			fmt.Fprintln(out, -1)
		} else {
			fmt.Fprintln(out, ans[i])
		}
	}
}

func solve(points []Point, queries []Query, ans []int) {
	sort.Slice(points, func(i, j int) bool {
		return points[i].lessThan(points[j])
	})
	last := Point{-INF, INF}
	toBeRemoved := Point{INF, INF}
	for i := range points {
		if last.x < points[i].x && last.y > points[i].y {
			last = points[i]
		} else {
			points[i] = toBeRemoved
		}
	}
	tmp := make([]Point, 0)
	for _, p := range points {
		if !reflect.DeepEqual(p, toBeRemoved) {
			tmp = append(tmp, p)
		}
	}
	points = tmp

	N := 1
	for N < len(points) {
		N *= 2
	}
	nodes := make([]Node, N*2)
	for i := range nodes {
		nodes[i] = Node{INF, -INF}
	}
	neighbors := make([]Neighbors, len(points))
	for i := range neighbors {
		neighbors[i] = Neighbors{-1, -1}
	}
	eventPQ := &Heap{}

	var addEvent func(int)
	addEvent = func(j int) {
		i := neighbors[j].prev
		k := neighbors[j].next
		if i == -1 || k == -1 {
			return
		}
		x := calcCrossX(points[i], points[j], points[k])
		if points[j].x < x {
			heap.Push(eventPQ, Event{x, i, j, k})
		}
	}

	sort.Slice(queries, func(i, j int) bool {
		return queries[i].lessThan(queries[j])
	})
	pi := 0
	for _, q := range queries {
		X := q.p.x
		for ; pi != len(points) && points[pi].x <= X; pi++ {
			lasti := nodes[1].right
			nodes[N+pi] = Node{pi, pi}
			for i := (N + pi) >> 1; i > 0; i >>= 1 {
				nodes[i].left = min(nodes[i].left, pi)
				nodes[i].right = max(nodes[i].right, pi)
			}
			if lasti != -INF {
				neighbors[lasti].next = pi
				neighbors[pi].prev = lasti
				addEvent(lasti)
			}
		}
		for eventPQ.Len() != 0 && (*eventPQ)[0].x <= X {
			e := heap.Pop(eventPQ).(Event)
			if neighbors[e.j].prev != e.i || neighbors[e.j].next != e.k {
				continue
			}

			nodes[N+e.j] = Node{INF, -INF}
			for i := (N + e.j) >> 1; i > 0; i >>= 1 {
				l := nodes[i<<1|0]
				r := nodes[i<<1|1]
				nodes[i] = Node{min(l.left, r.left), max(l.right, r.right)}
			}

			neighbors[e.j] = Neighbors{-1, -1}
			neighbors[e.i].next = e.k
			neighbors[e.k].prev = e.i

			addEvent(e.i)
			addEvent(e.k)
		}

		if nodes[1].left == INF {
			continue
		}

		i := 1
		for i < N {
			a := nodes[i<<1|0].right
			b := nodes[i<<1|1].left
			A := -INFL
			if a != -INF {
				A = calc(points[a], q.p)
			}
			B := -INFL
			if b != INF {
				B = calc(points[b], q.p)
			}
			tmp := 1
			if A >= B {
				tmp = 0
			}
			i = (i << 1) | tmp
		}
		value := calc(points[i-N], q.p)
		ans[q.i] = max(ans[q.i], value)
	}
}

func calc(a, b Point) int {
	return (b.x - a.x) * (b.y - a.y)
}

func calcCrossX(a, b, c Point) int {
	a.minus(b)
	c.minus(b)
	den := a.y*c.x - c.y*a.x
	num1 := a.x * c.x
	num2 := a.y - c.y
	if den < 0 {
		den = -den
		num2 = -num2
	}
	approx := float64(num1) * float64(num2) / float64(den)
	if approx+float64(b.x) > 1.1e9 {
		return INF
	}
	if approx+float64(b.x) < -1e8 {
		return -INF
	}
	quot := int(math.Ceil(approx))
	rem := num1*num2 - den*quot
	for rem > 0 {
		quot++
		rem -= den
	}
	for rem <= -den {
		quot--
		rem += den
	}
	return quot + b.x
}

type Node struct {
	left  int
	right int
}

type Neighbors struct {
	prev int
	next int
}

type Heap []Event

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return !h[i].greaterThan(h[j]) }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(Event)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type Event struct {
	x       int
	i, j, k int
}

func (a Event) greaterThan(b Event) bool {
	return a.x > b.x
}

type Query struct {
	i int
	p Point
}

func (a Query) lessThan(b Query) bool {
	return a.p.lessThan(b.p)
}

type pair struct {
	first  int
	second int
}

type Points struct {
	first  Point
	second pair
}

type Point struct {
	x, y int
}

func (a Point) lessThan(b Point) bool {
	if a.x != b.x {
		return a.x < b.x
	}
	return a.y < b.y
}

func (a *Point) minus(b Point) {
	a.x -= b.x
	a.y -= b.y
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
