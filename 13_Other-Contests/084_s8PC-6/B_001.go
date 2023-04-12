package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i], &b[i])
	}
	sort.Ints(a)
	sort.Ints(b)
	s := 0
	for i := 0; i < n; i++ {
		s += abs(a[i]-a[n/2]) + b[i] - a[i] + abs(b[i]-b[n/2])
	}
	fmt.Println(s)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
