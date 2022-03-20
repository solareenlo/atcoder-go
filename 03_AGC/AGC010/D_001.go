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

	y := 0
	x := 0
	i := 0
	for ; ; y ^= 1 {
		for i, x = 1, 0; i <= n; i++ {
			x += a[i] & 1
		}
		if x > 1 || ((n^x)&1) != 0 {
			if (n^x^y)&1 != 0 {
				fmt.Println("First")
			} else {
				fmt.Println("Second")
			}
			return
		}
		for i = 1; i <= n; i++ {
			if a[i]&1 != 0 {
				break
			}
		}
		if a[i] == 1 {
			if y != 0 {
				fmt.Println("First")
			} else {
				fmt.Println("Second")
			}
			return
		} else {
			a[i]--
		}
		for i, x = 1, 0; i <= n; i++ {
			x = gcd(x, a[i])
		}
		for i = 1; i <= n; i++ {
			a[i] /= x
		}
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
