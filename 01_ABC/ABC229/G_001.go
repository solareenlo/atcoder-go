package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var s string
	var w int
	fmt.Fscan(in, &s, &w)

	a := make([]int, 1<<20)
	t := 0
	for i := 0; i < len(s); i++ {
		if s[i]&1 != 0 {
			a[t] = i
			t++
		}
	}

	x, l := 0, 0
	for i := 0; i < t; i++ {
		d := (l + i) / 2
		w -= a[i] - a[d] - i + d
		for w < 0 {
			d = (l + i + 1) / 2
			w += a[d] - a[l] - d + l
			l++
		}
		x = max(x, i-l+1)
	}
	fmt.Println(x)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
