package main

import (
	"bufio"
	"fmt"
	"os"
)

func fx(x int) int {
	if x > 2 {
		x -= 3
	}
	return x
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s string
	var q int
	fmt.Fscan(in, &s, &q)
	s = " " + s

	for i := 0; i < q; i++ {
		var t, x int
		fmt.Fscan(in, &t, &x)
		val := 0
		for t != 0 && x > 1 {
			t--
			val += 2 - (x & 1)
			val = fx(val)
			x = (x + 1) >> 1
		}
		fmt.Fprintln(out, string('A'+(int(s[x]-'A')+val+t)%3))
	}
}
