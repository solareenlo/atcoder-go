package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)
	x := make([]int, N)
	y := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &x[i], &y[i])
	}

	if N == 1 {
		fmt.Println(0)
		return
	}

	a, b := 0, 0
	for i := 1; i < N; i++ {
		x[i] -= x[0]
		y[i] -= y[0]
		if i == 1 {
			a = x[i]
			b = y[i]
		} else {
			res := get(a, b, x[i], y[i])
			a = res.x
			b = res.y
		}
	}

	mx, Mx, my, My := 0, 0, 0, 0
	for i := 1; i < N; i++ {
		p := x[i]
		q := y[i]
		r := a
		s := b
		g := r*r + s*s
		u := (p*r + q*s) / g
		v := (q*r - p*s) / g
		chmin(&mx, u)
		chmax(&Mx, u)
		chmin(&my, v)
		chmax(&My, v)
	}

	d := max(Mx-mx+1, My-my+1)
	fmt.Println(d*d - N)
}

type pair struct {
	x, y int
}

func get(p, q, r, s int) pair {
	if r == 0 && s == 0 {
		return pair{p, q}
	}
	g := r*r + s*s
	np := near(p*r+q*s, g)
	nq := near(q*r-p*s, g)
	a := p - (np*r - nq*s)
	b := q - (nq*r + np*s)
	return get(r, s, a, b)
}

func near(x, d int) int {
	p := false
	if x >= 0 {
		p = true
	}
	x = abs(x)
	var t int
	if x%d*2 >= d {
		t = x/d + 1
	} else {
		t = x / d
	}
	if !p {
		t = -t
	}
	return t
}

func chmax(a *int, b int) bool {
	if *a < b {
		*a = b
		return true
	}
	return false
}

func chmin(a *int, b int) bool {
	if *a > b {
		*a = b
		return true
	}
	return false
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
