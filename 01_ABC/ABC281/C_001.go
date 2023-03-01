package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, t int
	fmt.Fscan(in, &n, &t)
	a := make([]int, n+1)
	m := 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		m += a[i]
	}
	t %= m
	l := 1
	for t-a[l] > 0 {
		t -= a[l]
		l++
	}
	fmt.Println(l, t)
}
