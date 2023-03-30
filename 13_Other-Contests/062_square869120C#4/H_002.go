package main

import (
	"fmt"
	"sort"
)

const MAXN = 200
const MAXM = MAXN - 1

type pair struct {
	x, y int
}

var n, nq int
var locghead [MAXN]int
var locgnxt, locgto [2 * MAXM]int
var local bool = false
var qloc QLOC
var inq [MAXN]bool
var can [MAXN][MAXN]bool
var ghead [MAXN]int
var gnxt, gto [2 * MAXM]int
var m int
var ret []pair
var nret int
var q [MAXN]int
var qhead, qtail int
var col [MAXN]int

func main() {
	ret = make([]pair, MAXM)

	fmt.Scan(&n)
	solve()
	fmt.Printf("!")
	for i := 0; i < nret; i++ {
		fmt.Printf(" (%d,%d)", ret[i].x, ret[i].y)
	}
	fmt.Println()
}

func solve() {
	nret = 0
	m = 0
	nq = 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			can[i][j] = i != j
		}
	}
	for i := 0; i < n; i++ {
		ghead[i] = -1
	}
	for {
		colorize()
		i := -1
		var icnt int
		for ii := 0; ii < n; ii++ {
			cnt := 0
			for j := 0; j < n; j++ {
				if can[ii][j] {
					cnt++
				}
			}
			if cnt > 0 && (i == -1 || cnt > icnt) {
				i = ii
				icnt = cnt
			}
		}
		if i == -1 {
			break
		}
		curcol := -1
		for j := 0; j < n; j++ {
			if can[i][j] {
				curcol = col[j]
				break
			}
		}
		opt := make([]int, 0)
		for j := 0; j < n; j++ {
			if can[i][j] && col[j] == curcol {
				opt = append(opt, j)
			}
		}
		one := make([]int, 0)
		now := make([]int, len(opt))
		copy(now, opt)
		now = append(now, i)
		cur := query(now)
		if cur == 0 {
			clear(now)
			continue
		} else if cur == 1 && len(now) > len(one) {
			copy(one, now)
		}
		var lst int
		l := -1
		r := len(opt) - 1
		for l+1 < r {
			m := l + (r-l)/2
			now := make([]int, 0)
			for j := 0; j <= m; j++ {
				now = append(now, opt[j])
			}
			now = append(now, i)
			cur := query(now)
			if cur == 0 {
				clear(now)
				l = m
			} else {
				r = m
				if cur == 1 && len(now) > len(one) {
					copy(one, now)
				}
			}
		}
		lst = r
		now = make([]int, 0)
		now = append(now, i)
		now = append(now, opt[lst])
		if lst > 0 && query(now) == 0 {
			clear(now)
		} else {
			found(i, opt[lst])
			clear(one)
			continue
		}
		var fst int
		l = 0
		r = lst
		for l+1 < r {
			m := l + (r-l)/2
			now := make([]int, 0)
			for j := m; j <= lst; j++ {
				now = append(now, opt[j])
			}
			cur := query(now)
			if cur == 0 {
				clear(now)
				r = m
			} else {
				l = m
				if cur == 1 && len(now) > len(one) {
					copy(one, now)
				}
			}
		}
		fst = l
		found(opt[fst], opt[lst])
		clear(one)
	}
	tmp := ret[:nret]
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].x == tmp[j].x {
			return tmp[i].y < tmp[j].y
		}
		return tmp[i].x < tmp[j].x
	})
}

func colorize() {
	for i := 0; i < n; i++ {
		col[i] = -1
	}
	for i := 0; i < n; i++ {
		if col[i] == -1 {
			qhead = 0
			qtail = 0
			col[i] = 0
			q[qhead] = i
			qhead++
			for qtail < qhead {
				at := q[qtail]
				qtail++
				for x := ghead[at]; x != -1; x = gnxt[x] {
					to := gto[x]
					if col[to] != -1 {
						continue
					}
					col[to] = 1 - col[at]
					q[qhead] = to
					qhead++
				}
			}
			for j := 0; j < qhead; j++ {
				for k := j + 1; k < qhead; k++ {
					a := q[j]
					b := q[k]
					can[a][b] = false
					can[b][a] = false
				}
			}
		}
	}
}

func clear(now []int) {
	for i := 0; i < len(now); i++ {
		for j := i + 1; j < len(now); j++ {
			a := now[i]
			b := now[j]
			can[a][b] = false
			can[b][a] = false
		}
	}
}

func found(a, b int) {
	ret[nret] = pair{min(a, b), max(a, b)}
	nret++
	can[a][b] = false
	can[b][a] = false
	gnxt[2*m+0] = ghead[a]
	ghead[a] = 2*m + 0
	gto[2*m+0] = b
	gnxt[2*m+1] = ghead[b]
	ghead[b] = 2*m + 1
	gto[2*m+1] = a
	m++
}

func query(now []int) int {
	nq++
	for i := 0; i < n; i++ {
		inq[i] = false
	}
	for i := 0; i < len(now); i++ {
		inq[now[i]] = true
	}
	if local {
		for i := 0; i < n; i++ {
			qloc.alive[i] = inq[i]
		}
		return qloc.calc()
	} else {
		fmt.Printf("? ")
		for i := 0; i < n; i++ {
			if inq[i] {
				fmt.Printf("%d", 1)
			} else {
				fmt.Printf("%d", 0)
			}
		}
		fmt.Println()
		var ret int
		fmt.Scan(&ret)
		return ret
	}
}

type QLOC struct {
	alive        [MAXN]bool
	state        [MAXN]int
	q            [MAXN]int
	qhead, qtail int
	d            [MAXN]int
}

func (q *QLOC) calc() int {
	ret := 0
	for i := 0; i < n; i++ {
		q.state[i] = 0
	}
	for i := 0; i < n; i++ {
		if q.state[i] == 0 && q.alive[i] {
			q.qhead = 0
			q.qtail = 0
			q.q[q.qhead] = i
			q.qhead++
			q.d[i] = 0
			q.state[i] = 1
			for q.qtail < q.qhead {
				at := q.q[q.qtail]
				q.qtail++
				for x := locghead[at]; x != -1; x = locgnxt[x] {
					to := locgto[x]
					if !q.alive[to] || q.state[to] == q.state[at] {
						continue
					}
					q.q[q.qhead] = to
					q.qhead++
					q.d[to] = q.d[at] + 1
					q.state[to] = q.state[at]
				}
			}
			j := q.q[q.qhead-1]
			q.qhead = 0
			q.qtail = 0
			q.q[q.qhead] = j
			q.qhead++
			q.d[j] = 0
			q.state[j] = 2
			for q.qtail < q.qhead {
				at := q.q[q.qtail]
				q.qtail++
				for x := locghead[at]; x != -1; x = locgnxt[x] {
					to := locgto[x]
					if !q.alive[to] || q.state[to] == q.state[at] {
						continue
					}
					q.q[q.qhead] = to
					q.qhead++
					q.d[to] = q.d[at] + 1
					q.state[to] = q.state[at]
				}
			}
			k := q.q[q.qhead-1]
			diam := q.d[k]
			ret += diam * diam
		}
	}
	return ret
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
