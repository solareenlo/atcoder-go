package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, d int
	fmt.Fscan(in, &n, &d)
	a := make([]int, 1<<n)
	c := make([]int, 1<<n)
	c[0] = 1
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[1<<i])
	}
	r := 1 << d
	for i := 1; i < 1<<n; i++ {
		a[i] = a[i&(i-1)] | a[i&(-i)]
		c[i] = -c[i&(i-1)]
		r += c[i] * (1 << (d - popcount(uint64(a[i]))))
	}
	fmt.Println(r)
}

func popcount(n uint64) int {
	return bits.OnesCount64(n)
}
