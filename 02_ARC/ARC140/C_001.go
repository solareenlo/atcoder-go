package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)

	defer out.Flush()
	var n, x int
	fmt.Scan(&n, &x)

	y := 0
	if x > (n+1)/2 {
		x = n - x + 1
		y = 1
	}

	var a [200002]int
	a[n] = n
	a[1] = x
	b := -1
	c := n - 1
	for i := n - 1; i > 1; i-- {
		if a[i+1]+b*c == a[1] {
			c = c - 1
		}
		a[i] = a[i+1] + b*c
		b = -b
		c = c - 1
	}
	for i := 1; i <= n; i++ {
		if y == 0 {
			fmt.Fprintf(out, "%d ", a[i])
		} else {
			fmt.Fprintf(out, "%d ", n+1-a[i])
		}
	}
}
