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
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &A[i])
	}
	ans := 0
	for i := 0; i < M; i++ {
		var a int
		fmt.Fscan(in, &a)
		ans += A[a-1]
	}
	fmt.Println(ans)
}
