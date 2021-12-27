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

	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	c := make([]int, n+1)
	c[0] = 3
	res := 1
	mod := int(1e9 + 7)
	for i := 0; i < n; i++ {
		res *= c[a[i]]
		res %= mod
		c[a[i]]--
		c[a[i]+1]++
	}
	fmt.Println(res)
}
