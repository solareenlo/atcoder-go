package main

import "fmt"

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	A := make([]int, M)
	for i := range A {
		fmt.Scan(&A[i])
	}

	j := 0
	for i := 1; i <= N; i++ {
		if i > A[j] {
			j++
		}
		fmt.Println(A[j] - i)
	}
}
