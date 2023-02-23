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
	var points string
	fmt.Fscan(in, &points)
	num := make([]int, 4)
	for i := 0; i < n; i++ {
		num[points[i]-'0'-1]++
	}

	maxi, mini := 0, int(1e18)
	for i := 0; i < 4; i++ {
		mini = min(mini, num[i])
		maxi = max(maxi, num[i])
	}

	fmt.Println(maxi, mini)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
