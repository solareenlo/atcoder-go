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
	var a [2][100005]int
	for k := 0; k < 2; k++ {
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[k][i])
		}
		u := 0
		var t int
		for j := 0; j < 60; j++ {
			for t = j - u; t < n && (^a[k][t]>>j)&1 != 0; t++ {

			}
			if t == n {
				u++
				continue
			}
			a[k][t], a[k][j-u] = a[k][j-u], a[k][t]
			for i := 0; i < n; i++ {
				if i != j-u && (a[k][i]>>j)&1 != 0 {
					a[k][i] ^= a[k][j-u]
				}
			}
		}
	}
	ok := true
	for i := 0; i < n; i++ {
		if a[0][i] != a[1][i] {
			ok = false
		}
	}
	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
