package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var N, M, K int
	fmt.Scan(&N, &M, &K)

	F := [12][12]int{}
	for i := 0; i < M; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		F[x][y] = 1
		F[y][x] = 1
	}

	p := [12]int{}
	c := 0
	r := rand.New(rand.NewSource(1 << 32))
	for t := int(1e6); t > 0; t-- {
		for i := 0; i < N; i++ {
			p[i] = i
		}
		for k := 0; k < K; k++ {
			x := r.Int() % N
			y := r.Int() % (N - 1)
			if y >= x {
				y++
			}
			p[x], p[y] = p[y], p[x]
		}
		x := 0
		for i, j := 0, N-1; i < N; {
			x |= F[p[i]][p[j]]
			j = i
			i++
		}
		if x == 0 {
			c++
		}
	}

	fmt.Println(float64(c) * 1e-6)
}
