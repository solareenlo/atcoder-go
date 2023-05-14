package main

import (
	"bufio"
	"fmt"
	"os"
)

const inf = 0x20202020

func main() {
	in := bufio.NewReader(os.Stdin)

	var T int
	fmt.Fscan(in, &T)
	for T > 0 {
		T--
		var h, w int
		fmt.Fscan(in, &h, &w)
		var s [210]string
		for i := 0; i < h; i++ {
			fmt.Fscan(in, &s[i])
		}
		var flow Flow
		flow.init(2*h*w, 2*h*w+1, 2*h*w+2)
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				if s[i][j] != '#' {
					if s[i][j] == 'S' {
						flow.add(2*h*w, i*w+j, 10, 0)
					}
					if s[i][j] == 'G' {
						flow.add(i*w+j+h*w, 2*h*w+1, 10, 0)
					}
					if s[i][j] == '.' {
						flow.add(i*w+j, i*w+j+h*w, 1, 0)
					} else {
						flow.add(i*w+j, i*w+j+h*w, 10, 0)
					}
					for dx := -1; dx < 2; dx++ {
						for dy := -1; dy < 2; dy++ {
							if dx*dy == 0 {
								if dx == 0 && dy == 0 {
									continue
								}
								px := i + dx
								py := j + dy
								if px >= 0 && px < h && py >= 0 && py < w && s[px][py] != '#' {
									flow.add(i*w+j+h*w, px*w+py, 10, 0)
								}
							}
						}
					}
				}
			}
		}
		f := flow.sap()
		if f == 1 {
			fmt.Println("Alice")
		} else {
			cc := 0
			for i := 0; i < h; i++ {
				for j := 0; j < w; j++ {
					if s[i][j] == '.' {
						cc++
					}
				}
			}
			if cc%2 != 0 {
				fmt.Println("Alice")
			} else {
				fmt.Println("Bob")
			}
		}
	}
}

const M = 1000000
const N = 21000

type Flow struct {
	y, nxt, f                [M]int
	gap, fst, c, pre, q, dis [N]int
	S, T, tot, Tn            int
}

func (flow *Flow) init(s, t, tn int) {
	flow.tot = 1
	for i := 0; i < tn; i++ {
		flow.fst[i] = 0
	}
	flow.S = s
	flow.T = t
	flow.Tn = tn
}

func (flow *Flow) add(u, v, c1, c2 int) {
	flow.tot++
	flow.y[flow.tot] = v
	flow.f[flow.tot] = c1
	flow.nxt[flow.tot] = flow.fst[u]
	flow.fst[u] = flow.tot
	flow.tot++
	flow.y[flow.tot] = u
	flow.f[flow.tot] = c2
	flow.nxt[flow.tot] = flow.fst[v]
	flow.fst[v] = flow.tot
}

func (f *Flow) sap() int {
	u := f.S
	t := 1
	flow := 0
	for i := 0; i < f.Tn; i++ {
		f.c[i] = f.fst[i]
		f.dis[i] = f.Tn
		f.gap[i] = 0
	}
	f.q[0] = f.T
	f.dis[f.T] = 0
	f.pre[f.S] = 0
	for i := 0; i < t; i++ {
		u := f.q[i]
		for j := f.fst[u]; j > 0; j = f.nxt[j] {
			if f.dis[f.y[j]] > f.dis[u]+1 && f.f[j^1] != 0 {
				f.q[t] = f.y[j]
				f.dis[f.y[j]] = f.dis[u] + 1
				t++
			}
		}
	}
	for i := 0; i < f.Tn; i++ {
		f.gap[f.dis[i]]++
	}
	for f.dis[f.S] <= f.Tn {
		for f.c[u] != 0 && (f.f[f.c[u]] == 0 || f.dis[f.y[f.c[u]]]+1 != f.dis[u]) {
			f.c[u] = f.nxt[f.c[u]]
		}
		if f.c[u] != 0 {
			f.pre[f.y[f.c[u]]] = f.c[u] ^ 1
			u = f.y[f.c[u]]
			if u == f.T {
				minf := inf
				for p := f.pre[f.T]; p > 0; p = f.pre[f.y[p]] {
					minf = min(minf, f.f[p^1])
				}
				for p := f.pre[f.T]; p > 0; p = f.pre[f.y[p]] {
					f.f[p^1] -= minf
					f.f[p] += minf
				}
				flow += minf
				u = f.S
			}
		} else {
			f.gap[f.dis[u]]--
			if f.gap[f.dis[u]] == 0 {
				break
			}
			mind := f.Tn
			f.c[u] = f.fst[u]
			for j := f.fst[u]; j > 0; j = f.nxt[j] {
				if f.f[j] != 0 && f.dis[f.y[j]] < mind {
					mind = f.dis[f.y[j]]
					f.c[u] = j
				}
			}
			f.dis[u] = mind + 1
			f.gap[f.dis[u]]++
			if u != f.S {
				u = f.y[f.pre[u]]
			}
		}
	}
	return flow
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
