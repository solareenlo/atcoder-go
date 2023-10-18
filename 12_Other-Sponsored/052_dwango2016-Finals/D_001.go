package main

import (
	"bufio"
	"fmt"
	"os"
)

const SIZE = 300005

var n int
var vec [SIZE][]int
var nd, dep, mx [SIZE]int
var back, par, up, down [SIZE]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &n)
	for i := 0; i < n-1; i++ {
		var p int
		fmt.Fscan(in, &p)
		p--
		vec[p] = append(vec[p], i+1)
	}
	DFS(0, -1, 0)
	MAKE(0, -1)
	solve1(0, -1)
	solve2(0, -1)
	for i := 1; i < n; i++ {
		if i != 1 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, down[i]+up[i]+nd[i]*(n-nd[i]))
	}
	fmt.Fprintln(out)
}

func DFS(v, p, d int) {
	nd[v] = 1
	dep[v] = d
	mx[v] = 0
	back[v] = 0
	for i := 0; i < len(vec[v]); i++ {
		to := vec[v][i]
		if to != p {
			DFS(to, v, d+1)
			nd[v] += nd[to]
			mx[v] = max(mx[v], mx[to]+1)
		}
	}
}

func MAKE(v, p int) {
	m1 := back[v]
	m2 := 0
	for i := 0; i < len(vec[v]); i++ {
		to := vec[v][i]
		if to != p {
			vl := mx[to] + 1
			if m1 < vl {
				m1, vl = vl, m1
			}
			if m2 < vl {
				m2, vl = vl, m2
			}
		}
	}
	for i := 0; i < len(vec[v]); i++ {
		to := vec[v][i]
		if to != p {
			vl := mx[to] + 1
			zan := m1
			if m1 == vl {
				zan = m2
			}
			back[to] = zan + 1
			MAKE(to, v)
		}
	}
}

func solve1(v, p int) {
	m1 := back[v]
	m2 := 0
	for i := 0; i < len(vec[v]); i++ {
		to := vec[v][i]
		if to != p {
			solve1(to, v)
			vl := mx[to] + 1
			if m1 < vl {
				m1, vl = vl, m1
			}
			if m2 < vl {
				m2, vl = vl, m2
			}
		}
	}
	s1, s2 := 0, 0
	for i := 0; i < len(vec[v]); i++ {
		to := vec[v][i]
		if to != p {
			vl := mx[to] + 1
			zan := m1
			if vl == m1 {
				zan = m2
			}
			if m1 == zan {
				s1 += down[to] + nd[to]
			} else if m2 == zan {
				s2 += down[to] + nd[to]
			}
		}
	}
	for i := 0; i < len(vec[v]); i++ {
		to := vec[v][i]
		if to != p {
			vl := mx[to] + 1
			zan := m1
			if vl == m1 {
				zan = m2
			}
			if zan < back[v]+1 {
				down[v] += down[to] + nd[to]
			}
			if zan == m1 {
				s1 -= down[to] + nd[to]
			} else if zan == m2 {
				s2 -= down[to] + nd[to]
			}
			if vl+1 > m1 {
				up[to] += s1
			}
			if vl+1 > m2 {
				up[to] += s2
			}
			if zan == m1 {
				s1 += down[to] + nd[to]
			} else if zan == m2 {
				s2 += down[to] + nd[to]
			}
		}
	}
}

func solve2(v, p int) {
	for i := 0; i < len(vec[v]); i++ {
		to := vec[v][i]
		if to != p {
			if mx[to]+2 > mx[v] {
				up[to] += up[v] + n - nd[v]
			}
			solve2(to, v)
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
