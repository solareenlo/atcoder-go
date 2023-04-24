package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct {
	x, y int
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, K int
	fmt.Fscan(in, &n, &m, &K)
	root := n + m + 1
	var mcc MinCostCirculation
	mcc.init(root + 1)

	Ed := make([]pair, 0)

	Map = make(map[pair]int)
	for i := 0; i < K; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		Map[pair{b, a}] = i + 1
		Map[pair{a, b}] = i + 1
		Deg[a]++
		Deg[b]++
		Ed = append(Ed, pair{a, b})
	}
	for _, tmp := range Ed {
		a := tmp.x
		b := tmp.y
		if Deg[a]%2 == Deg[b]%2 {
			mcc.addEdge(a, b, 0, 1, -1)
		} else {
			mcc.addEdge(b, a, 0, 1, -1)
		}
	}
	for i := 1; i <= n; i++ {
		mcc.addEdge(root, i, 0, 1, 0)
	}
	for i := n + 1; i <= n+m; i++ {
		mcc.addEdge(i, root, 0, 1, 0)
	}

	sln := mcc.solve()
	edges := mcc.Edges()

	fmt.Printf("%d\n", -sln.y)

	c := 0
	var OO [6010]int

	for _, tmp := range edges {
		a := tmp.x
		b := tmp.y
		if a != root && b != root {
			if (Deg[a]%2 == Deg[b]%2) != (a < b) {
				c++
				E[b] = append(E[b], pair{a, c})
				OO[a]--
				OO[b]++
			}
		}
	}
	for {
		val := 0
		for i := 1; i <= n; i++ {
			if OO[i] != 0 {
				TV = make([]int, 0)
				TE = make([]int, 0)
				ck := DFS(i)
				if ck == 2 {
					OO[i]--
				}
				val |= ck
				if ck != 0 {
					break
				}
			}
		}
		if val == 0 {
			break
		}
	}
	for {
		val := false
		for i := 1; i <= n; i++ {
			TV = make([]int, 0)
			TE = make([]int, 0)
			ck := DFS(i)
			if ck != 0 {
				val = true
				break
			}
		}
		if !val {
			break
		}
	}
	for _, T := range AA {
		l := len(T)
		ck := 0
		if T[0] > n {
			ck = 1
		}
		for i := 0; i < l; i++ {
			if i%2 == ck {
				add(T[i], T[(i+1)%l])
			}
		}
		for i := 0; i < l; i++ {
			if i%2 != ck {
				add(T[i], T[(i+1)%l])
			}
		}
	}
	for _, T := range BB {
		l := len(T)
		vv := make([]bool, l)
		for {
			ck := false
			for i := 0; i < l-1; i++ {
				if !vv[i] && add(T[i], T[i+1]) {
					ck = true
					vv[i] = true
					break
				}
			}
			if !ck {
				break
			}
		}
	}
	for _, t := range res {
		fmt.Printf("%d ", t)
	}
}

var E [6100][]pair
var TV, TE []int
var Vis [6100]bool
var vE [10000000]int
var AA, BB [][]int

func DFS(a int) int {
	if Vis[a] {
		TV = reverseOrderInt(TV)
		TE = reverseOrderInt(TE)
		r := make([]int, 0)
		ck := false
		for i := 0; i < len(TV); i++ {
			Vis[TV[i]] = false
			if !ck {
				r = append(r, TV[i])
			} else {
				vE[TE[i]] = 0
			}
			if TV[i] == a {
				ck = true
			}
		}
		r = reverseOrderInt(r)
		AA = append(AA, r)
		return 1
	}
	TV = append(TV, a)
	Vis[a] = true
	for _, tmp := range E[a] {
		x := tmp.x
		num := tmp.y
		if vE[num] != 0 {
			continue
		}
		vE[num] = 1
		TE = append(TE, num)
		return DFS(x)
	}
	for _, x := range TV {
		Vis[x] = false
	}
	if len(TV) == 1 {
		return 0
	}
	BB = append(BB, TV)
	return 2
}

var Map map[pair]int
var Deg, Used [6010]int
var res []int
var S int

func add(a, b int) bool {
	if Deg[a]%2 != (Deg[b]-Used[b])%2 {
		return false
	}
	if Deg[b]%2 != (Deg[a]-Used[a])%2 {
		return false
	}
	Used[a]++
	Used[b]++
	res = append(res, Map[pair{a, b}])
	S++
	return true
}

type EdgeHLPP struct {
	to, inv, rem, Cap int
}

type HLPP struct {
	G                            [][]EdgeHLPP
	excess                       []int
	hei, arc, prv, nxt, act, bot []int
	Q                            []int
	n, high, cut, work           int
}

func (h *HLPP) init(k int) {
	h.G = make([][]EdgeHLPP, k)
}

func (h *HLPP) addEdge(u, v, Cap, rcap int) int {
	h.G[u] = append(h.G[u], EdgeHLPP{v, len(h.G[v]), Cap, Cap})
	h.G[v] = append(h.G[v], EdgeHLPP{u, len(h.G[u]) - 1, rcap, rcap})
	return len(h.G[u]) - 1
}

func (hl *HLPP) raise(v, h int) {
	hl.nxt[hl.prv[v]] = hl.nxt[v]
	hl.prv[hl.nxt[hl.prv[v]]] = hl.prv[v]
	hl.hei[v] = h
	if hl.excess[v] > 0 {
		hl.bot[v] = hl.act[h]
		hl.act[h] = v
		hl.high = max(hl.high, h)
	}
	if h < hl.n {
		hl.cut = max(hl.cut, h+1)
	}
	h += hl.n
	hl.prv[v] = h
	hl.nxt[v] = hl.nxt[hl.prv[v]]
	hl.nxt[h] = v
	hl.prv[hl.nxt[hl.nxt[h]]] = v
}

func (h *HLPP) global(s, t int) {
	h.hei = make([]int, h.n)
	for i := range h.hei {
		h.hei[i] = h.n * 2
	}
	h.act = make([]int, h.n*2)
	for i := range h.act {
		h.act[i] = -1
	}
	for i := range h.prv {
		h.prv[i] = i
	}
	for i := range h.nxt {
		h.nxt[i] = i
	}
	h.hei[t] = 0
	h.high = 0
	h.cut = 0
	h.work = 0
	h.hei[s] = h.n
	for _, x := range []int{t, s} {
		h.Q = append(h.Q, x)
		for len(h.Q) > 0 {
			v := h.Q[0]
			h.Q = h.Q[1:]
			for _, e := range h.G[v] {
				if h.hei[e.to] == h.n*2 && h.G[e.to][e.inv].rem != 0 {
					h.Q = append(h.Q, e.to)
					h.raise(e.to, h.hei[v]+1)
				}
			}
		}
	}
}

func (h *HLPP) push(v int, e EdgeHLPP, z bool) {
	f := min(h.excess[v], e.rem)
	if f > 0 {
		if z && h.excess[e.to] == 0 {
			h.bot[e.to] = h.act[h.hei[e.to]]
			h.act[h.hei[e.to]] = e.to
		}
		e.rem -= f
		h.G[e.to][e.inv].rem += f
		h.excess[v] -= f
		h.excess[e.to] += f
	}
}

func (hl *HLPP) discharge(v int) {
	h := hl.n * 2
	k := hl.hei[v]

	for j := 0; j < len(hl.G[v]); j++ {
		e := hl.G[v][hl.arc[v]]
		if e.rem != 0 {
			if k == hl.hei[e.to]+1 {
				hl.push(v, e, true)
				if hl.excess[v] <= 0 {
					return
				}
			} else {
				h = min(h, hl.hei[e.to]+1)
			}
		}
		hl.arc[v]++
		if hl.arc[v] >= len(hl.G[v]) {
			hl.arc[v] = 0
		}
	}

	if k < hl.n && hl.nxt[k+hl.n] == hl.prv[k+hl.n] {
		for j := k; j < hl.cut; j++ {
			for hl.nxt[j+hl.n] < hl.n {
				hl.raise(hl.nxt[j+hl.n], hl.n)
			}
		}
		hl.cut = k
	} else {
		hl.raise(v, h)
		hl.work++
	}
}

func (h *HLPP) flow(src, dst int) int {
	h.n = len(h.G)
	h.excess = make([]int, h.n)
	h.arc = make([]int, h.n)
	h.prv = make([]int, h.n*3)
	h.nxt = make([]int, h.n*3)
	h.bot = make([]int, h.n)
	for _, e := range h.G[src] {
		h.excess[src] = e.rem
		h.push(src, e, false)
	}

	h.global(src, dst)

	for ; h.high > 0; h.high-- {
		for h.act[h.high] != -1 {
			v := h.act[h.high]
			h.act[h.high] = h.bot[v]
			if v != src && h.hei[v] == h.high {
				h.discharge(v)
				if h.work > 4*h.n {
					h.global(src, dst)
				}
			}
		}
	}

	return h.excess[dst]
}

func (h HLPP) getFlow(v, e int) int {
	return h.G[v][e].Cap - h.G[v][e].rem
}

func (h HLPP) cutSide(v int) bool { return h.hei[v] >= h.n }

type Circulation struct {
	lowerBoundSum int
	mf            HLPP
}

func (c *Circulation) init(k int) {
	c.mf.init(k + 2)
}

func (c *Circulation) addEdge(s, e, l, r int) {
	c.mf.addEdge(s+2, e+2, r-l, 0)
	if l > 0 {
		c.mf.addEdge(0, e+2, l, 0)
		c.mf.addEdge(s+2, 1, l, 0)
		c.lowerBoundSum += l
	} else {
		c.mf.addEdge(0, s+2, -l, 0)
		c.mf.addEdge(e+2, 1, -l, 0)
		c.lowerBoundSum += -l
	}
}

func (c *Circulation) solve(s, e int) bool {
	return c.lowerBoundSum == c.mf.flow(0, 1)
}

const SCALE = 3
const INF = int(1e18)

type EdgeStack struct {
	s, e, l, r, cost int
}

type Edge struct {
	pos, rev, rem, Cap, cost int
}

type MinCostCirculation struct {
	n    int
	estk []EdgeStack
	circ Circulation
	gph  [][]Edge
	p    []int
}

func (mcc *MinCostCirculation) init(k int) {
	mcc.n = k
	mcc.circ.init(k)
	mcc.gph = make([][]Edge, k)
	mcc.p = make([]int, k)
}

func (mcc *MinCostCirculation) addEdge(s, e, l, r, cost int) {
	mcc.estk = append(mcc.estk, EdgeStack{s, e, l, r, cost})
}

type P struct {
	x bool
	y int
}

func (mcc *MinCostCirculation) solve() P {
	for _, i := range mcc.estk {
		if i.s != i.e {
			mcc.circ.addEdge(i.s, i.e, i.l, i.r)
		}
	}
	if !mcc.circ.solve(-1, -1) {
		return P{false, 0}
	}
	ptr := make([]int, mcc.n)
	eps := 0
	for _, i := range mcc.estk {
		var curFlow int
		if i.s != i.e {
			curFlow = i.r - mcc.circ.mf.G[i.s+2][ptr[i.s]].rem
		} else {
			curFlow = i.r
		}
		srev := len(mcc.gph[i.e])
		erev := len(mcc.gph[i.s])
		if i.s == i.e {
			srev++
		}
		mcc.gph[i.s] = append(mcc.gph[i.s], Edge{i.e, srev, i.r - curFlow, i.r, i.cost * (mcc.n + 1)})
		mcc.gph[i.e] = append(mcc.gph[i.e], Edge{i.s, erev, -i.l + curFlow, -i.l, -i.cost * (mcc.n + 1)})
		eps = max(eps, abs(i.cost)*(mcc.n+1))
		if i.s != i.e {
			ptr[i.s] += 2
			ptr[i.e] += 2
		}
	}
	for {
		var cost func(Edge, int, int) int
		cost = func(e Edge, s, t int) int {
			return e.cost + mcc.p[s] - mcc.p[t]
		}
		eps = 0
		for i := 0; i < mcc.n; i++ {
			for _, e := range mcc.gph[i] {
				if e.rem > 0 {
					eps = max(eps, -cost(e, i, e.pos))
				}
			}
		}
		if eps <= 1 {
			break
		}
		eps = max(1, eps>>SCALE)
		excess := make([]int, mcc.n)
		que := make([]int, 0)
		var push func(*Edge, int, int)
		push = func(e *Edge, src, flow int) {
			(*e).rem -= flow
			mcc.gph[(*e).pos][(*e).rev].rem += flow
			excess[src] -= flow
			excess[(*e).pos] += flow
			if excess[(*e).pos] <= flow && excess[(*e).pos] > 0 {
				que = append(que, (*e).pos)
			}
		}
		ptr := make([]int, mcc.n)
		var relabel func(int)
		relabel = func(v int) {
			ptr[v] = 0
			mcc.p[v] = -INF
			for _, e := range mcc.gph[v] {
				if e.rem > 0 {
					mcc.p[v] = max(mcc.p[v], mcc.p[e.pos]-e.cost-eps)
				}
			}
		}
		for i := 0; i < mcc.n; i++ {
			for j := range mcc.gph[i] {
				if mcc.gph[i][j].rem > 0 && cost(mcc.gph[i][j], i, mcc.gph[i][j].pos) < 0 {
					push(&mcc.gph[i][j], i, mcc.gph[i][j].rem)
				}
			}
		}
		for len(que) > 0 {
			x := que[0]
			que = que[1:]
			for excess[x] > 0 {
				for ; ptr[x] < len(mcc.gph[x]); ptr[x]++ {
					e := mcc.gph[x][ptr[x]]
					if e.rem > 0 && cost(e, x, e.pos) < 0 {
						push(&mcc.gph[x][ptr[x]], x, min(e.rem, excess[x]))
						if excess[x] == 0 {
							break
						}
					}
				}
				if excess[x] == 0 {
					break
				}
				relabel(x)
			}
		}
	}
	ans := 0
	for i := 0; i < mcc.n; i++ {
		for _, j := range mcc.gph[i] {
			j.cost /= (mcc.n + 1)
			ans += j.cost * (j.Cap - j.rem)
		}
	}
	return P{true, ans / 2}
}

func (mcc MinCostCirculation) Edges() []pair {
	ret := make([]pair, 0)
	for i := 0; i < mcc.n; i++ {
		for _, j := range mcc.gph[i] {
			if j.rem > 0 {
				ret = append(ret, pair{i, j.pos})
			}
		}
	}
	return ret
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

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
