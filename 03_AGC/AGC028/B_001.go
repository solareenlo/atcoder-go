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

	const mod = 1_000_000_007
	INV := make([]int, 100100)
	INV[1] = 1
	for i := 2; i <= n; i++ {
		INV[i] = INV[mod%i] * (mod - mod/i) % mod
	}
	for i := 1; i <= n; i++ {
		INV[i] += INV[i-1]
		INV[i] %= mod
	}

	res := 0
	for i := 1; i <= n; i++ {
		var x int
		fmt.Fscan(in, &x)
		res = x*(INV[n-i+1]+INV[i]-1) + res
		res %= mod
	}

	for i := 1; i <= n; i++ {
		res = res * i % mod
	}
	fmt.Println(res)
}
