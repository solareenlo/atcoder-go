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
	A := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &A[i])
	}
	ans := 0.0
	for i := 1; i < n; i++ {
		ans += A[i] - A[i-1]
	}
	fmt.Println(ans / float64(n-1))
}
