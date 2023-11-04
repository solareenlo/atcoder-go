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

	var n int
	fmt.Fscan(in, &n)
	l, r := 1, n
	for l+1 < r {
		mid := (l + r) >> 1
		fmt.Fprintln(out, "?", mid)
		out.Flush()
		var x int
		fmt.Fscan(in, &x)
		if x != 0 {
			r = mid
		} else {
			l = mid
		}
	}
	fmt.Fprintln(out, "!", l)
	out.Flush()
}
