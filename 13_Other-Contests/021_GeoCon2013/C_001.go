package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const EPS = 1e-7

var T []bool
var X, Y []int
var Xs []int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int
	fmt.Fscan(in, &N)
	T = make([]bool, N)
	X = make([]int, N)
	Y = make([]int, N)
	Xs = make([]int, N)
	for i := 0; i < N; i++ {
		var S string
		var x, y int
		fmt.Fscan(in, &S, &x, &y)
		if S[0] == 'T' {
			T[i] = false
		} else {
			T[i] = true
		}
		X[i] = x
		Y[i] = y
		Xs[i] = x
	}
	sort.Ints(Xs)
	tmp := unique(Xs)
	Xl := len(tmp)
	upp := newConvex(Xl, Xs, 1)
	low := newConvex(Xl, Xs, -1)
	for i := 0; i < N; i++ {
		if T[i] {
			upp.update(X[i], Y[i])
			low.update(X[i], Y[i])
		} else {
			upv := upp.at(X[i])
			lwv := low.at(X[i])
			if upv > 0 {
				if lwv-EPS < float64(Y[i]) && float64(Y[i]) < upv+EPS {
					fmt.Fprintln(out, "DANGER")
					continue
				}
			}
			fmt.Fprintln(out, "SAFE")
		}
	}
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
	sort.Ints(result)
	n := len(result)
	for i := 0; i < n; i++ {
		a[i] = result[i]
	}
	return result
}

const BIT = 1 << 18

type convex struct {
	N    int
	x, y []float64
	sum  []int
	sgn  float64
}

func newConvex(N int, Xs []int, sg float64) *convex {
	res := new(convex)
	res.N = N
	res.x = make([]float64, N)
	res.y = make([]float64, N)
	for i := 0; i < N; i++ {
		res.x[i] = float64(Xs[i])
		res.y[i] = -1
	}
	res.sum = make([]int, 1<<19)
	for i := 0; i < 2*BIT; i++ {
		res.sum[i] = 0
	}
	res.sgn = sg
	return res
}

func (con *convex) update(xv, yv int) {
	a := con.at(xv)
	ps := lowerBound(con.x, float64(xv))
	if a < 0 {
		con.set(ps, 1)
		con.y[ps] = float64(yv)
	} else {
		if (float64(yv)-a)*con.sgn > 0 {
			con.set(ps, 1)
			con.y[ps] = float64(yv)
		} else {
			return
		}
	}
	p1 := con.right(ps)
	p2 := con.right(p1)
	for p2 != -1 {
		nv := con.y[p1] + (con.y[p2]-con.y[p1])*(con.x[ps]-con.x[p1])/(con.x[p2]-con.x[p1])
		if (float64(yv)-nv)*con.sgn < EPS {
			break
		}
		con.set(p1, 0)
		con.y[p1] = -1
		tmp := con.right(p2)
		p1 = p2
		p2 = tmp
	}
	p1 = con.left(ps)
	p2 = con.left(p1)
	for p2 != -1 {
		nv := con.y[p1] + (con.y[p2]-con.y[p1])*(con.x[ps]-con.x[p1])/(con.x[p2]-con.x[p1])
		if (float64(yv)-nv)*con.sgn < EPS {
			break
		}
		con.set(p1, 0)
		con.y[p1] = -1
		tmp := con.left(p2)
		p1 = p2
		p2 = tmp
	}
}

func (con convex) at(xv int) float64 {
	ps := lowerBound(con.x, float64(xv))
	if con.query(ps+1) == 0 {
		return -1
	}
	if con.y[ps] != -1 {
		return float64(con.y[ps])
	}
	lf := con.left(ps)
	rg := con.right(ps)
	if lf == -1 || rg == -1 {
		return -1
	}
	return con.y[lf] + (con.y[rg]-con.y[lf])*(con.x[ps]-con.x[lf])/(con.x[rg]-con.x[lf])
}

func lowerBound(a []float64, x float64) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func (con convex) query(r int) int {
	r += BIT
	ret := 0
	for r != 0 {
		if (r & 1) != 0 {
			r--
			ret += con.sum[r]
		}
		r >>= 1
	}
	return ret
}

func (con convex) left(p int) int {
	if p <= 0 {
		return -1
	}
	bas := con.query(p)
	if bas == 0 {
		return -1
	}
	ret := 1
	for ret < BIT {
		if bas > con.sum[ret*2] {
			bas -= con.sum[ret*2]
			ret = ret*2 + 1
		} else {
			ret = ret * 2
		}
	}
	return ret - BIT
}

func (con convex) right(p int) int {
	if p == -1 {
		return -1
	}
	bas := con.query(p+1) + 1
	ret := 1
	for ret < BIT {
		if bas > con.sum[ret*2] {
			bas -= con.sum[ret*2]
			ret = ret*2 + 1
		} else {
			ret = ret * 2
		}
	}
	if ret == 2*BIT-1 {
		return -1
	}
	return ret - BIT
}

func (con *convex) set(p, v int) {
	p += BIT
	con.sum[p] = v
	p >>= 1
	for p != 0 {
		con.sum[p] = con.sum[p*2] + con.sum[p*2+1]
		p >>= 1
	}
}
