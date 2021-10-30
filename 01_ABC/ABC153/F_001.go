package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, d, a int
	fmt.Fscan(in, &n, &d, &a)

	type node struct{ x, y int }
	b := make([]node, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &b[i].x, &b[i].y)
	}
	sort.Slice(b, func(i, j int) bool {
		return b[i].x < b[j].x
	})

	c := [200006]int{}
	res := 0
	for i, j := 0, 0; i < n; i++ {
		for j < n && b[j].x <= b[i].x+2*d {
			j++
		}
		need := max((b[i].y-c[i]*a+a-1)/a, 0)
		res += need
		c[i] += need
		c[j] -= need
		c[i+1] += c[i]
	}

	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
