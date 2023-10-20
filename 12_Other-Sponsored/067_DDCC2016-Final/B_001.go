package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var R float64
	var N, M int
	fmt.Scan(&R, &N, &M)
	L := make([]float64, 0)
	for i := 1; i < N; i++ {
		length := 4.0 * R / float64(N) * math.Pow(float64(i*N-i*i), 0.5)
		L = append(L, length)
	}
	sort.Slice(L, func(i, j int) bool {
		return L[i] > L[j]
	})
	ans := 0.0
	for i := 0; i < N-1; i++ {
		if i < M {
			ans += L[i]
		}
		if i >= M*2 && i%2 == 0 {
			ans += L[i]
		}
	}
	fmt.Println(ans)
}
