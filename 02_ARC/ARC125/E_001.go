package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	A := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &A[i])
	}
	B := make([]int, m+1)
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &B[i])
	}
	C := make([]int, m+1)
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &C[i])
	}
	tmp := A[1 : n+1]
	sort.Ints(tmp)

	sum2 := 0
	D := make([]int, 200002)
	E := make([]int, 200002)
	for i := 1; i <= m; i++ {
		x := C[i] / B[i]
		sum2 += C[i]
		D[max(n-x, 0)] += B[i]
		E[max(n-x, 0)] += C[i]
	}

	sum0, sum1 := 0, 0
	ans := 1 << 60
	for i := 0; i <= n; i++ {
		sum0 += A[i]
		sum1 += D[i]
		sum2 -= E[i]
		ans = min(ans, sum0+sum1*(n-i)+sum2)
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
