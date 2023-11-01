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
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 1; i <= n; i++ {
		x := 2 * n
		for j := 1; j <= x; j++ {
			if i > j {
				x = min(x, abs(a[i]-a[i-j])+j)
			}
			if i+j <= n {
				x = min(x, abs(a[i]-a[i+j])+j)
			}
		}
		fmt.Printf("%d ", x)
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
