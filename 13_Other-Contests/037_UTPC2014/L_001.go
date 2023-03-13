package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

const inf = int(1e9)

var n = 30
var e [30]int
var col, ans int

func main() {
	in := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscan(in, &s)
	use := make([]int, n)
	g := make([][]int, n)
	for i := 0; i < len(s); i += 4 {
		c := s[i]
		d := s[i+2]
		var a int
		if c <= 'Z' {
			a = int(c - 'A')
		} else {
			a = int(c-'a') + 26
		}
		var b int
		if d <= 'Z' {
			b = int(d - 'A')
		} else {
			b = int(d-'a') + 26
		}
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
		use[a] = 1
		use[b] = 1
	}
	vis := make([]int, n)
	cmp := make([][]int, 0)
	for i := 0; i < n; i++ {
		if use[i] != 0 && vis[i] == 0 {
			v := make([]int, 0)
			DFs(g, i, &v, &vis)
			cmp = append(cmp, v)
		}
	}

	res := 1
	for _, v := range cmp {
		INIT(v, g)
		ans = 0
		col = 0
		res *= 2 * rec(0, 0).y
	}
	fmt.Println(res)
}

func DFs(g [][]int, c int, v, vis *[]int) {
	(*vis)[c] = 1
	*v = append(*v, c)
	for _, i := range g[c] {
		if (*vis)[i] == 0 {
			DFs(g, i, v, vis)
		}
	}
}

func INIT(v []int, g [][]int) {
	n = len(v)
	to := make([]int, 30)
	for i := 0; i < n; i++ {
		to[v[i]] = i
	}
	for i := 0; i < 30; i++ {
		e[i] = 0
	}
	for _, i := range v {
		for _, j := range g[i] {
			e[to[i]] |= 1 << to[j]
			e[to[j]] |= 1 << to[i]
		}
	}
}

func rec(c, now int) pi {
	ans = max(ans, now)
	if c == n {
		return pi{0, 1}
	}
	if now+(n-c)*4 < ans {
		return pi{-inf, 1}
	}

	res := pi{0, 0}
	for i := 0; i < 2; i++ {
		if c != 0 || i != 0 {
			tmp1 := col
			if ((col >> c) & 1) != 0 {
				tmp1 = ^col
			}
			cnt := bits.OnesCount(uint(tmp1 & e[c] & ((1 << c) - 1)))
			tmp := rec(c+1, now+cnt)
			if res.x < tmp.x+cnt {
				res = pi{tmp.x + cnt, tmp.y}
			} else if res.x == tmp.x+cnt {
				res.y += tmp.y
			}
			col ^= 1 << c
		}
	}
	return res
}

type pi struct {
	x, y int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
