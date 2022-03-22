package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, a int
	fmt.Fscan(in, &n, &a)

	x := 0
	c := 0
	for i := 1; i < n; i++ {
		var b int
		fmt.Fscan(in, &b)
		if b == a {
			continue
		}
		if x != 0 {
			if x == 1 && b < a || x == -1 && b > a {
				x = 0
				c++
			}
		} else {
			if b > a {
				x = 1
			} else {
				x = -1
			}
		}
		a = b
	}
	fmt.Println(c + 1)
}
