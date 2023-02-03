package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 100100

	var n int
	fmt.Fscan(in, &n)
	g, sum := 0, 0
	var a [N]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		g = gcd(g, a[i])
	}
	for i := 1; i <= n; i++ {
		x := a[i] / g
		for x%2 == 0 {
			sum++
			x /= 2
		}
		for x%3 == 0 {
			sum++
			x /= 3
		}
		if x != 1 {
			fmt.Println(-1)
			return
		}
	}
	fmt.Println(sum)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
