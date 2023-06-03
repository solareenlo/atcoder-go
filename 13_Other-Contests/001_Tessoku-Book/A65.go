package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	A := make([]int, n+1)
	B := make([]int, n+1)
	for i := 2; i <= n; i++ {
		fmt.Fscan(in, &A[i])
	}
	for i := n; i >= 1; i-- {
		B[A[i]] += (B[i] + 1)
	}
	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", B[i])
	}
}
