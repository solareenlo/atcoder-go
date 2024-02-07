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
	var tim [77]int
	ans := 0
	for i := 1; i <= n; i++ {
		var w, x int
		fmt.Fscan(in, &w, &x)
		for i := x; i <= x+8; i++ {
			tim[i%24] += w
			ans = max(ans, tim[i%24])
		}
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
