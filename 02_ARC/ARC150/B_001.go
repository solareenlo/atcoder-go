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
		var a, b int
		fmt.Fscan(in, &a, &b)
		ans := abs(a - b)
		for l := a; l < b; {
			r := (b + l - 1) / l
			ans = min(ans, l*r-a-b+l)
			l = (b-1)/(r-1) + 1
		}
		fmt.Fprintln(out, ans)
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
