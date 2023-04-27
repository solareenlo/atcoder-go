package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var t int
	fmt.Fscan(in, &t)
	for t > 0 {
		t--
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		if (a&b) == b && ((b&c) == b || (b&c) == 0) && (a|c) == a {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
