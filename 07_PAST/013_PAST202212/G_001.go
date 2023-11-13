package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, res int
	fmt.Fscan(in, &n, &res)

	ans := res
	for i := 1; i < n; i++ {
		var x int
		fmt.Fscan(in, &x)
		res = max(res+x, x)
		ans = max(ans, res)
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
