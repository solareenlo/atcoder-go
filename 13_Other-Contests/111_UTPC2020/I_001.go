package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var f [1 << 19]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	type pair struct {
		x, y int
	}

	type tuple struct {
		x, y, z int
	}

	var n, m int
	fmt.Fscan(in, &n, &m)
	used := make([]bool, m)
	mst := make([]bool, m)
	a := make([]int, m)
	b := make([]int, m)
	c := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &a[i], &b[i], &c[i])
		a[i]--
		b[i]--
	}
	g := make([][]pair, n)
	G := make([][]pair, n)
	v := make([]int, m)
	for i := 0; i < m; i++ {
		v[i] = i
	}
	for i := 0; i < n; i++ {
		f[i] = i
	}
	sort.Slice(v, func(i, j int) bool {
		return c[v[i]] < c[v[j]]
	})
	for i := 0; i < m; i++ {
		j := v[i]
		G[a[j]] = append(G[a[j]], pair{b[j], j})
		G[b[j]] = append(G[b[j]], pair{a[j], j})
		y := find(a[j])
		z := find(b[j])
		if y != z {
			f[y] = z
			g[a[j]] = append(g[a[j]], pair{b[j], j})
			g[b[j]] = append(g[b[j]], pair{a[j], j})
			mst[j] = true
		}
	}
	id := make([]int, n)
	nid := 1
	var q int
	fmt.Fscan(in, &q)
	sum := 0
	seen := make([]int, 0)
	for q > 0 {
		q--
		var x int
		fmt.Fscan(in, &x)
		T := (sum ^ x)
		T--
		if used[T] {
			sum = 0
		} else {
			sum = T + 1
		}
		used[T] = true
		if !mst[T] {
			fmt.Fprintln(out, sum)
			continue
		}
		turn := 0
		v := make([]tuple, 0)
		vv := make([]tuple, 0)
		v = append(v, tuple{a[T], b[T], 0})
		vv = append(vv, tuple{b[T], a[T], 0})
		seen = []int{a[T], b[T]}
		for {
			for len(v) > 0 {
				ls := v[len(v)-1]
				ed := true
				s := ls.x
				p := ls.y
				id := ls.z
				for i := id; i < len(g[s]); i++ {
					nx := g[s][i].x
					eid := g[s][i].y
					if nx == p {
						continue
					}
					if used[eid] {
						g[s][i], g[s][len(g[s])-1] = g[s][len(g[s])-1], g[s][i]
						g[s] = g[s][:len(g[s])-1]
						i--
						continue
					}
					v = v[:len(v)-1]
					v = append(v, tuple{s, p, i + 1})
					v = append(v, tuple{nx, s, 0})
					seen = append(seen, nx)
					ed = false
					break
				}
				if ed {
					v = v[:len(v)-1]
				} else {
					break
				}
			}
			if len(v) == 0 {
				break
			}
			turn ^= 1
			v, vv = vv, v
		}
		for j := turn; j < len(seen); j += 2 {
			id[seen[j]] = nid
		}
		for k := turn; k < len(seen); k += 2 {
			s := seen[k]
			for i := 0; i < len(G[s]); i++ {
				j := G[s][i].y
				if used[j] {
					G[s][i], G[s][len(G[s])-1] = G[s][len(G[s])-1], G[s][i]
					G[s] = G[s][:len(G[s])-1]
					i--
					continue
				}
				if id[a[j]] != id[b[j]] {
					sum += j + 1
					used[j] = true
					G[s][i], G[s][len(G[s])-1] = G[s][len(G[s])-1], G[s][i]
					G[s] = G[s][:len(G[s])-1]
					i--
					continue
				}
			}
		}
		nid++
		fmt.Fprintln(out, sum)
	}
}

func find(x int) int {
	if x == f[x] {
		return x
	}
	f[x] = find(f[x])
	return f[x]
}
