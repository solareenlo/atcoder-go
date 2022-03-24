package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var a, b, c int
	fmt.Fscan(in, &a, &b, &c)

	d := c
	a--
	for i := 0; i < a; i++ {
		var e int
		fmt.Fscan(in, &e)
		c = gcd(e, c)
		d = max(e, d)
	}

	if b%c == 0 && b <= d {
		fmt.Println("POSSIBLE")
	} else {
		fmt.Println("IMPOSSIBLE")
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
