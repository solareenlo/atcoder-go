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

	p := make([]int, n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		p[a]++
	}

	o := 0
	for i := 0; i < n; i++ {
		k = min(k, p[i])
		o += k
	}
	fmt.Println(o)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
