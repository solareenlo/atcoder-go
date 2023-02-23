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

	var N int
	X := [1000]float64{}
	Y := [1000]float64{}
	T := [1000]float64{}
	R := [1000]float64{}
	g := [1000][1000]float64{}
	cost := make([]float64, 1000)
	v := [1000]bool{}
	for i := range cost {
		cost[i] = 1e9
	}

	fmt.Fscan(in, &N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &X[i], &Y[i], &T[i], &R[i])
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			g[i][j] = math.Hypot(X[i]-X[j], Y[i]-Y[j]) / math.Min(T[i], R[j])
		}
	}

	cost[0] = 0
	for i := 0; i < N; i++ {
		idx := -1
		for j := 0; j < N; j++ {
			if (idx < 0 || cost[idx] >= cost[j]) && !v[j] {
				idx = j
			}
		}
		for j := 1; j < N; j++ {
			cost[j] = math.Min(cost[j], cost[idx]+g[idx][j])
		}
		v[idx] = true
	}
	sort.Slice(cost[:N], func(i, j int) bool {
		return cost[i] < cost[j]
	})

	res := 0.0
	for i := 1; i < N; i++ {
		res = math.Max(res, cost[i]+float64(N-i-1))
	}
	fmt.Println(res)
}
