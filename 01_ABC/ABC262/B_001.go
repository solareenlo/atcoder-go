package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	var s, S [5000]int
	var a [5000][5000]bool
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &s[i], &S[i])
		a[s[i]][S[i]] = true
		a[S[i]][s[i]] = true
	}

	ans := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if a[s[i]][j] && a[S[i]][j] {
				ans++
			}
		}
	}
	fmt.Println(ans / 3)
}
