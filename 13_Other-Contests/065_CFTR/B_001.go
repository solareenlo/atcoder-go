package main

import (
	"bufio"
	"fmt"
	"os"
)

var n int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var q int
	fmt.Scan(&n, &q)

	for i := 0; i < q; i++ {
		var v, w int
		fmt.Fscan(in, &v, &w)
		fmt.Fprintln(out, f(v, w))
	}
}

func f(v, w int) int {
	if v == w {
		return v
	}
	if v < w {
		return f(w, v)
	}
	if n == 1 {
		return w
	}
	return f((v+n-2)/n, w)
}
