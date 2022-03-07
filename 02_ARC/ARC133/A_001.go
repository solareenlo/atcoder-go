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
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	x := 0
	for i := 1; i <= n; i++ {
		if i == n || a[i] > a[i+1] {
			x = a[i]
			break
		}
	}

	for i := 1; i <= n; i++ {
		if a[i] != x {
			fmt.Print(a[i], " ")
		}
	}
}
