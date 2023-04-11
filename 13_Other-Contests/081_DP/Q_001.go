package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 200009

	var n int
	fmt.Fscan(in, &n)
	var h [N]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &h[i])
	}
	var t [N]int
	z := 0
	for i := 1; i <= n; i++ {
		y := 0
		for x := h[i]; x > 0; x -= x & -x {
			y = max(y, t[x])
		}
		var x int
		fmt.Fscan(in, &x)
		y += x
		z = max(z, y)
		for x := h[i]; x <= n; x += x & -x {
			t[x] = max(t[x], y)
		}
	}
	fmt.Println(z)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
