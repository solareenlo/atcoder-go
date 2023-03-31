package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const MAX = 100002
const MAX_LOG = 22
const MIN_VAL = -int(1e18)

var n, m, k int
var pos []int
var ed, mx [MAX]int
var dp1, dp2 []int
var vv []pair
var rig [MAX]int
var query [MAX][]pair
var rr [MAX][]int
var bas [MAX]int
var ms MAX_SP
var dd [MAX]int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &m, &k)
	for i := 0; i < k; i++ {
		var p int
		fmt.Fscan(in, &p)
		p--
		pos = append(pos, p)
	}
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		vv = append(vv, pair{a, b})
		rr[a] = append(rr[a], b)
	}
	dp1 = make([]int, n)
	solve(dp1, vv)
	for i := 0; i < len(vv); i++ {
		vv[i].x, vv[i].y = vv[i].y, vv[i].x
		vv[i].x = n - vv[i].x - 1
		vv[i].y = n - vv[i].y - 1
	}
	dp2 = make([]int, n)
	solve(dp2, vv)
	for i := 0; i < len(vv); i++ {
		vv[i].x, vv[i].y = vv[i].y, vv[i].x
		vv[i].x = n - vv[i].x - 1
		vv[i].y = n - vv[i].y - 1
	}
	dp2 = reverseOrderInt(dp2)
	ans := 0
	for i := -1; i < n; i++ {
		nex := upperBound(pos, i)
		if nex == len(pos) {
			break
		}
		nex = pos[nex]
		lef := i + 1
		rig := nex
		cost := 0
		if i >= 0 {
			cost += dp1[i]
		}
		query[lef] = append(query[lef], pair{rig, cost})
	}

	ff := -1
	uu := 0
	dp2 = append(dp2, 0)
	for i := 0; i < len(pos); i++ {
		cur := pos[i]
		cost := 0
		if cur != 0 {
			cost += dp1[cur-1]
		}
		if cur+1 < n {
			cost += dp2[cur+1]
		}
		ans = max(ans, cost)
	}
	for i := len(dp2) - 1; i >= 0; i-- {
		bas[i] = uu
		dp2[i] = -uu + dp2[i]
		uu++
	}
	ms.set_array(dp2)
	dd[0] = 0
	ff = -1
	ff = -1
	for i := 0; i < n; i++ {
		for _, Go := range rr[i] {
			ff = max(ff, Go)
		}
		for _, it := range query[i] {
			if it.x <= ff {
				L := it.x
				R := ff
				cost := it.y
				cost += bas[i] - (-ms.get_max(L+1, R+1))
				ans = max(ans, cost)
			}
		}
	}
	fmt.Println(ans)
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
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

type MAX_SP struct {
	lcc  [MAX_LOG][]int
	l2   []int
	SIZE int
}

func (m *MAX_SP) init(n int) {
	m.SIZE = n
	emp := make([]int, n)
	for i := range emp {
		emp[i] = MIN_VAL
	}
	for i := 0; i < MAX_LOG; i++ {
		m.lcc[i] = make([]int, n)
		copy(m.lcc[i], emp)
	}
	m.l2 = make([]int, m.SIZE+1)
	for i := 0; i <= m.SIZE; i++ {
		m.l2[i] = log2(i)
	}
}

func (m *MAX_SP) set_array(A []int) {
	if m.SIZE < len(A) {
		m.SIZE = len(A)
		m.init(m.SIZE)
	}
	for i := 0; i < m.SIZE; i++ {
		m.lcc[0][i] = A[i]
	}
	for i := 0; i+1 < MAX_LOG; i++ {
		for j := 0; j+(1<<i) < m.SIZE; j++ {
			m.lcc[i+1][j] = max(m.lcc[i][j], m.lcc[i][j+(1<<i)])
		}
	}
}

func (m MAX_SP) get_max(a, b int) int {
	rng := b - a + 1
	L := m.l2[rng]
	return max(m.lcc[L][a], m.lcc[L][b-(1<<L)+1])
}

type pair struct {
	x, y int
}

func solve(dp []int, v []pair) {
	for i := 0; i < n; i++ {
		ed[i] = 1145141919
	}
	for i := 0; i < len(v); i++ {
		ed[v[i].y] = Min(ed[v[i].y], v[i].x)
	}
	for i := 0; i < n; i++ {
		dp[i] = 0
	}
	base := 0
	d := make([]pair, 0)
	d = append(d, pair{-1, 0})
	for i := 0; i < n; i++ {
		base++
		if i != 0 {
			dp[i] = dp[i-1]
		}
		cost2 := base - dp[i]
		if ed[i] == 1145141919 {
			for len(d) != 0 && d[len(d)-1].y > cost2 {
				d = d[:len(d)-1]
			}
			d = append(d, pair{i, cost2})
			continue
		}
		lef := ed[i] - 1
		mm := dp[i]
		if lef >= 0 {
			mm = max(mm, mx[lef])
		}
		id := LowerBound(d, pair{lef, -1})
		if id != len(d) {
			mm = max(mm, base-d[id].y)
		}
		dp[i] = mm
		cost := base - mm
		for len(d) != 0 && d[len(d)-1].y > cost {
			d = d[:len(d)-1]
		}
		d = append(d, pair{i, cost})
	}
}

func LowerBound(a []pair, x pair) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i].x >= x.x
	})
	return idx
}

func Min(a, b int) int {
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

func log2(n int) int {
	var k int
	for k = 0; n != 0; n >>= 1 {
		k++
	}
	return k - 1
}
