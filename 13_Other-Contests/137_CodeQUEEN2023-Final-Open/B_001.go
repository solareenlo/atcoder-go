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
	X := make([]int, N-1)
	Y := make([]int, N-1)
	a := N * (N - 1) / 2
	b := (N - 1) * N / 2
	for i := 0; i < N-1; i++ {
		fmt.Fscan(in, &X[i], &Y[i])
		X[i]--
		Y[i]--
		a -= X[i]
		b -= Y[i]
	}
	for i := 0; i < N-1; i++ {
		if abs(a-X[i]) == abs(b-Y[i]) {
			fmt.Println(-1)
			return
		}
	}
	fmt.Println(a+1, b+1)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
