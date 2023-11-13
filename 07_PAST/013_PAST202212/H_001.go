package main

import (
	"fmt"
	"strconv"
)

func main() {
	var N int
	var S string
	fmt.Scan(&N, &S)
	A := make([]int, 10)
	for i := 0; i < N; i++ {
		tmp, _ := strconv.Atoi(string(S[i]))
		A[tmp]++
	}
	ans := 0
	for i := 0; i < 10; i++ {
		for j := i + 1; j < 10; j++ {
			ans += (j - i) * A[i] * A[j]
		}
	}
	fmt.Println(ans)
}
