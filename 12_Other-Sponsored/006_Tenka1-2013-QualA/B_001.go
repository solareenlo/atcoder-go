package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	s := 0
	for i := 1; i <= n; i++ {
		var a, b, c, d, e int
		fmt.Fscan(in, &a, &b, &c, &d, &e)
		if a+b+c+d+e < 20 {
			s++
		}
	}
	fmt.Println(s)
}
