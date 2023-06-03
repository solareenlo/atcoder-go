package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, s int
	fmt.Fscan(in, &n, &s)
	A := make([]int, n+1)
	B := make([]int, s+1)
	for i := range B {
		B[i] = -1
	}
	B[0] = 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &A[i])
		for j := s - A[i]; j >= 0; j-- {
			if B[j] >= 0 && B[j+A[i]] == -1 {
				B[j+A[i]] = i
			}
		}
	}
	if B[s] == -1 {
		fmt.Println(-1)
		return
	}
	P := make([]int, 0)
	for s > 0 {
		P = append(P, B[s])
		s -= A[B[s]]
	}
	fmt.Println(len(P))
	for i := len(P) - 1; i >= 0; i-- {
		fmt.Printf("%d ", P[i])
	}
}
