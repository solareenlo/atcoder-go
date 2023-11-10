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

	var n, K int
	fmt.Fscan(in, &n, &K)
	if n%2 == 0 {
		fmt.Fprintln(out, -1)
		return
	}
	fmt.Fprintln(out, n/2)
	m := gcd(n, K)
	var work func(int, int)
	work = func(x, y int) {
		x %= n
		for ; x != y; x = (x + K*2) % n {
			fmt.Fprintln(out, x, (x+K)%n)
		}
	}
	for i := 0; i < m; i++ {
		if (i & 1) != 0 {
			work(i+K*3, i)
			fmt.Fprintln(out, (i+K)%n, i+1)
		} else {
			work(i+K, i)
			if i != m-1 {
				fmt.Fprintln(out, i, (i+1+K*2)%n)
			}
		}
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
