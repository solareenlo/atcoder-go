package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)
	a := make([]int, N-1)
	b := make([]int, N)
	for i := 0; i < N-1; i++ {
		fmt.Fscan(in, &a[i])
	}
	b[0] = a[0]
	b[N-1] = a[N-2]
	for i := 1; i < N-1; i++ {
		b[i] = min(a[i-1], a[i])
	}
	for i := 0; i < N; i++ {
		fmt.Printf("%d ", b[i])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
