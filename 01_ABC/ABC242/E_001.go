package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	const mod = 998244353
	for j := 0; j < t; j++ {
		var n int
		var s string
		fmt.Fscan(in, &n, &s)
		flag := false
		for i := n/2 - 1; i >= 0; i-- {
			if s[i] != s[n-i-1] {
				flag = s[i] > s[n-i-1]
				break
			}
		}
		res := 0
		for i := 0; i+i < n; i++ {
			res *= 26
			res += int(s[i] - 'A')
			res %= mod
		}
		if !flag {
			res++
		}
		fmt.Fprintln(out, res%mod)
	}
}
