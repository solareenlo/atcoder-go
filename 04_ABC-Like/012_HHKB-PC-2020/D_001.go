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

	mod := int(1e9 + 7)
	for i := 0; i < t; i++ {
		var n, a, b int
		fmt.Fscan(in, &n, &a, &b)
		c := n - a + 1
		d := n - b + 1
		e := n - a - b
		f := (e + 2) * (e + 1) / 2 % mod
		if e < 0 {
			f = 0
		}
		res := 4 * f * (c*d%mod - f) % mod
		fmt.Fprintln(out, (res+mod)%mod)
	}
}
