package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var l uint
	fmt.Scan(&l)

	n := 0
	for ; (1 << n) <= l; n++ {
	}
	n--

	fmt.Println(n+1, 2*n+bits.OnesCount(l)-1)

	for i := 1; i <= n; i++ {
		fmt.Println(i, i+1, 0)
		fmt.Println(i, i+1, 1<<(i-1))
	}

	for i := 0; i < n; i++ {
		if (l>>i)&1 != 0 {
			fmt.Println(i+1, n+1, l-(l&((1<<(i+1))-1)))
		}
	}
}
