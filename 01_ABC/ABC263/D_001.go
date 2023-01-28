package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, l, r int
	fmt.Fscan(in, &n, &l, &r)

	ans := r * n
	s := 0
	for i := 1; i <= n; i++ {
		var a int
		fmt.Fscan(in, &a)
		s = min(s+a, i*l)
		ans = min(ans, s+(n-i)*r)
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
