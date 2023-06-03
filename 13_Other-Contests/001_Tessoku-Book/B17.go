package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	A := make([]int, n)
	for i := range A {
		fmt.Fscan(in, &A[i])
	}
	B := make([]int, n)
	B[1] = abs(A[1] - A[0])
	C := make([]int, n)
	for i := range C {
		C[i] = -1
	}
	C[1] = 0
	c := n - 1
	for i := 2; i < n; i++ {
		a := B[i-1] + abs(A[i]-A[i-1])
		b := B[i-2] + abs(A[i]-A[i-2])
		if a > b {
			B[i] = b
			C[i] = i - 2
		} else {
			B[i] = a
			C[i] = i - 1
		}
	}
	P := make([]int, 0)
	for c >= 0 {
		P = append(P, c)
		c = C[c]
	}
	fmt.Println(len(P))
	for i := len(P) - 1; i >= 0; i-- {
		fmt.Printf("%d ", P[i]+1)
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
