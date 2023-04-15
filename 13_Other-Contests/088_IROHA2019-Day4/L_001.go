package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type Coin struct {
	tl, tr, v, x int
}

var coins []Coin
var queries []int
var Ans []float64

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var Q int
	fmt.Fscan(in, &Q)
	coin_pos_to_id := make(map[int]int)
	for i := 0; i < Q; i++ {
		var t int
		fmt.Fscan(in, &t)
		if t == 1 {
			var x, v int
			fmt.Fscan(in, &x, &v)
			coin_pos_to_id[x] = len(coins)
			coins = append(coins, Coin{len(queries), -1, v, x})
		}
		if t == 2 {
			var x int
			fmt.Fscan(in, &x)
			coins[coin_pos_to_id[x]].tr = len(queries)
			delete(coin_pos_to_id, x)
		}
		if t == 3 {
			var x int
			fmt.Fscan(in, &x)
			queries = append(queries, x)
		}
	}

	query_t_max := len(queries)
	for a := range coins {
		if coins[a].tr < 0 {
			coins[a].tr = query_t_max
		}
	}

	sort.Slice(coins, func(l, r int) bool {
		return coins[l].x < coins[r].x
	})

	Ans = make([]float64, query_t_max)

	for way := 0; way < 2; way++ {
		allcoins := make([]int, len(coins))
		for i := 0; i < len(allcoins); i++ {
			allcoins[i] = i
		}
		sort.Slice(allcoins, func(l, r int) bool {
			return coins[allcoins[l]].x < coins[allcoins[r]].x
		})
		allqueries := make([]int, len(queries))
		for i := 0; i < len(allqueries); i++ {
			allqueries[i] = i
		}
		sort.Slice(allqueries, func(l, r int) bool {
			return queries[allqueries[l]] < queries[allqueries[r]]
		})
		solve(0, len(queries), allcoins, allqueries)
		for c := range coins {
			coins[c].x *= -1
		}
		for x := range queries {
			queries[x] *= -1
		}
	}

	for _, a := range Ans {
		fmt.Fprintln(out, a)
	}
}

type Frac struct {
	a, b int
}

func (f *Frac) init(a, b int) {
	f.a = a
	f.b = b
	if b < 0 {
		f.a *= -1
		f.b *= -1
	}
}

func (l Frac) lessThan(r Frac) bool {
	return l.a*r.b < r.a*l.b
}

func solve(l, r int, coinid, queryid []int) {
	var coinidL, coinidM, coinidR []int
	var queryidL, queryidR []int
	m := (l + r) / 2
	for _, i := range coinid {
		tl := coins[i].tl
		tr := coins[i].tr
		if tl <= l && r <= tr {
			coinidM = append(coinidM, i)
		} else {
			if tl < m {
				coinidL = append(coinidL, i)
			}
			if m < tr {
				coinidR = append(coinidR, i)
			}
		}
	}
	for _, i := range queryid {
		if i < m {
			queryidL = append(queryidL, i)
		} else {
			queryidR = append(queryidR, i)
		}
	}
	if r-l > 1 {
		solve(l, m, coinidL, queryidL)
		solve(m, r, coinidR, queryidR)
	}

	type pair struct {
		x, y Frac
	}
	type P struct {
		x int
		y pair
	}
	lines := make([]P, 0)
	coini := 0
	for _, q := range queryid {
		queryx := queries[q]
		for coini < len(coinidM) {
			coin := coins[coinidM[coini]]
			if queryx < coin.x {
				break
			}
			for len(lines) > 0 {
				p1 := lines[len(lines)-1].x
				if coins[p1].v >= coin.v {
					break
				}
				lines = lines[:len(lines)-1]
			}
			for 2 <= len(lines) {
				p1 := lines[len(lines)-1].x
				y23 := lines[len(lines)-1].y.y
				y12 := Frac{coin.x - coins[p1].x, coins[p1].v - coin.v}
				if y12.lessThan(y23) {
					break
				}
				lines = lines[:len(lines)-1]
			}
			if len(lines) == 0 {
				lines = append(lines, P{coinidM[coini], pair{Frac{1001001001, 1}, Frac{1001001001, 1}}})
			} else {
				p1 := lines[len(lines)-1].x
				x12 := Frac{coins[p1].v*coin.x - coin.v*coins[p1].x, coins[p1].v - coin.v}
				y12 := Frac{coin.x - coins[p1].x, coins[p1].v - coin.v}
				lines = append(lines, P{coinidM[coini], pair{x12, y12}})
			}
			coini++
		}
		if len(lines) == 0 {
			continue
		}
		for lines[len(lines)-1].y.x.lessThan(Frac{queryx, 1}) {
			lines = lines[:len(lines)-1]
		}
		coin := coins[lines[len(lines)-1].x]
		Ans[q] = math.Max(Ans[q], float64(coin.v)/float64(queryx-coin.x))
	}
}
