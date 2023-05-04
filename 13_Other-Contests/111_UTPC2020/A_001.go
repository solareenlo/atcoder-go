package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, L int
	fmt.Fscan(in, &N, &L)

	T := 0
	min := 1
	bx := 0
	for i := 0; i < N; i++ {
		var X, A int
		fmt.Fscan(in, &X, &A)
		T += X - bx
		if T > 0 {
			T = 0
		}
		T -= A
		if min > T {
			min = T
		}
		bx = X
	}
	fmt.Println(abs(min))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
