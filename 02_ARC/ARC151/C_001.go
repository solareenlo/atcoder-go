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
	to := -1
	tx := 1
	ans := 0
	for i := 1; i <= m+1; i++ {
		var x, o int
		if i <= m {
			fmt.Fscan(in, &x, &o)
		} else {
			x = n
			o = -1
		}
		if to < 0 || o < 0 {
			ans ^= x - tx
		}
		if o == to {
			ans ^= 1
		} else {
			ans ^= 0
		}
		tx = x
		to = o
	}
	var tmp int
	if m != 0 {
		tmp = ans
	} else {
		tmp = n & 1
	}
	if tmp != 0 {
		fmt.Println("Takahashi")
	} else {
		fmt.Println("Aoki")
	}
}
