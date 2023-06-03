package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	A := make([]int, n)
	B := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &A[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &B[i])
	}
	sort.Ints(A)
	sort.Slice(B, func(i, j int) bool {
		return B[i] > B[j]
	})
	ans := 0
	for i := 0; i < n; i++ {
		ans += A[i] * B[i]
	}
	fmt.Println(ans)
}
