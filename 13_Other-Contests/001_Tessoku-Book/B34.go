package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, x, y int
	fmt.Fscan(in, &n, &x, &y)
	a := []int{0, 0, 1, 1, 2}
	c := 0
	for i := 0; i < n; i++ {
		var w int
		fmt.Fscan(in, &w)
		c ^= a[w%5]
	}

	if c > 0 {
		fmt.Println("First")
	} else {
		fmt.Println("Second")
	}
}
