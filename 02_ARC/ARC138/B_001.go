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
	r := n
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	l := 1
	flip := 0
	for l <= r {
		if a[r]^flip == 0 {
			r--
		} else if a[l]^flip == 0 {
			l++
			flip ^= 1
		} else {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
