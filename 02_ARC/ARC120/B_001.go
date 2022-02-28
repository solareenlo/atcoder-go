package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	r := make([]int, 1010)
	b := make([]int, 1010)
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)
		for j := 0; j < m; j++ {
			if s[j] == 'R' {
				r[i+j] = 1
			}
			if s[j] == 'B' {
				b[i+j] = 1
			}
		}
	}

	res := 1
	for i := 0; i < n+m-1; i++ {
		res = res * (2 - r[i] - b[i]) % 998244353
	}
	fmt.Println(res)
}
