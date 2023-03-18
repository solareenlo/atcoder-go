package main

import (
	"bufio"
	"fmt"
	"os"
)

type e struct {
	a int
	b int
	c float64
}

var (
	n, m, k int
	s, t    []int
	val     []float64
	z       []float64
	d       []float64
	g       []e
)

func ch(x float64) bool {
	g = g[:0]
	for i := 0; i < m; i++ {
		a := s[i]
		b := t[i]
		if val[a] != 1e9 && val[b] != 1e9 {
			if val[a]+z[i]-val[b] > x {
				return false
			}
		} else if val[a] == 1e9 && val[b] == 1e9 {
			g = append(g, e{b, a, x - z[i]})
		} else if val[a] == 1e9 {
			g = append(g, e{n, a, val[b] + x - z[i]})
		} else {
			g = append(g, e{b, n, x - val[a] - z[i]})
		}
	}
	for i := 0; i <= n; i++ {
		d[i] = 1e9
	}
	d[0] = 0
	c := 0
	for {
		o := false
		for i := 0; i < len(g); i++ {
			v := g[i].a
			u := g[i].b
			w := g[i].c
			if d[u] > d[v]+w {
				o = true
				d[u] = d[v] + w
			}
		}
		c++
		if o == false {
			return true
		}
		if c == n+1 {
			return false
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	fmt.Fscan(in, &m)
	fmt.Fscan(in, &k)
	s = make([]int, m)
	t = make([]int, m)
	val = make([]float64, n+1)
	z = make([]float64, m)
	d = make([]float64, n+1)
	for i := 0; i <= n; i++ {
		val[i] = 1e9
	}
	for i := 0; i < k; i++ {
		var x int
		var y float64
		fmt.Fscan(in, &x)
		fmt.Fscan(in, &y)
		x--
		val[x] = y
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &s[i])
		fmt.Fscan(in, &t[i])
		fmt.Fscan(in, &z[i])
		s[i]--
		t[i]--
	}
	if ch(-1e9) {
		fmt.Println("#")
		return
	}
	l := -1e9
	r := 1e9
	for r-l > 0.000001 {
		o := (l + r) / 2
		if ch(o) {
			r = o
		} else {
			l = o
		}
	}
	fmt.Printf("%.8f\n", r)
}
