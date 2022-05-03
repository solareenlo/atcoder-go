package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

var (
	par = make([]int, 65)
	sz  = make([]int, 65)
)

func find(u int) int {
	if par[u] == u {
		return u
	}
	par[u] = find(par[u])
	return par[u]
}

func join(u, v int) {
	if sz[u] < sz[v] {
		u, v = v, u
	}
	sz[u] += sz[v]
	par[v] = u
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)
	x := make([]float64, 65)
	y := make([]float64, 65)
	r := make([]float64, 15)
	for i := 1; i <= N+M; i++ {
		fmt.Fscan(in, &x[i], &y[i])
		if i > N {
			fmt.Fscan(in, &r[i-N])
		}
	}

	type tuple struct {
		d    float64
		u, v int
	}
	vec := make([]tuple, 0)
	for i := 1; i <= N+M; i++ {
		for j := i + 1; j <= N+M; j++ {
			dist := math.Sqrt(1.0*(x[i]-x[j])*(x[i]-x[j]) + 1.0*(y[i]-y[j])*(y[i]-y[j]))
			if i > N && j > N {
				if 1.0*r[i-N]+r[j-N] <= dist {
					dist = dist - r[i-N] - r[j-N]
				} else if dist+min(r[i-N], r[j-N]) <= max(r[i-N], r[j-N]) {
					dist = max(r[i-N], r[j-N]) - dist - min(r[i-N], r[j-N])
				} else {
					dist = 0
				}
			} else if i > N {
				dist = abs(dist - r[i-N])
			} else if j > N {
				dist = abs(dist - r[j-N])
			}
			vec = append(vec, tuple{dist, i, j})
		}
	}
	sort.Slice(vec, func(i, j int) bool {
		return vec[i].d < vec[j].d
	})

	ans := 1e18
	for k := 0; k < (1 << M); k++ {
		for i := 1; i <= N+M; i++ {
			par[i] = i
			sz[i] = 1
		}
		cost := 0.0
		for i := range vec {
			d := vec[i].d
			u := vec[i].u
			v := vec[i].v
			if u > N && (k&(1<<(u-N-1))) == 0 {
				continue
			}
			if v > N && (k&(1<<(v-N-1))) == 0 {
				continue
			}
			u = find(u)
			v = find(v)
			if u == v {
				continue
			}
			join(u, v)
			cost += d
		}
		ans = min(ans, cost)
	}
	fmt.Println(ans)
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}
