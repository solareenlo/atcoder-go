package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const mod = 1000000007

	var q int
	fmt.Fscan(in, &q)

	for i := 0; i < q; i++ {
		var n int
		fmt.Fscan(in, &n)
		if n > 1 {
			tmp := 0
			if n%3 == 1 {
				tmp = 1
			}
			c := n/3 - tmp
			if n%3 != 0 {
				if n%3 == 2 {
					n = 2
				} else {
					n = 4
				}
			} else {
				n = 1
			}
			for b := 3; c > 0; c, b = c/2, b*b%mod {
				if (c & 1) != 0 {
					n = n * b % mod
				}
			}
		}
		if i < q-1 {
			fmt.Printf("%d ", n)
		} else {
			fmt.Println(n)
		}
	}
}
