package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var a [2009]int

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for n > 1 {
		for i := 1; i <= n; i++ {
			if n%2 == 0 {
				a[i] = max(a[i], a[i+1])
			} else {
				a[i] = min(a[i], a[i+1])
			}
		}
		n--
	}
	fmt.Println(a[1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
