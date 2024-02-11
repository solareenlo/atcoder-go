package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 998244353

	var s string
	fmt.Fscan(in, &s)

	ans := 1
	cur := 0
	for _, c := range s {
		if c == '*' {
			ans = ans * cur % MOD
			cur = 0
		} else {
			cur = (cur*10 + int(c-'0')) % MOD
		}
	}
	fmt.Println(ans * cur % MOD)
}
