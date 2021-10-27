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

	mod := int(1e9 + 7)
	res := 0
	for i := 0; i < 60; i++ {
		one := 0
		for j := 0; j < n; j++ {
			if a[j]>>i&1 != 0 {
				one++
			}
		}
		zero := n - one
		now := one * zero % mod
		for j := 0; j < i; j++ {
			now = now * 2 % mod
		}
		res += now
		res %= mod
	}

	fmt.Println(res)
}
