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

	a := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &a[i])
	}

	sum := 0
	for i := 0; i < N; i++ {
		var x int
		fmt.Fscan(in, &x)
		sum += abs(a[i] - x)
	}

	if K-sum >= 0 && (K-sum)%2 == 0 {
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
