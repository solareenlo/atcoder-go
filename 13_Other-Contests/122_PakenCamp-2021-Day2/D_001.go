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
	p := make(map[int]int)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		p[a]++
	}
	x := int(1e9)
	y := 0
	for _, v := range p {
		x = min(x, v)
		y = max(y, v)
	}
	if len(p) < m {
		x = 0
	}
	fmt.Println(x, y)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
