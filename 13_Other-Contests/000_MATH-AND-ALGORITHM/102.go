package main

import (
	"bufio"
	"fmt"
	"os"
)

func nCr(n, r int) int {
	if n < 3 && r < 3 {
		if n < r {
			return 0
		}
		if n == 2 && r == 1 {
			return n
		} else {
			return 1
		}
	}
	return (nCr(n/3, r/3) * nCr(n%3, r%3)) % 3
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	var s string
	fmt.Fscan(in, &n, &s)

	mp := make(map[byte]int)
	mp['B'] = 0
	mp['W'] = 1
	mp['R'] = 2

	ans := 0
	for i := 0; i < n; i++ {
		ans += mp[s[i]] * nCr(n-1, i)
		ans %= 3
	}

	if n&1 == 0 {
		ans = (6 - ans) % 3
	}

	anstr := "BWR"
	fmt.Println(string(anstr[ans]))
}
