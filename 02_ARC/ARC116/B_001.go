package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a)

	const mod = 998244353
	res := 0
	pre := 0
	for i := 0; i < n; i++ {
		res += ((pre + a[i]) % mod) * a[i] % mod
		res %= mod
		pre = (pre*2%mod + a[i]) % mod
	}
	fmt.Println(res)
}
