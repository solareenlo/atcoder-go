package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var t, n int
	fmt.Fscan(in, &t, &n)
	var a [500001]int
	for i := 0; i < n; i++ {
		var l, r int
		fmt.Fscan(in, &l, &r)
		a[l]++
		a[r]--
	}
	z := 0
	for i := 0; i < t; i++ {
		fmt.Println(z + a[i])
		z += a[i]
	}
}
