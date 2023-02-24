package main

import (
	"bufio"
	"fmt"
	"os"
)

const EPS float64 = 1e-10

var ps [][]Pt

func main() {
	in := bufio.NewReader(os.Stdin)

	var TC int
	fmt.Fscan(in, &TC)
	P := make([]Pt, 3030)
	sigs := make([]int, 3030)
	ps = make([][]Pt, 4)
	for i := range ps {
		ps[i] = make([]Pt, 20)
	}
	for TC > 0 {
		TC--
		var N int
		var H float64
		fmt.Fscan(in, &N, &H)
		for i := 0; i < N; i++ {
			fmt.Fscan(in, &P[i].x, &P[i].y)
		}
		P[N] = P[0]
		area := 0.0
		for i := 0; i < N; i++ {
			area += P[i].det(P[i+1])
		}
		if area < 0 {
			tmp := P[1:N]
			tmp = reverseOrder(tmp)
			for i := 0; i < N-1; i++ {
				P[i+1] = tmp[i]
			}
			area *= -1.0
		}
		for i := 1; i < N-1; i++ {
			sigs[i] = sig(tri(P[0], P[i], P[i+1]))
		}
		ans := area
		for i := 1; i < N-1; i++ {
			for j := 1; j < N-1; j++ {
				res := solve(
					sigs[i],
					sigs[j],
					P[0].plus(Pt{H, H}), P[i].plus(Pt{H, H}), P[i+1].plus(Pt{H, H}),
					P[0].minus(Pt{H, H}), P[j].minus(Pt{H, H}), P[j+1].minus(Pt{H, H}))
				ans -= float64(sigs[i]) * float64(sigs[j]) * res
			}
		}
		fmt.Printf("%.2f\n", ans-EPS)
	}
}

func reverseOrder(a []Pt) []Pt {
	n := len(a)
	res := make([]Pt, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func solve(sa, sb int, a0, a1, a2, b0, b1, b2 Pt) float64 {
	if sa < 0 {
		a1, a2 = a2, a1
	}
	if sb < 0 {
		b1, b2 = b2, b1
	}
	n := 3
	ps[0][0] = a0
	ps[0][1] = a1
	ps[0][2] = a2
	ps[0][3] = a0
	if n >= 3 {
		n = convexCut(n, ps[0], b0, b1, ps[1])
	}
	if n >= 3 {
		n = convexCut(n, ps[1], b1, b2, ps[2])
	}
	if n >= 3 {
		n = convexCut(n, ps[2], b2, b0, ps[3])
	}
	ret := 0.0
	if n >= 3 {
		for i := 0; i < n; i++ {
			ret += ps[3][i].det(ps[3][i+1])
		}
	}
	ret /= 2
	return ret
}

func convexCut(n int, p []Pt, a, b Pt, q []Pt) int {
	m := 0
	p[n] = p[0]
	for i := 0; i < n; i++ {
		if sig(tri(a, b, p[i])) >= 0 {
			q[m] = p[i]
			m++
		}
		if sig(tri(a, b, p[i]))*sig(tri(a, b, p[i+1])) < 0 {
			q[m] = pLL(a, b, p[i], p[i+1])
			m++
		}
	}
	q[m] = q[0]
	return m
}

func pLL(a, b, c, d Pt) Pt {
	b = b.minus(a)
	d = d.minus(c)
	return a.plus(((b.mul((c.minus(a)).det(d))).div(b.det(d))))
}

type Pt struct {
	x, y float64
}

func (b Pt) plus(a Pt) Pt     { return Pt{b.x + a.x, b.y + a.y} }
func (b Pt) minus(a Pt) Pt    { return Pt{b.x - a.x, b.y - a.y} }
func (b Pt) mul(k float64) Pt { return Pt{b.x * k, b.y * k} }
func (b Pt) div(k float64) Pt { return Pt{b.x / k, b.y / k} }
func (b Pt) det(a Pt) float64 { return b.x*a.y - b.y*a.x }
func tri(a, b, c Pt) float64  { return (b.minus(a)).det(c.minus(a)) }

func sig(r float64) int {
	if r < -EPS {
		return -1
	}
	if r > +EPS {
		return 1
	}
	return 0
}
