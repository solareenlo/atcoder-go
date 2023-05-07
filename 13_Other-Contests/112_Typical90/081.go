package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)

	var x [5002][5002]int
	for n > 0 {
		n--
		var a, b int
		fmt.Fscan(in, &a, &b)
		x[a+1][b+1]++
	}

	for i := 2; i < 5002; i++ {
		for j := 2; j < 5002; j++ {
			x[i][j] += x[i-1][j] + x[i][j-1] - x[i-1][j-1]
			if i > k && j > k {
				n = max(n, x[i][j]-x[i-k-1][j]-x[i][j-k-1]+x[i-k-1][j-k-1])
			}
		}
	}
	fmt.Println(n)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
