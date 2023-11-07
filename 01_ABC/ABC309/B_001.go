package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	fmt.Scan(&N)
	A := make([][]string, N+1)
	for i := 0; i < N; i++ {
		var a string
		fmt.Scan(&a)
		a = a + " "
		A[i] = strings.Split(a, "")
	}
	A[N] = strings.Split(strings.Repeat(" ", len(A[0])), "")
	B := make([][]string, N+1)
	for i := range B {
		B[i] = make([]string, len(A[i]))
		copy(B[i], A[i])
	}
	for i := 0; i < N-1; i++ {
		B[0][i+1] = A[0][i]
		B[i+1][N-1] = A[i][N-1]
	}
	for i := 0; i < N-1; {
		i++
		B[N-1][i-1] = A[N-1][i]
		B[i-1][0] = A[i][0]
	}
	for i := 0; i < N; i++ {
		fmt.Println(strings.Join(B[i], ""))
	}
}
