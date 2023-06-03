package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var D int
	fmt.Fscan(in, &D)
	a := make([]int, D)
	fmt.Fscan(in, &a[0])
	for i := 1; i < D; i++ {
		var v int
		fmt.Fscan(in, &v)
		a[i] = a[i-1] + v
	}
	var Q int
	fmt.Fscan(in, &Q)
	for Q > 0 {
		Q--
		var l, r int
		fmt.Fscan(in, &l, &r)
		l--
		r--
		if a[l] == a[r] {
			fmt.Println("Same")
		} else {
			if a[l] > a[r] {
				fmt.Println(l + 1)
			} else {
				fmt.Println(r + 1)
			}
		}
	}
}
