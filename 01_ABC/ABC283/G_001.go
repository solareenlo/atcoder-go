package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var a [62]int

	var n, L, R int
	fmt.Fscan(in, &n, &L, &R)
	for i := 1; i <= n; i++ {
		var x int
		fmt.Fscan(in, &x)
		for i := 59; i >= 0; i-- {
			if ((x >> i) & 1) != 0 {
				if a[i] == 0 {
					a[i] = x
					break
				}
				x ^= a[i]
			}
		}
	}
	p := make([]int, 0)
	P := 0
	for i := 0; i < 60; i++ {
		if a[i] != 0 {
			p = append(p, i)
			P++
		}
	}
	for i := L - 1; i < R; i++ {
		val := 0
		for j := P - 1; j >= 0; j-- {
			if (((i >> j) & 1) ^ ((val >> p[j]) & 1)) != 0 {
				val ^= a[p[j]]
			}
		}
		fmt.Printf("%d ", val)
	}
}
