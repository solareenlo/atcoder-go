package main

import (
	"bufio"
	"fmt"
	"os"
)

var dp = [3][3]int{{1, 0, 0}, {1, 1, 0}, {1, 2, 1}}

func nCr(n, r int) int {
	ret := 1
	for n > 0 {
		ret = (ret * dp[n%3][r%3]) % 3
		n /= 3
		r /= 3
	}
	return ret
}

func main() {
	in := bufio.NewReader(os.Stdin)

	m := map[byte]int{}
	m['W'] = 1
	m['R'] = 2

	var n int
	var s string
	fmt.Fscan(in, &n, &s)

	ans := 0
	for i := 0; i < n; i++ {
		ans = (ans + m[s[i]]*nCr(n-1, i)) % 3
	}
	if n%2 == 0 {
		ans = (ans * 2) % 3
	}

	fmt.Println(string("BWR"[ans]))
}
