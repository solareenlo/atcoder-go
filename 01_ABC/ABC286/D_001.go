package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var f [10005]int

	var n, m int
	fmt.Fscan(in, &n, &m)
	f[0] = 1
	for i := 1; i <= n; i++ {
		var v, c int
		fmt.Fscan(in, &v, &c)
		for c > 0 {
			c--
			for j := m; j >= v; j-- {
				f[j] |= f[j-v]
			}
		}
	}
	if f[m] != 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
