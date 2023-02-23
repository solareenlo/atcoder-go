package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	A := make([]int, 50050)
	M := make([]int, 50050)
	var i int
	for {
		var N, K int
		fmt.Fscan(in, &N, &K)
		if N == -1 {
			break
		}
		for i = 0; i < N; i++ {
			fmt.Fscan(in, &A[i])
		}
		K++
		M[0] = A[0]
		for i = 1; i < N; i++ {
			M[i] = A[i]
			if M[i] < M[i-1] {
				M[i] = M[i-1]
			}
		}
		for i = K; i < N; i++ {
			if M[i-K] > A[i] {
				break
			}
		}
		if i < N {
			fmt.Println("No")
		} else {
			fmt.Println("Yes")
		}
	}
}
