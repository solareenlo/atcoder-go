package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M, K, E int
	fmt.Fscan(in, &N, &M, &K, &E)

	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}
	B := make([]int, M)
	for j := 0; j < M; j++ {
		fmt.Fscan(in, &B[j])
	}
	sort.Ints(A)
	sort.Ints(B)
	B = reverseOrderInt(B)

	id := 0
	for i := 0; i < N; i++ {
		for E < A[i] {
			if id == K {
				fmt.Println("No")
				fmt.Println(i)
				return
			}
			E += B[id]
			id++
		}
		E -= A[i]
	}
	fmt.Println("Yes")
	fmt.Println(id)
}

func reverseOrderInt(a []int) []int {
	n := len(a)
	res := make([]int, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
