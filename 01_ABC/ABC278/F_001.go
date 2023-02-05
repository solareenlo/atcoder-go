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

	var s [16]string
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s[i])
	}

	var f [1 << 16][16]int
	for i := (1 << n) - 1; i > 0; i-- {
		for j := 0; j < n; j++ {
			if ((i >> j) & 1) != 0 {
				f[i][j] = 1
				for k := 0; k < n; k++ {
					if ((i>>k)&1) == 0 && s[j][len(s[j])-1] == s[k][0] {
						f[i][j] &= f[i^(1<<k)][k]
					}
				}
				f[i][j] ^= 1
			}
		}
	}

	a := 1
	for i := 0; i < n; i++ {
		a &= f[1<<i][i]
	}
	if a != 0 {
		fmt.Println("Second")
	} else {
		fmt.Println("First")
	}
}
