package main

import (
	"bufio"
	"fmt"
	"math"
	"math/cmplx"
	"os"
	"sort"
)

const eps = 1e-8

type CC struct {
	p    complex128
	r, w float64
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	v := make([]CC, n)
	dp := make([]RANGE, n)
	for i := 0; i < n; i++ {
		var x, y, r, w float64
		fmt.Fscan(in, &x, &y, &r, &w)
		v[i] = CC{complex(x, y), r, w}
	}
	var norm func(complex128, complex128) float64
	norm = func(a, b complex128) float64 {
		tmp1 := real(a) * real(b)
		tmp2 := imag(a) * imag(b)
		return tmp1 - tmp2
	}
	sort.Slice(v, func(i, j int) bool {
		tmp1 := norm(v[i].p, cmplx.Conj(v[i].p)) - v[i].r*v[i].r
		tmp2 := norm(v[j].p, cmplx.Conj(v[j].p)) - v[j].r*v[j].r
		return tmp1 < tmp2
	})
	for i := 0; i < n; i++ {
		dp[i] = circle_range(CC{complex(0.0, 0.0), 0.0, 0.0}, v[i])
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			Range := circle_range(v[i], v[j])
			dp[j] = dp[j].plus(dp[i].mul(Range))
		}
	}
	res := 0.0
	for i := 0; i < n; i++ {
		res += dp[i].width * v[i].w / 2.0 / math.Pi
	}
	fmt.Println(res)
}

func circle_range(c1, c2 CC) RANGE {
	w := math.Asin((c1.r + c2.r) / cmplx.Abs(c1.p-c2.p))
	x := c2.p - c1.p
	a := math.Atan2(imag(x), real(x)) - w
	if a < -math.Pi {
		a += 2.0 * math.Pi
	}
	return RANGE{a, w * 2.0}
}

type RANGE struct {
	angle float64
	width float64
}

func (a RANGE) mul(b RANGE) RANGE {
	a1 := a.angle
	w1 := a.width
	a2 := b.angle
	w2 := b.width
	if a1 > a2 {
		a1, a2 = a2, a1
		w1, w2 = w2, w1
	}
	if a2-a1 < math.Pi {
		d := a1 + w1 - a2
		if d < eps {
			return RANGE{0.0, 0.0}
		}
		return RANGE{a2, math.Min(d, w2)}
	} else {
		d := a2 + w2 - a1 - 2.0*math.Pi
		if d < eps {
			return RANGE{0.0, 0.0}
		}
		return RANGE{a1, math.Min(d, w1)}
	}
}

func (a RANGE) plus(b RANGE) RANGE {
	if b.width < eps {
		return a
	}
	a1 := a.angle
	w1 := a.width
	a2 := b.angle
	w2 := b.width
	if a1 > a2 {
		a1, a2 = a2, a1
		w1, w2 = w2, w1
	}
	if a2-a1 < math.Pi {
		return RANGE{a1, math.Max(w1, a2+w2-a1)}
	} else {
		return RANGE{a2, math.Max(w2, a1+w1+2*math.Pi-a2)}
	}
}
