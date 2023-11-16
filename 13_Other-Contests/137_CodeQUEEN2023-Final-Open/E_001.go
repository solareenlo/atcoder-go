package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(1e18)

	var n int
	fmt.Fscan(in, &n)
	x, y, z := 0, -INF, -INF
	for n > 0 {
		n--
		var a int
		fmt.Fscan(in, &a)
		m := max(x, y+a, z-a)
		y = max(y, x-a)
		z = max(z, x+a)
		x = m
	}
	fmt.Println(x)
}

func max(a ...int) int {
	res := a[0]
	for i := range a {
		if res < a[i] {
			res = a[i]
		}
	}
	return res
}
