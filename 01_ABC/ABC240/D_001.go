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

	a := make([]int, 200002)
	b := make([]int, 200002)
	t := 0
	for i := 1; i <= n; i++ {
		var x int
		fmt.Fscan(in, &x)
		t++
		a[t] = x
		if t == 1 || a[t] != a[t-1] {
			b[t] = 1
		} else {
			b[t] = b[t-1] + 1
		}
		if b[t] == x {
			t = t - x
		}
		fmt.Fprintln(out, t)
	}
}
