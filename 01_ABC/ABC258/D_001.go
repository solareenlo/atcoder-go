package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, x int
	fmt.Fscan(in, &n, &x)

	var a, b int
	c := 0
	l := math.MaxInt64
	for n > 0 && x > 0 {
		n--
		x--
		fmt.Fscan(in, &a, &b)
		c += a + b
		l = min(c+x*b, l)
	}
	fmt.Println(l)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
