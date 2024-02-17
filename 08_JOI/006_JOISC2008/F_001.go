package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

const PI = 3.1415926535

var V []pair

func pu(a, b float64) int {
	for b >= 0 {
		a -= 2 * PI
		b -= 2 * PI
	}
	a += 2 * PI
	b += 2 * PI
	if a >= 0 {
		V = append(V, pair{a, +1})
	}
	V = append(V, pair{b, -1})
	if a < 0 {
		return 1
	}
	return 0
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var x, y [1010]float64

	var n int
	var d float64
	fmt.Fscan(in, &n, &d)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i], &y[i])
	}
	res := 0
	for i := 0; i < n; i++ {
		V = make([]pair, 0)
		cnt := 0
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			p := x[j] - x[i]
			q := y[j] - y[i]
			r := math.Sqrt(p*p + q*q)
			if 2*d > r {
				t := math.Atan2(p, q)
				cnt += pu(t, t+PI)
			} else {
				t := math.Atan2(p, q)
				s := math.Asin(2 * d / r)
				cnt += pu(t, t+s)
				t += PI
				cnt += pu(t-s, t)
			}
		}
		sortPair(V)
		for _, e := range V {
			if res < cnt {
				res = cnt
			}
			cnt += e.y
		}
		if res < cnt {
			res = cnt
		}
	}
	fmt.Println(res + 1)
}

type pair struct {
	x float64
	y int
}

func sortPair(tmp []pair) {
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].x == tmp[j].x {
			return tmp[i].y < tmp[j].y
		}
		return tmp[i].x < tmp[j].x
	})
}
