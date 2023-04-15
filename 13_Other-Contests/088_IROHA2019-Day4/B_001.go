package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M, L int
	fmt.Fscan(in, &N, &M, &L)
	ans := L * M
	for i := 0; i < N; i++ {
		var A, B int
		fmt.Fscan(in, &A, &B)
		ans = min(ans, A+B*M)
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
