package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var dpR, dpL []int
var T [300000]int
var cur, lis int

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = 1145141919

	type P struct {
		x, y int
	}

	var N, W, H int
	fmt.Fscan(in, &N, &W, &H)
	var X, Y [300000]int
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &X[i+1], &Y[i+1])
	}
	X[N+1] = W
	Y[N+1] = H
	N += 2
	dpR = make([]int, N+1)
	for i := 0; i < N+1; i++ {
		dpR[i] = INF
	}
	dpR[0] = -INF
	evt := make([]P, N+1)
	for i := N - 1; i >= 0; i-- {
		pos := upperBound(dpR, -Y[i])
		evt[i] = P{pos, dpR[pos]}
		dpR[pos] = -Y[i]
	}
	lis = lowerBound(dpR, INF) - 1

	dpL = make([]int, N+1)
	s := 0
	for i := 0; i < N+1; i++ {
		dpL[i] = INF
	}
	dpL[0] = -INF
	for i := 0; i < N-1; i++ {
		pos := upperBound(dpL, Y[i])
		dpL[pos] = Y[i]
		dpR[evt[i].x] = evt[i].y
		recalc(pos)
		recalc(lis - evt[i].x)
		s += ((X[i+1] - X[i]) / 2) * cur
	}
	fmt.Println(s)
}

func change(x, val int) {
	cur += val - T[x]
	T[x] = val
}

func recalc(a int) {
	l := dpL[a]
	r := -dpR[lis-a]
	change(a, max(0, (r-l)/2))
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
