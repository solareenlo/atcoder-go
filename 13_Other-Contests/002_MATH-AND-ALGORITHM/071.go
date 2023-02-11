package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)

	A := make([]float64, N)
	B := make([]float64, N)
	C := make([]float64, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i], &B[i], &C[i])
	}

	ans := 0.0

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if A[i]*B[j] == B[i]*A[j] {
				continue
			}
			y := (C[i]*A[j] - C[j]*A[i]) / (B[i]*A[j] - B[j]*A[i])
			x := (C[i]*B[j] - C[j]*B[i]) / (A[i]*B[j] - A[j]*B[i])
			satisfied := true
			for k := 0; k < N; k++ {
				if A[k]*x+B[k]*y > C[k] {
					satisfied = false
					break
				}
			}
			if satisfied {
				ans = math.Max(ans, x+y)
			}
		}
	}

	fmt.Println(ans)
}
