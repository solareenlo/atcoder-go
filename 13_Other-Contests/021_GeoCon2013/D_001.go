package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const eps = 1e-7

var N, CP int
var AA [13]derKreis
var S, TT derPunkt
var P [111111]derPunkt
var W [13]float64
var V float64

func main() {
	in := bufio.NewReader(os.Stdin)

	var sx, sy, gx, gy float64
	fmt.Fscan(in, &sx, &sy, &gx, &gy, &V, &N)
	S = derPunkt{sx, sy}
	TT = derPunkt{gx, gy}
	for i := 1; i <= N; i++ {
		var wx, wy float64
		fmt.Fscan(in, &wx, &wy, &AA[i].r, &W[i])
		AA[i].O = derPunkt{wx, wy}
	}
	if N == 1 && W[1] > 0 {
		uno()
	} else {
		dos()
	}
}

func uno() {
	var a, b, c, x0, y0, r0, w, r1, t1, d1, x2, y2, r2, t2, d2, t3, r3, k float64
	w = W[1]
	r0 = AA[1].r
	S = S.minus(AA[1].O)
	TT = TT.minus(AA[1].O)
	x0 = S.x
	y0 = S.y
	a = V*V - w*w
	vx := V * math.Cos((TT.minus(S)).theta())
	vy := V * math.Sin((TT.minus(S)).theta())
	b = 2 * (x0*vx + y0*vy - r0*w)
	c = x0*x0 + y0*y0 - r0*r0
	s1 := (-b + math.Sqrt(b*b-4*a*c)) / (2 * a)
	s2 := (-b - math.Sqrt(b*b-4*a*c)) / (2 * a)
	if fle(b*b-4*a*c, 0) || (fle(s1, 0) && fle(s1, 0)) || (fle(Dis(S, TT)/V, s1) && fle(Dis(S, TT)/V, s2)) {
		fmt.Println(Dis(S, TT) / V)
		return
	}
	t1 = math.Sqrt((S.sqrlen() - r0*r0) / (V*V - w*w))
	r3 = math.Sqrt(TT.sqrlen())
	k = math.Sqrt(V*V - w*w)
	r1 = r0 + t1*w
	gamma := math.Acos(w / V)
	d1 = S.theta() - math.Asin(math.Sin(gamma)*V*t1/math.Sqrt(S.sqrlen()))
	d3 := TT.theta()
	if d3 > d1 {
		d3 -= 2 * math.Pi
	}
	l := 0.0
	r := d1 - d3
	res := 1e18
	if !fls(k/w*math.Log(r3/r1), r) {
		for r-l > eps {
			mid := (l + r) / 2
			d2 = d1 - mid
			r2 = math.Cos(d2-d3)*r3 - math.Sin(d2-d3)*r3*w/k
			if k/w*math.Log(r2/r1) < mid {
				l = mid
			} else {
				r = mid
			}
		}
		d2 = d1 - l
		r2 = math.Cos(d2-d3)*r3 - math.Sin(d2-d3)*r3*w/k
		t2 = (r2 - r0) / w
		x2 = math.Cos(d2) * r2
		y2 = math.Sin(d2) * r2
		t3 = t2 + Dis(TT, derPunkt{x2, y2})/V
		if res > t3 {
			res = t3
		}
	}

	d1 = S.theta() + math.Asin(math.Sin(gamma)*V*t1/math.Sqrt(S.sqrlen()))
	d3 = TT.theta()
	if d3 < d1 {
		d3 += 2 * math.Pi
	}
	l = 0
	r = d3 - d1
	if !fls(k/w*math.Log(r3/r1), r) {
		for r-l > eps {
			mid := (l + r) / 2
			d2 = d1 + mid
			r2 = math.Cos(d3-d2)*r3 - math.Sin(d3-d2)*r3*w/k
			if k/w*math.Log(r2/r1) < mid {
				l = mid
			} else {
				r = mid
			}
		}
		d2 = d1 + l
		r2 = math.Cos(d3-d2)*r3 - math.Sin(d3-d2)*r3*w/k
		t2 = (r2 - r0) / w
		x2 = math.Cos(d2) * r2
		y2 = math.Sin(d2) * r2
		t3 = t2 + Dis(TT, derPunkt{x2, y2})/V
		if res > t3 {
			res = t3
		}
	}
	if res > 1e17 {
		fmt.Println("Impossible")
	} else {
		fmt.Println(res)
	}
	return
}

func dos() {
	N += 2
	AA[N] = derKreis{S, 0}
	AA[N-1] = derKreis{TT, 0}
	for i := 1; i <= N; i++ {
		for j := i + 1; j <= N; j++ {
			A := AA[i]
			B := AA[j]
			if A.r < B.r {
				A, B = B, A
			}
			C := derKreis{A.O, A.r - B.r}
			D := B.O
			l := dieGerade{C.O, D.minus(C.O)}
			if fle(l.len(), C.r) {
				continue
			}
			theta := math.Acos(C.r / l.len())
			alpha := l.theta() + theta
			beta := l.theta() - theta
			CP++
			P[CP] = derPunkt{A.O.x + math.Cos(alpha)*A.r, A.O.y + math.Sin(alpha)*A.r}
			CP++
			P[CP] = derPunkt{B.O.x + math.Cos(alpha)*B.r, B.O.y + math.Sin(alpha)*B.r}
			CP++
			P[CP] = derPunkt{A.O.x + math.Cos(beta)*A.r, A.O.y + math.Sin(beta)*A.r}
			CP++
			P[CP] = derPunkt{B.O.x + math.Cos(beta)*B.r, B.O.y + math.Sin(beta)*B.r}
		}
	}
	for i := 1; i <= N; i++ {
		for j := i + 1; j <= N; j++ {
			A := AA[i]
			B := AA[j]
			if A.r < B.r {
				A, B = B, A
			}
			C := derKreis{A.O, A.r + B.r}
			D := B.O
			l := dieGerade{C.O, D.minus(C.O)}
			if fle(l.len(), C.r) {
				continue
			}
			theta := math.Acos(C.r / l.len())
			alpha := l.theta() + theta
			beta := l.theta() - theta
			CP++
			P[CP] = derPunkt{A.O.x + math.Cos(alpha)*A.r, A.O.y + math.Sin(alpha)*A.r}
			CP++
			P[CP] = derPunkt{B.O.x - math.Cos(alpha)*B.r, B.O.y - math.Sin(alpha)*B.r}
			CP++
			P[CP] = derPunkt{A.O.x + math.Cos(beta)*A.r, A.O.y + math.Sin(beta)*A.r}
			CP++
			P[CP] = derPunkt{B.O.x - math.Cos(beta)*B.r, B.O.y - math.Sin(beta)*B.r}
		}
	}
	var G derGraf
	for i := 0; i < CP; i += 2 {
		G.seg(i+1, i+2)
	}
	for i := 1; i <= CP; i++ {
		for j := i + 1; j <= CP; j++ {
			for k := 1; k <= N; k++ {
				if equal(P[i], AA[k]) && equal(P[j], AA[k]) {
					A := P[i]
					B := P[j]
					O := AA[k].O
					alpha := (A.minus(O)).theta()
					beta := (B.minus(O)).theta()
					if arc(A, B, O, AA[k].r, alpha, beta) {
						G.AE(i, j, dk(alpha, beta, AA[k].r))
					}
					if arc(B, A, O, AA[k].r, beta, alpha) {
						G.AE(j, i, dk(beta, alpha, AA[k].r))
					}
				}
			}
		}
	}
	G.s = CP + 1
	G.t = CP + 2
	P[CP+1] = S
	P[CP+2] = TT
	G.n = CP + 2
	for i := 1; i <= CP; i++ {
		if feq(Dis(S, P[i]), 0) {
			G.AE(G.s, i, 0)
		}
		if feq(Dis(TT, P[i]), 0) {
			G.AE(i, G.t, 0)
		}
	}
	res := G.bf()
	if res > 1e17 {
		fmt.Println("Impossible")
		return
	}
	for i := 1; i <= N; i++ {
		if lessThan(TT, derKreis{AA[i].O, AA[i].r + res/V*W[i]}) {
			fmt.Println("Impossible")
			return
		}
	}
	fmt.Println(G.bf() / V)
	return
}

type derGraf struct {
	s, t, m, n int
	u, v       []int
	w          []float64
	dis        [111111]float64
}

func (d *derGraf) ins(a, b int, c float64) {
	d.u = append(d.u, a)
	d.v = append(d.v, b)
	d.w = append(d.w, c+eps)
	d.m++
}

func (d *derGraf) bf() float64 {
	for i := 1; i <= d.n; i++ {
		d.dis[i] = 1e18
	}
	d.dis[d.s] = 0.0
	for j := 0; j < d.n; j++ {
		for i := 0; i < d.m; i++ {
			d.dis[d.v[i]] = math.Min(d.dis[d.v[i]], d.dis[d.u[i]]+d.w[i])
		}
	}
	return d.dis[d.t]
}

func (g *derGraf) AE(a, b int, c float64) {
	g.ins(a, b, c)
	g.ins(b, a, c)
}

func (g *derGraf) seg(alpha, beta int) {
	A := P[alpha]
	B := P[beta]
	l := dieGerade{A, B.minus(A)}
	for j := 1; j <= N; j++ {
		if mul(AA[j], l) {
			return
		}
	}
	g.AE(alpha, beta, l.len())
}

func dans(a, b, c float64) bool {
	if fls(a, b) {
		return fls(a, c) && fls(c, b)
	} else {
		return fls(c, b) || fls(a, c)
	}
}

func arc(A, B, O derPunkt, r, alpha, beta float64) bool {
	for i := 1; i <= N; i++ {
		if lessThan(A, AA[i]) || lessThan(B, AA[i]) {
			return false
		}
		if fls(math.Abs(AA[i].r-r), Dis(O, AA[i].O)) && fls(Dis(O, AA[i].O), AA[i].r+r) {
			gamma := (AA[i].O.minus(O)).theta()
			if dans(alpha, beta, gamma) {
				return false
			}
		}
	}
	return true
}

func dk(alpha, beta, r float64) float64 {
	if alpha <= beta {
		return (beta - alpha) * r
	} else {
		return (beta - alpha + 2*math.Pi) * r
	}
}

type dieGerade struct {
	O, A derPunkt
}

func (d dieGerade) theta() float64 { return math.Atan2(d.A.y, d.A.x) }
func (d dieGerade) len() float64   { return math.Sqrt(d.A.sqrlen()) }

type derKreis struct {
	O derPunkt
	r float64
}

func lessThan(A derPunkt, B derKreis) bool { return fls((A.minus(B.O)).sqrlen(), B.r*B.r) }
func equal(A derPunkt, B derKreis) bool    { return feq((A.minus(B.O)).sqrlen(), B.r*B.r) }
func mul(O derKreis, l dieGerade) bool {
	A := l.O
	B := l.O.plus(l.A)
	if lessThan(A, O) || lessThan(B, O) {
		return true
	}
	a := -l.A.y
	b := l.A.x
	c := A.div(B)
	d1 := a*O.O.x + b*O.O.y + c
	d2 := O.r * O.r * (a*a + b*b)
	d1 = d1 * d1
	if fls(d1, d2) {
		d1 = (O.O.minus(A)).mul(B.minus(A))
		d2 = (O.O.minus(B)).mul(A.minus(B))
		return fls(0, d1) && fls(0, d2)
	} else {
		return false
	}
}

type derPunkt struct{ x, y float64 }

func (a derPunkt) plus(b derPunkt) derPunkt  { return derPunkt{a.x + b.x, a.y + b.y} }
func (a derPunkt) minus(b derPunkt) derPunkt { return derPunkt{a.x - b.x, a.y - b.y} }
func (a derPunkt) mul(b derPunkt) float64    { return a.x*b.x + a.y*b.y }
func (a derPunkt) div(b derPunkt) float64    { return a.x*b.y - b.x*a.y }
func (d derPunkt) theta() float64            { return math.Atan2(d.y, d.x) }
func (d derPunkt) sqrlen() float64           { return d.x*d.x + d.y*d.y }
func fle(a, b float64) bool                  { return a < b+eps }
func fls(a, b float64) bool                  { return a+eps < b }
func feq(a, b float64) bool                  { return math.Abs(a-b) < eps }
func Dis(A, B derPunkt) float64              { return math.Sqrt((A.minus(B)).sqrlen()) }
