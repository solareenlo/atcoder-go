package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, p, q, r, s int
	fmt.Fscan(in, &n, &p, &q, &r, &s)
	var a [101]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := p; i <= q; i++ {
		a[i], a[r-p+i] = a[r-p+i], a[i]
	}
	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", a[i])
	}
	fmt.Println()
}
