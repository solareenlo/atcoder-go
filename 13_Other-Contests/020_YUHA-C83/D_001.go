package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const EPS = 1e-10
const INF = 1e+10

var p1, p2 []Pt

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 41000

	var a int
	var b float64
	fmt.Fscan(in, &a, &b)
	p1 = make([]Pt, N)
	p2 = make([]Pt, N)
	for i := 0; i < a; i++ {
		var x1, y1, x2, y2 float64
		fmt.Fscan(in, &x1, &y1, &x2, &y2)
		p1[i] = Pt{x1, y1}
		p2[i] = Pt{x2, y2}
	}
	r := make([]float64, N)
	len := make([]float64, N)
	for i := 0; i < a-1; i++ {
		if iLL(p1[i], p2[i], p1[i+1], p2[i+1]) == 0 {
			d := (p2[i].minus(p1[i+1])).ABS()
			r[i] = d / 2
			len[i] = d * math.Pi / 2
		} else {
			tmp1 := p2[i].plus(((p1[i].minus(p2[i])).mulPt(Pt{0, 1})))
			tmp2 := p1[i+1].plus(((p2[i+1].minus(p1[i+1])).mulPt(Pt{0, 1})))
			p := pLL(p2[i], tmp1, p1[i+1], tmp2)
			r[i] = p.minus(p2[i]).ABS()
			d := p2[i].minus(p1[i+1]).ABS() / 2
			th := math.Asin(d/r[i]) * 2
			if p2[i].minus(p1[i]).dot(p1[i+1].minus(p2[i])) < -EPS {
				th += math.Pi
			}
			len[i] = r[i] * th
		}
	}
	now := 0.0
	v := make([]float64, N)
	for i := 0; i < a-1; i++ {
		now = math.Sqrt(p2[i].minus(p1[i]).ABS()*2*b + now*now)
		now = math.Min(now, math.Sqrt(r[i]*b))
		v[i+1] = now
	}
	now = 0
	for i := a - 1; i > 0; i-- {
		now = math.Sqrt(p2[i].minus(p1[i]).ABS()*2*b + now*now)
		now = math.Min(now, math.Sqrt(r[i-1]*b))
		now = math.Min(now, v[i])
		v[i] = now
	}
	ret := 0.0
	for i := 0; i < a; i++ {
		ret += gett(b, p2[i].minus(p1[i]).ABS(), v[i], v[i+1])
		if i < a-1 {
			ret += len[i] / v[i+1]
		}
	}
	fmt.Println(ret)
}

func gett(a, s, v1, v2 float64) float64 {
	vt := math.Sqrt(a*s + (v1*v1+v2*v2)/2)
	return (2*vt - v1 - v2) / a
}

func pLL(a, b, c, d Pt) Pt {
	b = b.minus(a)
	d = d.minus(c)
	tmp1 := c.minus(a).det(d)
	tmp2 := b.det(d)
	return a.plus(b.mul(tmp1).div(tmp2))
}

func iLL(a, b, c, d Pt) int {
	if sig((b.minus(a)).det(d.minus(c))) != 0 {
		return 1
	}
	if sig((b.minus(a)).det(c.minus(a))) != 0 {
		return 0
	}
	return -1
}

func sig(r float64) int {
	if r < -EPS {
		return -1
	}
	if r > +EPS {
		return 1
	}
	return 0
}

type Pt struct {
	x, y float64
}

func (b Pt) plus(a Pt) Pt {
	return Pt{b.x + a.x, b.y + a.y}
}

func (b Pt) minus(a Pt) Pt {
	return Pt{b.x - a.x, b.y - a.y}
}

func (b Pt) det(a Pt) float64 {
	return b.x*a.y - b.y*a.x
}

func (b Pt) dot(a Pt) float64 {
	return b.x*a.x + b.y*a.y
}

func (b Pt) ABS() float64 {
	return math.Sqrt(b.x*b.x + b.y*b.y)
}

func (b Pt) mulPt(a Pt) Pt {
	return Pt{b.x*a.x - b.y*a.y, b.x*a.y + b.y*a.x}
}

func (b Pt) mul(k float64) Pt {
	return Pt{b.x * k, b.y * k}
}

func (b Pt) div(k float64) Pt {
	return Pt{b.x / k, b.y / k}
}
