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
	x := make([]int, n)
	y := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i], &y[i])
	}
	var s string
	fmt.Fscan(in, &s)

	sa := make(map[int]int)
	for i := 0; i < n; i++ {
		if s[i] == 'L' {
			if sa[y[i]] < x[i] {
				sa[y[i]] = x[i]
			}
		}
	}

	for i := 0; i < n; i++ {
		if s[i] == 'R' {
			if sa[y[i]] > x[i] {
				fmt.Println("Yes")
				return
			}
		}
	}

	fmt.Println("No")
}
