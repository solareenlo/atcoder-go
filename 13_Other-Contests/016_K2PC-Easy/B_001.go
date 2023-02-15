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
	a := make([]int, 10)
	for i := 0; i < 10; i++ {
		fmt.Fscan(in, &a[i])
	}
	b := make([]int, 7)
	c := make([]int, 7)
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)
		for j := 0; j < 7; j++ {
			if s[j] == 'X' {
				c[j]++
			} else {
				c[j] = 0
			}
			b[j] = max(b[j], c[j])
		}
	}
	sort.Ints(a)
	sort.Ints(b)
	ok := 1
	for i := 0; i < 7; i++ {
		if b[i] <= a[i+3] {
			ok &= 1
		} else {
			ok &= 0
		}
	}
	if ok != 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
