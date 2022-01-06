package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	N int
	M int
	a = [20]float64{}
	b = [20]float64{}
	c = [20]float64{}
	p = [20]float64{}
	q = [20]float64{}
)

func score(x, y float64) float64 {
	road := float64(1 << 62)
	pnt2 := float64(1 << 62)
	for i := 0; i < N; i++ {
		road = math.Min(road, math.Abs(a[i]*x+b[i]*y+c[i]))
	}
	for i := 0; i < M; i++ {
		pnt2 = math.Min(pnt2, math.Pow(x-p[i], 2)+math.Pow(y-q[i], 2))
	}
	return road + pnt2
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var R float64
	fmt.Fscan(in, &N, &M, &R)

	for i := 0; i < N; i++ {
		fmt.Fscan(in, &a[i], &b[i], &c[i])
	}
	for i := 0; i < M; i++ {
		fmt.Fscan(in, &p[i], &q[i])
	}

	for i := 0; i < N; i++ {
		norm := math.Hypot(a[i], b[i])
		a[i] /= norm
		b[i] /= norm
		c[i] /= norm
	}

	x := 0.0
	y := 0.0
	s := R / 100.0
	for i := 0; i < 15; i++ {
		best := 0.0
		bx := 0.0
		by := 0.0
		for dx := -100; dx < 101; dx++ {
			for dy := -100; dy < 101; dy++ {
				nx := x + s*float64(dx)
				ny := y + s*float64(dy)
				if !(-R <= nx && nx <= R && -R <= ny && ny <= R) {
					continue
				}
				v := score(nx, ny)
				if best < v {
					best = v
					bx = nx
					by = ny
				}
			}
		}
		x = bx
		y = by
		s /= 10.0
	}

	fmt.Println(score(x, y))
}
