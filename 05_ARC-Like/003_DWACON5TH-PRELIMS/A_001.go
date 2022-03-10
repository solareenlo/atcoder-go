package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n+1)
	s := 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		s += a[i]
	}

	id := 0
	mn := 1 << 60
	for i := 1; i <= n; i++ {
		if abs(a[i]*n-s) < mn {
			mn = abs(a[i]*n - s)
			id = i
		}
	}
	fmt.Println(id - 1)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
