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

	ANS := 0.0
	var N int
	fmt.Fscan(in, &N)
	X := make([]float64, N)
	Y := make([]float64, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &X[i], &Y[i])
	}
	for i := 0; i < N; i++ {
		A := make([]float64, 0)
		for j := 0; j < N; j++ {
			if i == j {
				continue
			}
			A = append(A, math.Atan2(Y[j]-Y[i], X[j]-X[i])*180.0/math.Pi)
		}
		sort.Float64s(A)
		for j := 0; j < N-1; j++ {
			X := LowerBound(A, A[j]+180.0)
			if ANS < A[X-1]-A[j] {
				ANS = A[X-1] - A[j]
			}
			if X < N-1 {
				if A[X]-A[j] >= 180.0 && ANS < 360.0-(A[X]-A[j]) {
					ANS = 360.0 - A[X] + A[j]
				}
			}
		}
	}
	fmt.Println(ANS)
}

func LowerBound(a []float64, x float64) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}
