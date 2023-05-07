package main

import (
	"bufio"
	"fmt"
	"os"
)

const p = 1000000007

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)
	k %= p

	ans := 0
	if n == 1 {
		ans = k % p
	} else {
		ans = k * (k - 1)
		ans %= p
		ans = (ans * f(max(k-2, 0), n-2)) % p
	}
	fmt.Println(ans)
}

func f(a, b int) int {
	if b == 0 {
		return 1
	}
	if b%2 != 0 {
		return ((a % p) * (f(a, b-1) % p)) % p
	}
	res := f(a, b/2) % p
	return (res * res) % p
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
