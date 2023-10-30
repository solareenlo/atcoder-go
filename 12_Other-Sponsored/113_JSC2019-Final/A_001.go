package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var a, b, c [2000020]int

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &b[i])
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if c[a[i]+b[j]] != 0 {
				fmt.Println(i, j, c[a[i]+b[j]]/300000, c[a[i]+b[j]]%300000)
				return
			}
			c[a[i]+b[j]] = i*300000 + j
		}
	}
	fmt.Println(-1)
}
