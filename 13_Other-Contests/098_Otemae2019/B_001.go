package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var m, n, k int
	fmt.Fscan(in, &m, &n, &k)

	var y [5005]int
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(in, &x)
		y[x-1]++
	}

	ans := 0
	for i := 0; i < m; i++ {
		t := y[i]
		for j := 0; j < k; j++ {
			if (i-j-1 >= 0 && y[i-j-1] != 0) || (i+j+1 < m && y[i+j+1] != 0) {
				t++
			}
		}
		ans = max(ans, t)
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
