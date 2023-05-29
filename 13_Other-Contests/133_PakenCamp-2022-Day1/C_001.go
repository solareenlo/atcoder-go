package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const mod = 998244353

	var s string
	fmt.Fscan(in, &s)
	n := len(s)
	ans := 1
	for i := 0; i < (n+1)/2; i++ {
		if s[i] == '?' && s[n-1-i] == '?' {
			ans *= 26
			ans += mod
			ans %= mod
		} else if s[i] != '?' && s[n-1-i] != '?' && s[i] != s[n-1-i] {
			ans = 0
		}
	}
	fmt.Println(ans)
}
