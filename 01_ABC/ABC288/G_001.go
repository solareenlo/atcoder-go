package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var a [531441]int

	var n int
	fmt.Fscan(in, &n)
	m := pow(3, n)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &a[i])
	}
	for k, t := 0, 1; k < n; k, t = k+1, t*3 {
		for i := 0; i < m; i++ {
			if i/t%3 == 0 {
				x := a[i]
				y := a[i+t]
				z := a[i+(t<<1)]
				a[i] = y - z
				a[i+t] = x + z - y
				a[i+(t<<1)] = y - x
			}
		}
	}
	for i := 0; i < m; i++ {
		fmt.Fprintf(out, "%d ", a[i])
	}
	fmt.Fprintln(out)
}

func pow(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a
		}
		a = a * a
		n /= 2
	}
	return res
}
