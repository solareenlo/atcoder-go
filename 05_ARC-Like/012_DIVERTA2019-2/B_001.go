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

	x := make([]int, n)
	y := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i], &y[i])
	}

	type pair struct{ x, y int }
	mp := map[pair]int{}
	mx := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j {
				mp[pair{x[j] - x[i], y[j] - y[i]}]++
				mx = max(mx, mp[pair{x[j] - x[i], y[j] - y[i]}])
			}
		}
	}

	fmt.Println(n - mx)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
