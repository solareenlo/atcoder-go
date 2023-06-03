package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, l, k int
	fmt.Fscan(in, &n, &l, &k)
	A := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &A[i])
	}
	y := 1
	z := l + 1
	for y < z {
		x := (y + z + 1) / 2
		c := 0
		a := 0
		for i := 0; i < n; i++ {
			if A[i]-a >= x && l-A[i] >= x {
				c++
				a = A[i]
			}
		}
		if c >= k {
			y = x
		} else {
			z = x - 1
		}
	}
	fmt.Println(y)
}
