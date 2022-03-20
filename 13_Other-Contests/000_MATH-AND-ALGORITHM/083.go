package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)

	A := make([]int, N+1)
	for i := 1; i <= N; i++ {
		fmt.Fscan(in, &A[i])
	}
	B := make([]int, N+1)
	for i := 1; i <= N; i++ {
		fmt.Fscan(in, &B[i])
	}
	sort.Ints(A)
	sort.Ints(B)

	ans := 0
	for i := 1; i <= N; i++ {
		ans += abs(A[i] - B[i])
	}
	fmt.Println(ans)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
