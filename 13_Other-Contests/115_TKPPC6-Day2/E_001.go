package main

import "fmt"

func main() {
	var K int
	fmt.Scan(&K)
	N := 3
	for N*(N-1)*(N-2)/6 < K {
		N++
	}
	K = N*(N-1)*(N-2)/6 - K

	const MAX = 10000000000

	A := make([]int, N)
	for i := range A {
		A[i] = MAX
	}
	for i := 0; i < N-1; i++ {
		A[i] = MAX/2 + i
	}

	for i := N - 2; i >= 0; i-- {
		if K <= 0 {
			break
		}
		A[i] = MAX/2 + 1
		for j := i - 1; j >= 0; j-- {
			if K <= 0 {
				break
			}
			K--
			A[i]--
		}
		if K <= 0 {
			break
		}
		A[i]--
	}

	fmt.Println(N)
	for i := 0; i < N; i++ {
		fmt.Println(A[i])
	}
	return
}
