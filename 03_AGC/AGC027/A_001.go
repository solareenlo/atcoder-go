package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, x int
	fmt.Fscan(in, &n, &x)

	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a)

	nc := 0
	for nc < n && x >= a[nc] {
		x -= a[nc]
		nc++
	}

	if x > 0 && nc == n {
		nc--
	}
	fmt.Println(nc)
}
