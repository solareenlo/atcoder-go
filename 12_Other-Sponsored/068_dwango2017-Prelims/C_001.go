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
	A := make([]int, 5)
	for i := 0; i < N; i++ {
		var a int
		fmt.Fscan(in, &a)
		A[a]++
	}
	ans := A[4]
	a := min(A[3], A[1])
	ans += a
	A[3] -= a
	A[1] -= a
	if A[3] == 0 {
		ans += (A[1] + A[2]*2 + 3) / 4
	} else {
		ans += (A[2]+1)/2 + A[3]
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
