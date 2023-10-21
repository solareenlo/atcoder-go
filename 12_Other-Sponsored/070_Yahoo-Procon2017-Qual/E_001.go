package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const INF = int(1e18)

var mv, mv2 [][]int
var sccn, scc, low, st []int
var mt int

func scc_solve(dn, p int) int {
	if low[p] < 0 {
		w := mt
		low[p] = mt
		mt++
		st = append(st, p)
		for _, i := range mv[p] {
			low[p] = min(low[p], scc_solve(dn, i))
		}
		if w == low[p] {
			a := len(sccn)
			sccn = append(sccn, 0)
			tmp := make([]int, 0)
			mv2 = append(mv2, tmp)
			for {
				n := st[len(st)-1]
				st = st[:len(st)-1]
				low[n] = INF
				scc[n] = a
				if n < dn {
					sccn[len(sccn)-1]++
				}
				if n == p {
					break
				}
			}
		}
	}
	return low[p]
}

var ans []int

func solve(n int) int {
	if ans[n] < 0 {
		ans[n] = sccn[n]
		for _, i := range mv2[n] {
			ans[n] = max(ans[n], solve(i)+sccn[n])
		}
	}
	return ans[n]
}

const B = 200

var ok [51419]bool
var md [51419 / B][]int

var df, dt [51419]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var dn int
	fmt.Fscan(in, &dn)
	for i := 0; i < dn; i++ {
		fmt.Fscan(in, &df[i])
	}
	for i := 0; i < dn; i++ {
		fmt.Fscan(in, &dt[i])
	}
	as := dn/B + 1
	mv = make([][]int, dn*2)
	for i := 0; i < dn; i++ {
		mv[i+dn] = append(mv[i+dn], i)
	}
	tmp := make([]pair, 0)
	for i := 0; i < dn; i++ {
		tmp = append(tmp, pair{df[i] + i, i})
		tmp = append(tmp, pair{dt[i] + i, i + dn})
	}
	sortPair(tmp)
	for i := 0; i < dn; i++ {
		ok[i] = false
	}
	for _, i := range tmp {
		if i.y < dn {
			ok[i.y] = true
			md[i.y/B] = append(md[i.y/B], i.y+dn)
		} else {
			n := (i.y-dn)/B + 1
			for j := i.y - dn; j < min(n*B, dn); j++ {
				if ok[j] {
					mv[i.y-dn] = append(mv[i.y-dn], j+dn)
				}
			}
			for j := n; j < as; j++ {
				if len(md[j]) == 1 {
					mv[i.y-dn] = append(mv[i.y-dn], md[j][0])
				} else if len(md[j]) > 1 {
					tmp0 := make([]int, 0)
					mv = append(mv, tmp0)
					for _, k := range md[j] {
						mv[len(mv)-1] = append(mv[len(mv)-1], k)
					}
					md[j] = make([]int, 0)
					md[j] = append(md[j], len(mv)-1)
					mv[i.y-dn] = append(mv[i.y-dn], len(mv)-1)
				}
			}
		}
	}
	tmp = make([]pair, 0)
	for i := 0; i < dn; i++ {
		tmp = append(tmp, pair{df[i] - i, i})
		tmp = append(tmp, pair{dt[i] - i, i + dn})
	}
	sortPair(tmp)
	for i := 0; i < dn; i++ {
		ok[i] = false
	}
	for i := 0; i < as; i++ {
		md[i] = make([]int, 0)
	}
	for _, i := range tmp {
		if i.y < dn {
			ok[i.y] = true
			md[i.y/B] = append(md[i.y/B], i.y+dn)
		} else {
			n := (i.y - dn) / B
			for j := 0; j < n; j++ {
				if len(md[j]) == 1 {
					mv[i.y-dn] = append(mv[i.y-dn], md[j][0])
				}
				if len(md[j]) > 1 {
					tmp0 := make([]int, 0)
					mv = append(mv, tmp0)
					for _, k := range md[j] {
						mv[len(mv)-1] = append(mv[len(mv)-1], k)
					}
					md[j] = make([]int, 0)
					md[j] = append(md[j], len(mv)-1)
					mv[i.y-dn] = append(mv[i.y-dn], len(mv)-1)
				}
			}
			for j := n * B; j < i.y-dn; j++ {
				if ok[j] {
					mv[i.y-dn] = append(mv[i.y-dn], j+dn)
				}
			}
		}
	}
	scc = make([]int, len(mv))
	low = make([]int, len(mv))
	for i := range low {
		low[i] = -1
	}
	for i := 0; i < len(mv); i++ {
		scc_solve(dn, i)
	}
	for i := 0; i < len(mv); i++ {
		for _, j := range mv[i] {
			if scc[i] != scc[j] {
				mv2[scc[i]] = append(mv2[scc[i]], scc[j])
			}
		}
	}
	ans = make([]int, len(mv2))
	for i := range ans {
		ans[i] = -1
	}
	r := 0
	for i := 0; i < len(mv2); i++ {
		r = max(r, solve(i))
	}
	fmt.Println(r)
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

type pair struct {
	x, y int
}

func sortPair(tmp []pair) {
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].x == tmp[j].x {
			return tmp[i].y < tmp[j].y
		}
		return tmp[i].x < tmp[j].x
	})
}
