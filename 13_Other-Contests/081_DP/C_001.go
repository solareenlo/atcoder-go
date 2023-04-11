package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, a, b, c int
	fmt.Fscan(in, &n, &a, &b, &c)
	for n > 0 {
		n--
		var x, y, z int
		fmt.Fscan(in, &x, &y, &z)
		x += max(b, c)
		y += max(a, c)
		z += max(a, b)
		a = x
		b = y
		c = z
	}
	fmt.Println(max(a, b, c))
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
