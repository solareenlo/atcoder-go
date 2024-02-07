package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	var s string
	fmt.Fscan(in, &n, &s)
	m := len(s)
	var a, b [500500]int
	for i := 0; i < n; i++ {
		var t string
		fmt.Fscan(in, &t)
		k := len(t)
		g := 0
		for j := 0; j < k; j++ {
			if g < m && s[g] == t[j] {
				g++
			}
		}
		a[g]++
		g = 0
		for j := 0; j < k; j++ {
			if g < m && s[m-g-1] == t[k-j-1] {
				g++
			}
		}
		b[i] = g
	}
	for i := m; i >= 0; i-- {
		a[i] += a[i+1]
	}
	k := 0
	for i := 0; i < n; i++ {
		k += a[m-b[i]]
	}
	fmt.Println(k)
}
