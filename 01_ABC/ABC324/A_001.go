package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}
	sort.Ints(A)
	if A[0] == A[N-1] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
