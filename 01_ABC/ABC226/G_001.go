package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	a = make([]int, 6)
	b = make([]int, 6)
)

func pack(x, y int) {
	c := min(a[x], b[y])
	a[x] -= c
	b[y] -= c
	b[y-x] += c
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for tt := 0; tt < t; tt++ {
		for i := 0; i < 5; i++ {
			fmt.Fscan(in, &a[i+1])
		}
		for i := 0; i < 5; i++ {
			fmt.Fscan(in, &b[i+1])
		}
		a[0] = 0
		b[0] = 0

		pack(5, 5)
		pack(4, 4)
		pack(4, 5)
		pack(3, 3)
		pack(3, 5)
		pack(3, 4)
		for i := 0; i < 4; i++ {
			pack(2, 5-i)
		}
		for i := 0; i < 5; i++ {
			pack(1, 5-i)
		}

		ok := true
		for i := 0; i < 5; i++ {
			if a[i+1] > 0 {
				ok = false
			}
		}

		if ok {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}
}
