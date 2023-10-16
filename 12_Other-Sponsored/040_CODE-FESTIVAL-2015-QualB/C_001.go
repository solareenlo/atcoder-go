package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)

	A := make([]int, N)
	B := make([]int, M)

	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}
	for i := 0; i < M; i++ {
		fmt.Fscan(in, &B[i])
	}

	sort.Ints(A)
	sort.Ints(B)

	i := 0
	for j := 0; j < M; j++ {
		for i < N && A[i] < B[j] {
			i++
		}
		if i >= N {
			fmt.Println("NO")
			return
		}
		i++
	}

	fmt.Println("YES")
}
