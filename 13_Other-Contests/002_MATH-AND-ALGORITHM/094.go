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

	b := make([]int, n-1)
	for i := range b {
		fmt.Fscan(in, &b[i])
	}

	ans := b[0] + b[n-2]
	for i := 0; i < n-2; i++ {
		ans += min(b[i], b[i+1])
	}

	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
