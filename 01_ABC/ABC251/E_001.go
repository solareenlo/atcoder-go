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

	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	d := make([]int, n+1)
	e := make([]int, n+1)
	d[1] = 1 << 60
	e[0] = 1 << 60
	e[1] = a[n-1]
	for i := 0; i < n-1; i++ {
		d[i+2] = min(d[i], d[i+1]) + a[i]
		e[i+2] = min(e[i], e[i+1]) + a[i]
	}
	fmt.Println(min(d[n], e[n], e[n-1]))
}

func min(a ...int) int {
	res := a[0]
	for i := range a {
		if res > a[i] {
			res = a[i]
		}
	}
	return res
}
