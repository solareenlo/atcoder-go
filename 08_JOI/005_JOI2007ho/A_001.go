package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 200005

	var n, k int
	fmt.Fscan(in, &n, &k)
	var s [N]int
	for i := 1; i <= n; i++ {
		var x int
		fmt.Fscan(in, &x)
		s[i] = s[i-1] + x
	}
	ans := s[k]
	for i := k; i <= n; i++ {
		ans = max(ans, s[i]-s[i-k])
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
