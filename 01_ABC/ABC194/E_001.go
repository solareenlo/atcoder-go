package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	pos := make([]int, n)
	for i := 0; i < n; i++ {
		pos[i] = -1
	}

	res := n
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		if i > m+pos[a] {
			res = min(res, a)
		}
		pos[a] = i
	}

	for i := 0; i < n; i++ {
		if n > m+pos[i] {
			res = min(res, i)
		}
	}

	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
