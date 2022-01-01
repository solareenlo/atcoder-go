package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, s, g int
	fmt.Fscan(in, &n, &s, &g)

	a := make([]int, 3*n)
	for i := 0; i < 2*n; i++ {
		fmt.Fscan(in, &a[i])
	}

	a[2*n] = g
	a[n*2+1] = g
	q := 0
	p := s
	t := 0.0
	for {
		y := q + 1
		m := q + 1
		b := q + 1
		l := a[2*y]
		r := a[y*2+1]
		c := float64(l - p)
		f := float64(r - p)
		for ; y <= n; y++ {
			l = a[y*2]
			r = a[y*2+1]
			d := float64(l - p)
			d /= float64(y - q)
			e := float64(r - p)
			e /= float64(y - q)
			if f < d {
				x := a[b*2+1]
				t += math.Hypot(float64(p-x), float64(q-b))
				p = x
				q = b
				break
			}
			if c > e {
				x := a[m*2]
				t += math.Hypot(float64(p-x), float64(q-m))
				p = x
				q = m
				break
			}
			if c < d {
				c = d
				m = y
			}
			if f > e {
				f = e
				b = y
			}
		}
		if y > n {
			t += math.Hypot(float64(p-g), float64(q-n))
			break
		}
	}
	fmt.Println(t)
}
