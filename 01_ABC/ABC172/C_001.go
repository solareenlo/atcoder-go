package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)

	A := make([]int, n)
	B := make([]int, m)
	for i := range A {
		fmt.Fscan(in, &A[i])
	}
	for i := range B {
		fmt.Fscan(in, &B[i])
	}

	a := make([]int, n+1)
	b := make([]int, m+1)
	for i := 0; i < n; i++ {
		a[i+1] = a[i] + A[i]
	}
	for i := 0; i < m; i++ {
		b[i+1] = b[i] + B[i]
	}

	res := 0
	for i := 0; i < n+1; i++ {
		if a[i] > k {
			break
		}
		for b[m] > k-a[i] {
			m--
		}
		res = max(res, i+m)
	}

	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
