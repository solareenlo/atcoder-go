package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)
	P := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &P[i])
	}
	for i := range P {
		P[i] = 100 - P[i]
	}
	sort.Ints(P)
	ans := 0
	for i := 0; i < N; i++ {
		ans += 100 * P[i]
		if i < N-1 {
			ans -= P[i] * P[N-1]
		}
	}
	ans = 10000 - min(ans, 10000)
	fmt.Println(float64(ans) / 10000.0)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
