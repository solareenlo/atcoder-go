package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type pair struct {
	x, y int
}

func inc(a, b, c int) bool {
	return a <= b && b <= c
}

func muri() {
	fmt.Println(0)
	os.Exit(0)
}

type doubling struct {
	g            [][]int
	n, h, cnt    int
	par          [][]int
	dep, in, out []int
}

func (dou *doubling) dfs(v, p, d int) {
	dou.par[0][v] = p
	dou.dep[v] = d
	dou.in[v] = dou.cnt
	dou.cnt++
	for _, e := range dou.g[v] {
		if e != p {
			dou.dfs(e, v, d+1)
		}
	}
	dou.out[v] = dou.cnt
}

func (dou *doubling) init(gg [][]int, r int) {
	dou.g = gg
	dou.n = len(dou.g)
	dou.h = log2(dou.n) + 1
	dou.cnt = 0
	dou.par = make([][]int, dou.h)
	for i := range dou.par {
		dou.par[i] = make([]int, dou.n)
		for j := range dou.par[i] {
			dou.par[i][j] = -1
		}
	}
	dou.dep = make([]int, dou.n)
	dou.in = make([]int, dou.n)
	dou.out = make([]int, dou.n)

	dou.dfs(r, -1, 0)
	for i := 1; i < dou.h; i++ {
		for j := 0; j < dou.n; j++ {
			if dou.par[i-1][j] != -1 {
				dou.par[i][j] = dou.par[i-1][dou.par[i-1][j]]
			}
		}
	}
}

func (dou *doubling) lca(a, b int) int {
	if dou.dep[a] > dou.dep[b] {
		a, b = b, a
	}
	w := dou.dep[b] - dou.dep[a]
	for i := 0; i < dou.h; i++ {
		if (w & (1 << i)) != 0 {
			b = dou.par[i][b]
		}
	}
	if a == b {
		return a
	}
	for i := dou.h - 1; i >= 0; i-- {
		x := dou.par[i][a]
		y := dou.par[i][b]
		if x != y {
			a, b = x, y
		}
	}
	return dou.par[0][a]
}

func (dou *doubling) Len(a, b int) int {
	return dou.dep[a] + dou.dep[b] - dou.dep[dou.lca(a, b)]*2
}

func (dou *doubling) jump(a, b, d int) int {
	c := dou.lca(a, b)
	w := dou.dep[a] + dou.dep[b] - dou.dep[c]*2
	if d <= dou.dep[a]-dou.dep[c] {
		for i := 0; i < dou.h; i++ {
			if (d & (1 << i)) != 0 {
				a = dou.par[i][a]
			}
		}
		return a
	} else {
		d = w - d
		for i := 0; i < dou.h; i++ {
			if (d & (1 << i)) != 0 {
				b = dou.par[i][b]
			}
		}
		return b
	}
}

func getvals(n int, es []pair, rs [3]int) []int {
	t := make([][]int, n)
	for i := 0; i < n-1; i++ {
		t[es[i].x] = append(t[es[i].x], es[i].y)
		t[es[i].y] = append(t[es[i].y], es[i].x)
	}
	var getdist func(int) []int
	getdist = func(r int) []int {
		dist := make([]int, n)
		for i := range dist {
			dist[i] = -1
		}
		q := make([]int, 0)
		var reach func(int, int)
		reach = func(v, d int) {
			if dist[v] == -1 {
				dist[v] = d
				q = append(q, v)
			}
		}
		reach(r, 0)
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			for _, to := range t[v] {
				reach(to, dist[v]+1)
			}
		}
		return dist
	}
	as := make([][]int, n)
	for i := range as {
		as[i] = make([]int, 3)
	}
	for k := 0; k < 3; k++ {
		d := getdist(rs[k])
		for i := 0; i < n; i++ {
			as[i][k] = d[i]
		}
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		sort.Ints(as[i])
		res[i] = as[i][1]
	}
	return res
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	es := make([]pair, 0)
	for i := 0; i < n-1; i++ {
		var tmp int
		fmt.Fscan(in, &tmp)
		es = append(es, pair{i + 1, tmp - 1})
	}

	val := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &val[i])
	}

	valraw := make([]int, len(val))
	copy(valraw, val)
	resize(&val, n+n-1)
	t := make([][]int, n+n-1)
	var ae func(int, int)
	ae = func(a, b int) {
		t[a] = append(t[a], b)
		t[b] = append(t[b], a)
	}
	waf := make([]int, 0)
	for i := 0; i < n-1; i++ {
		a, b := es[i].x, es[i].y
		d := abs(val[a] - val[b])
		if d > 1 {
			muri()
		}
		if d == 0 {
			waf = append(waf, i)
		} else {
			val[n+i] = val[a] + val[b]
		}
		ae(a, n+i)
		ae(b, n+i)
	}
	for i := 0; i < n; i++ {
		val[i] *= 2
	}

	if len(waf) > 5 {
		muri()
	}

	n = n*2 - 1

	ans := 0
	var ysp doubling
	ysp.init(t, 0)
	for bit := 0; bit < 1<<len(waf); bit++ {
		for i := 0; i < len(waf); i++ {
			x := waf[i]
			a := es[x].x
			if (bit & (1 << i)) != 0 {
				val[(n+1)/2+x] = val[a] - 1
			} else {
				val[(n+1)/2+x] = val[a] + 1
			}
		}
		pa, pb := -1, -1
		ng := false
		for i := 0; i < n; i++ {
			mn := true
			for _, j := range t[i] {
				if val[j] < val[i] {
					mn = false
				}
			}
			if mn {
				if pa == -1 {
					pa = i
				} else if pb == -1 {
					pb = i
				} else {
					ng = true
				}
			}
		}
		if !ng {
			var tar int
			var dfs func(int, int, int) pair
			dfs = func(v, p, d int) pair {
				a, b := 0, 0
				if v <= n/2 {
					if tar == d {
						a++
					} else if tar < d {
						b++
					}
				}
				for _, to := range t[v] {
					if to != p {
						tmp := dfs(to, v, d+1)
						x, y := tmp.x, tmp.y
						a += x
						b += y
					}
				}
				return pair{a, b}
			}
			if pb == -1 {
				tar = val[pa]
				var dp [4][2]int
				dp[0][0] = 1
				for _, v := range t[pa] {
					tmp := dfs(v, pa, 1)
					a, b := tmp.x, tmp.y
					for x := 4 - 1; x >= 0; x-- {
						for y := 2 - 1; y >= 0; y-- {
							if x+1 < 4 {
								dp[x+1][y] += dp[x][y] * a
							}
							if y+1 < 2 {
								dp[x][y+1] += dp[x][y] * b
							}
						}
					}
				}
				ans += dp[3][0]
				ans += dp[2][1]
			} else {
				Len := ysp.Len(pa, pb)
				sum := val[pa] + val[pb]
				if sum >= Len && (sum-Len)%2 == 0 {
					z := (sum - Len) / 2
					k := val[pa] - z
					if inc(1, k, Len-1) {
						w := 1
						for i := 0; i < 2; i++ {
							tar = val[pa]
							x := dfs(pa, ysp.jump(pa, pb, 1), 0).x
							w *= x
							pa, pb = pb, pa
						}
						pc := ysp.jump(pa, pb, k)
						nga := ysp.jump(pc, pa, 1)
						ngb := ysp.jump(pc, pb, 1)
						x := 0
						tar = z
						if tar == 0 {
							x = 1
						} else {
							for _, to := range t[pc] {
								if to != nga && to != ngb {
									x += dfs(to, pc, 1).x
								}
							}
						}
						w *= x
						ans += w
					}
				}
			}
		}
	}
	fmt.Println(ans)
}

func resize(a *[]int, n int) {
	if len(*a) > n {
		*a = (*a)[:n]
	} else {
		n = n - len(*a)
		for i := 0; i < n; i++ {
			*a = append(*a, 0)
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func log2(n int) int {
	var k int
	for k = 0; n != 0; n >>= 1 {
		k++
	}
	return k - 1
}
