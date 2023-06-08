package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	type pair struct {
		x, y float64
	}

	var x1, x2, Y1, y2 [2100]float64

	var a int
	fmt.Fscan(in, &a)
	r := make([]pair, a*4)
	EPS := 1e-11
	PI := math.Acos(-1.0)
	for i := 0; i < a; i++ {
		fmt.Fscan(in, &x1[i], &Y1[i], &x2[i], &y2[i])
	}
	for i := 0; i < a; i++ {
		L := math.Atan2(Y1[i], x1[i])
		R := math.Atan2(y2[i], x2[i])
		if math.Max(L, R)-math.Min(L, R) > PI {
			if L < R {
				L += PI * 2
			} else {
				R += PI * 2
			}
		}
		r[i] = pair{math.Max(L, R) + EPS, math.Min(L, R) - EPS}
	}
	for i := 0; i < 3*a; i++ {
		r[i+a] = pair{r[i].x + PI*2, r[i].y + PI*2}
	}

	sort.Slice(r, func(i, j int) bool {
		if r[i].x == r[j].x {
			return r[i].y < r[j].y
		}
		return r[i].x < r[j].x
	})

	ret := 99999999
	for i := a; i < 2*a; i++ {
		last := -99999999.0
		val := 0
		for j := i; j < i+a; j++ {
			if last > r[j].y {
				continue
			}
			last = r[j].x
			val++
		}
		ret = min(ret, val)
	}
	fmt.Println(ret)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
