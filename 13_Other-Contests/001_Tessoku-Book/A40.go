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
	A := make([]int, 101)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		A[a]++
	}
	ans := 0
	for i := 1; i <= 100; i++ {
		if A[i] > 2 {
			ans += A[i] * (A[i] - 1) * (A[i] - 2) / 6
		}
	}
	fmt.Println(ans)
}
