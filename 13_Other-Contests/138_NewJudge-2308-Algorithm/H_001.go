package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"sort"
)

const mod = 998244353

type DSU struct {
	p []int
}

func (d *DSU) init(n int) {
	d.p = make([]int, n)
	for i := range d.p {
		d.p[i] = i
	}
}

func (d DSU) find(x int) int {
	if x == d.p[x] {
		return x
	}
	d.p[x] = d.find(d.p[x])
	return d.p[x]
}

func (d DSU) same(x, y int) bool {
	return d.find(x) == d.find(y)
}

func (d *DSU) merge(x, y int) int {
	if d.same(x, y) {
		return 0
	}
	d.p[d.p[x]] = d.p[y]
	return 1
}

func work(n, k, S int, e *[]pair) int {
	comp := popcount(S)
	if comp == 1 {
		return k % mod
	}
	for i := range *e {
		if (*e)[i].x > (*e)[i].y {
			(*e)[i].x, (*e)[i].y = (*e)[i].y, (*e)[i].x
		}
	}
	sortPair(*e)
	*e = unique(*e)
	var d DSU
	d.init(n)
	for _, tmp := range *e {
		u, v := tmp.x, tmp.y
		comp -= d.merge(u, v)
	}
	if comp > 1 {
		res := 1
		for i := 0; i < n; i++ {
			if ((S>>i)&1) != 0 && i == d.find(i) {
				ne := make([]pair, 0)
				for _, tmp := range *e {
					u, v := tmp.x, tmp.y
					if d.find(u) == i {
						ne = append(ne, pair{u, v})
					}
				}
				nS := 0
				for j := 0; j < n; j++ {
					if i == d.find(j) {
						nS |= 1 << j
					}
				}
				res = res * work(n, k, nS, &ne) % mod
			}
		}
		return res
	}
	deg := make([]int, n)
	for _, tmp := range *e {
		u, v := tmp.x, tmp.y
		deg[u]++
		deg[v]++
	}
	t := -1
	for i := 0; i < n; i++ {
		if (S>>i)&1 != 0 && (t == -1 || deg[i] < deg[t]) {
			t = i
		}
	}
	adj := make([]int, 0)
	e0 := make([]pair, 0)
	for _, tmp := range *e {
		u, v := tmp.x, tmp.y
		if u == t || v == t {
			adj = append(adj, t^u^v)
		} else {
			e0 = append(e0, pair{u, v})
		}
	}
	res := 0
	oth := work(n, k, S^(1<<t), &e0)
	for T := 0; T < 1<<len(adj); T++ {
		if T == 0 {
			res = (res + oth*k%mod) % mod
		} else if popcount(T) == 1 {
			res = (res + mod - oth) % mod
		} else {
			low := adj[ctz(T)]
			ne := make([]pair, len(e0))
			copy(ne, e0)
			merge := 0
			nS := S ^ (1 << t) ^ (1 << low)
			ok := true
			for i := 0; i < len(adj); i++ {
				if (T >> i & 1) != 0 {
					merge |= 1 << adj[i]
					nS ^= 1 << adj[i]
				}
			}
			for i := range ne {
				if ((merge >> ne[i].x) & 1) != 0 {
					ne[i].x = low
				}
				if ((merge >> ne[i].y) & 1) != 0 {
					ne[i].y = low
				}
				if ne[i].x == ne[i].y {
					ok = false
					break
				}
			}
			if !ok {
				continue
			}
			if parity(T) != 0 {
				res = (res + mod - work(n, k, nS, &ne)) % mod
			} else {
				res = (res + work(n, k, nS, &ne)) % mod
			}
		}
	}
	return res
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	e := make([]pair, m)
	for i := range e {
		fmt.Fscan(in, &e[i].x, &e[i].y)
		e[i].x--
		e[i].y--
	}
	fmt.Println(work(n, k, (1<<n)-1, &e))
}

func popcount(n int) int {
	return bits.OnesCount32(uint32(n))
}

func parity(x int) int {
	return popcount(x) % 2
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

func unique(a []pair) []pair {
	occurred := map[pair]bool{}
	result := []pair{}
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

func ctz(x int) int {
	return bits.TrailingZeros32(uint32(x))
}
