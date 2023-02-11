package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, K int
	fmt.Fscan(in, &N, &K)

	s := 0
	for i := 0; i < N; i++ {
		var A int
		fmt.Fscan(in, &A)
		s += abs(A)
	}

	if K >= s && (K-s)%2 == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
