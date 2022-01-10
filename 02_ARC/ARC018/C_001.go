package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, M, x, a, p int
	fmt.Scan(&N, &M, &x, &a, &p)

	type pair struct{ x, y int }
	D := make([]pair, N*M)
	D[0] = pair{x, 0}
	for i := 1; i < N*M; i++ {
		x += a
		x %= p
		D[i] = pair{x, i}
	}
	sort.Slice(D, func(i, j int) bool {
		return D[i].x < D[j].x || (D[i].x == D[j].x && D[i].y < D[j].y)
	})

	d := make([]int, M)
	ans := 0
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			ans += abs(D[i*M+j].y/M - i)
			d[j] = D[i*M+j].y % M
		}
		sort.Ints(d)
		for j := 0; j < M; j++ {
			ans += abs(d[j] - j)
		}
	}
	fmt.Println(ans)

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
