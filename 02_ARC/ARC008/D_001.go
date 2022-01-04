package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)

	P := make([]int, M)
	A := make([]float64, M)
	B := make([]float64, M)
	v := make([]int, 0)
	for i := 0; i < M; i++ {
		fmt.Fscan(in, &P[i], &A[i], &B[i])
		v = append(v, P[i])
	}
	sort.Ints(v)
	v = unique(v)

	SX := make([]float64, 1<<18-1)
	SY := make([]float64, 1<<18-1)
	for i := range SX {
		SX[i] = 1.0
	}

	mini, maxi := 1.0, 1.0
	for i := 0; i < M; i++ {
		p := lowerBound(v, P[i]) + (1 << 17) - 1
		SX[p] = A[i]
		SY[p] = B[i]
		for p > 0 {
			p = (p - 1) / 2
			SX[p] = SX[p*2+1] * SX[p*2+2]
			SY[p] = SX[p*2+2]*SY[p*2+1] + SY[p*2+2]
		}
		mini = min(mini, SX[0]+SY[0])
		maxi = max(maxi, SX[0]+SY[0])
	}

	fmt.Println(mini)
	fmt.Println(maxi)
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

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func unique(a []int) []int {
	occurred := map[int]bool{}
	result := []int{}
	for i := range a {
		if occurred[a[i]] != true {
			occurred[a[i]] = true
			result = append(result, a[i])
		}
	}
	sort.Ints(result)
	return result
}
