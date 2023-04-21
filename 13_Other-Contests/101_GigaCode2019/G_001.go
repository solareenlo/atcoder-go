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

	var n, m int
	fmt.Fscan(in, &n, &m)
	var a, b, deg [100010]int
	var g0 [100010][]int
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &a[i], &b[i])
		a[i]--
		b[i]--
		deg[a[i]]++
		deg[b[i]]++
		g0[a[i]] = append(g0[a[i]], b[i])
		g0[b[i]] = append(g0[b[i]], a[i])
	}
	var w [100010]int
	for i := range w {
		w[i] = 1
	}
	que := make([]int, 0)
	for i := 0; i < n; i++ {
		if deg[i] == 1 {
			que = append(que, i)
		}
	}
	ans := 0
	for len(que) > 0 {
		x := que[0]
		que = que[1:]
		deg[x]--
		for _, y := range g0[x] {
			if deg[y] >= 1 {
				deg[y]--
				if deg[y] == 1 {
					que = append(que, y)
				}
				w[y] += w[x]
				ans += w[x] * (n - w[x])
			}
		}
	}
	if m == n-1 {
		fmt.Println(ans)
		return
	}
	var g2 [100010][]int
	for i := 0; i < n; i++ {
		if deg[i] == 2 {
			for _, y := range g0[i] {
				if deg[y] > 0 {
					g2[i] = append(g2[i], y)
					g2[y] = append(g2[y], i)
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		sort.Ints(g2[i])
		g2[i] = unique(g2[i])
	}
	vs := make([][]int, 0)
	var used [100010]bool
	for i := 0; i < n; i++ {
		if deg[i] <= 2 {
			continue
		}
		for _, y := range g2[i] {
			if deg[y] == 2 && !used[y] {
				v := make([]int, 0)
				v = append(v, i)
				k := y
				pr := i
				for deg[k] == 2 {
					v = append(v, k)
					used[k] = true
					for _, z := range g2[k] {
						if pr != z && !used[z] {
							pr = k
							k = z
							break
						}
					}
				}
				v = append(v, k)
				vs = append(vs, v)
			}
		}
	}
	for i := 0; i < m; i++ {
		if deg[a[i]] > 2 && deg[b[i]] > 2 {
			v := make([]int, 2)
			v[0] = a[i]
			v[1] = b[i]
			vs = append(vs, v)
		}
	}
	if len(vs) == 0 {
		v := make([]int, 0)
		for i := 0; i < n; i++ {
			if deg[i] == 2 {
				k := i
				pr := -1
				for {
					v = append(v, k)
					for _, y := range g2[k] {
						if y != pr {
							pr = k
							k = y
							break
						}
					}
					if k == i {
						break
					}
				}
				v = append(v, i)
				vs = append(vs, v)
				break
			}
		}
	}
	num := make(map[int]int)
	var numr [2000]int
	for _, v := range vs {
		num[v[0]] = 0
		num[v[len(v)-1]] = 0
	}
	n1 := 0
	keys := make([]int, 0, len(num))
	for k := range num {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, p := range keys {
		num[p] = n1
		numr[n1] = p
		n1++
	}
	var g [2000][]P
	var d [2000][2000]int
	var a1, b1 [2500]int
	m1 := 0
	for _, v := range vs {
		x1 := num[v[0]]
		y1 := num[v[len(v)-1]]
		g[x1] = append(g[x1], P{y1, len(v) - 1})
		g[y1] = append(g[y1], P{x1, len(v) - 1})
		a1[m1] = x1
		b1[m1] = y1
		m1++
	}
	const INF = int(1e9) + 7
	for i := 0; i < n1; i++ {
		for j := range d[i] {
			d[i][j] = INF
		}
		d[i][i] = 0
		que := &HeapPair{}
		heap.Push(que, P{0, i})
		for que.Len() > 0 {
			p := heap.Pop(que).(P)
			x := p.y
			if p.x > d[i][x] {
				continue
			}
			for _, e := range g[x] {
				y := e.x
				if d[i][y] > d[i][x]+e.y {
					d[i][y] = d[i][x] + e.y
					heap.Push(que, P{d[i][y], y})
				}
			}
		}
	}
	for i := 0; i < n1; i++ {
		for j := 0; j < i; j++ {
			ans += w[numr[i]] * w[numr[j]] * d[i][j]
		}
	}
	ws := make([][]int, 0)
	ws2 := make([][]int, 0)
	for _, v := range vs {
		wv := make([]int, len(v)-1)
		wv2 := make([]int, len(v)-1)
		for i := 1; i < len(v)-1; i++ {
			wv[i] = wv[i-1] + w[v[i]]
			wv2[i] = wv2[i-1] + w[v[i]]*i
		}
		ws = append(ws, wv)
		ws2 = append(ws2, wv2)
	}
	for i := 0; i < n; i++ {
		used[i] = false
	}
	for k := 0; k < m1; k++ {
		v := vs[k]
		for i := 0; i < len(v); i++ {
			if used[v[i]] {
				continue
			}
			used[v[i]] = true
			myon := k
			if i == 0 || i == len(v)-1 {
				myon = m1
			}
			for j := 0; j < myon; j++ {
				d1 := min(i+d[a1[k]][a1[j]], len(v)-1-i+d[b1[k]][a1[j]])
				d2 := min(i+d[a1[k]][b1[j]], len(v)-1-i+d[b1[k]][b1[j]])
				if d1+len(vs[j])-1 <= d2 {
					ans += w[v[i]] * (ws2[j][len(ws2[j])-1] + d1*ws[j][len(ws[j])-1])
				} else if d2+len(vs[j])-1 <= d1 {
					ans += w[v[i]] * (ws[j][len(ws[j])-1]*(len(vs[j])-1+d2) - ws2[j][len(ws2[j])-1])
				} else {
					l := (len(vs[j]) - 1 + d2 - d1) / 2
					ans += w[v[i]] * (ws2[j][l] + d1*ws[j][l] + (len(vs[j])-1+d2)*(ws[j][len(ws[j])-1]-ws[j][l]) - ws2[j][len(ws2[j])-1] + ws2[j][l])
				}
			}
		}
	}
	for i := 0; i < m1; i++ {
		v := vs[i]
		if len(v) >= 4 {
			wv := make([]int, 0)
			for j := 1; j < len(v)-1; j++ {
				wv = append(wv, w[v[j]])
			}
			ans += cycle(wv, d[a1[i]][b1[i]]+2)
		}
	}
	fmt.Println(ans)
}

func cycle(v []int, d int) int {
	n1 := len(v)
	s := make([]int, n1)
	ss := make([]int, n1)
	s[0] = v[0]
	for i := 1; i < n1; i++ {
		s[i] = s[i-1] + v[i]
		ss[i] = ss[i-1] + v[i]*i
	}
	ret := 0
	for i := 0; i < n1; i++ {
		k := min((n1-1+2*i+d)/2, n1-1)
		ret += v[i] * (ss[k] - ss[i] - i*(s[k]-s[i]))
		ret += v[i] * (-ss[n1-1] + ss[k] + (n1-1+d+i)*(s[n1-1]-s[k]))
	}
	return ret
}

type P struct {
	x, y int
}

type HeapPair []P

func (h HeapPair) Len() int            { return len(h) }
func (h HeapPair) Less(i, j int) bool  { return h[i].x < h[j].x }
func (h HeapPair) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *HeapPair) Push(x interface{}) { *h = append(*h, x.(P)) }

func (h *HeapPair) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
