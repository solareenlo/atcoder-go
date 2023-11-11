package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for t > 0 {
		t--
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		if (a%2) != (b%2) || (b%2) != (c%2) || (a+b+c)%3 != 0 {
			fmt.Fprintln(out, -1)
			continue
		}
		x := (a + b + c) / 3
		fmt.Fprintln(out, (abs(a-x)+abs(b-x)+abs(c-x))/4)
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
