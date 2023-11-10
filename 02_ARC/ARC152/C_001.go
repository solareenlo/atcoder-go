package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var a [200020]int

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	d := a[n] - a[1]
	g := d
	for i := 2; i <= n; i++ {
		g = gcd(g, 2*(a[i]-a[1]))
	}
	fmt.Println(a[1]%g + d)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
