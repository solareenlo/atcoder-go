package main

import (
	"fmt"
	"math"
)

const epsilon = 1e-8

var (
	x         [64]float64
	y         [64]float64
	c         [64]float64
	X         [64 * 64 * 2]float64
	Y         [64 * 64 * 2]float64
	R         [64]float64
	tot, N, K int
)

func dist(x, y, c, X, Y float64) float64 {
	return c * math.Sqrt((X-x)*(X-x)+(Y-y)*(Y-y))
}

func cal(x1, y1, r1, x2, y2, r2 float64) {
	dd := (x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)
	if dd == 0.0 {
		return
	}
	ad := (r1*r1 + dd - r2*r2) / 2
	h := dd*r1*r1 - ad*ad
	if h < 0.0 {
		return
	}
	h = math.Sqrt(h)
	tot++
	X[tot] = x1 + ad*(x2-x1)/dd - (y2-y1)*h/dd
	Y[tot] = y1 + ad*(y2-y1)/dd + (x2-x1)*h/dd
	tot++
	X[tot] = x1 + ad*(x2-x1)/dd + (y2-y1)*h/dd
	Y[tot] = y1 + ad*(y2-y1)/dd - (x2-x1)*h/dd
}

func judge(t float64) bool {
	for i := 1; i <= N; i++ {
		R[i] = t / c[i]
	}
	tot = 0
	for i := 1; i <= N; i++ {
		tot++
		X[tot] = x[i]
		Y[tot] = y[i]
		for j := i + 1; j <= N; j++ {
			cal(x[i], y[i], R[i], x[j], y[j], R[j])
		}
	}
	for i := 1; i <= tot; i++ {
		cnt := 0
		for j := 1; j <= N; j++ {
			if dist(x[j], y[j], c[j], X[i], Y[i]) < t+epsilon {
				cnt++
			}
		}
		if cnt >= K {
			return true
		}
	}
	return false
}

func main() {
	fmt.Scan(&N, &K)
	for i := 1; i < N+1; i++ {
		fmt.Scan(&x[i], &y[i], &c[i])
	}
	l, r := 0.0, 1000000.0
	for l+epsilon < r {
		mid := (l + r) / 2.0
		if judge(mid) {
			r = mid
		} else {
			l = mid
		}
	}
	fmt.Println(r)
}
