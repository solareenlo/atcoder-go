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
	c := make([]int, 1000005)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		c[a[i]]++
	}

	const mx = 1000000
	for x := 1; x < mx; x *= 10 {
		for i := 0; i < mx; i++ {
			if i/x%10 != 0 {
				c[i] += c[i-x]
			}
		}
	}

	s := 0
	for i := 1; i <= n; i++ {
		s += c[999999-a[i]]
		y := a[i]
		f := 1
		for y > 0 {
			if y%10 > 4 {
				f = 0
			}
			y /= 10
		}
		s -= f
	}
	fmt.Println(s / 2)
}
