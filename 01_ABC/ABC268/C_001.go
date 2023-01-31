package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 200020

	var n int
	fmt.Fscan(in, &n)

	var c [N]int
	for i := 0; i < n; i++ {
		var p int
		fmt.Fscan(in, &p)
		c[(p-i+n)%n]++
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, c[i]+c[(i+1)%n]+c[(i+n-1)%n])
	}

	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
