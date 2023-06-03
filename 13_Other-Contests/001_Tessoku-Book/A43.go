package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, l int
	fmt.Fscan(in, &n, &l)
	ans := 0
	for i := 0; i < n; i++ {
		var a int
		var b string
		fmt.Fscan(in, &a, &b)
		if b == "E" {
			ans = max(ans, l-a)
		} else {
			ans = max(ans, a)
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
