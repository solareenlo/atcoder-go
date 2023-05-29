package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)
	var A [200020]int
	for i := 0; i < N; i++ {
		A[i] = 1
	}
	for i := 0; i < M; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		A[a] = 0
	}
	sum := 0
	for i := 0; i < N; i++ {
		sum += A[i]
	}
	fmt.Println(sum)
	for i := 0; i < N; i++ {
		if A[i] != 0 {
			fmt.Printf("%d ", i+1)
		}
	}
	fmt.Println()
}
