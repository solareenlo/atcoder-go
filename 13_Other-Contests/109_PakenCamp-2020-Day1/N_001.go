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

	var a [200005]int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	a[n] = n + 1
	x := n - 1
	for x > 0 && a[x-1] < a[x] {
		x--
	}
	y := 0
	for y < n-1 && a[y] < a[y+1] {
		y++
	}
	if x == 0 {
		fmt.Println(0)
		return
	}
	m := a[0] + a[x]
	for x < n && a[x] < a[y] {
		x++
	}
	t := a[y+1] + a[x-1]
	if m > t {
		m = t
	}
	if a[y] < a[x+1] {
		t = a[y+1] + a[x]
	}
	if m > t {
		m = t
	}
	fmt.Println(m)
}
