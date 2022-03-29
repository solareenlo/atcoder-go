package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var h, w, n int
	fmt.Fscan(in, &h, &w, &n)

	a := make([]int, h+1)
	for i := 0; i < h; i++ {
		a[i] = w
	}

	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		a[x-1] = min(a[x-1], y-1)
	}

	u := 0
	for i := 1; i <= h; i++ {
		if u >= a[i] {
			fmt.Println(i)
			break
		} else {
			if u+1 != a[i] {
				u++
			}
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
