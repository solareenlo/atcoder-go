package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, x int
	var s string
	fmt.Fscan(in, &n, &x, &s)

	const maxn = 1000000000000000000
	cnt := 0
	for i := 0; i < n; i++ {
		if s[i] == 'U' {
			if cnt != 0 {
				cnt--
			} else {
				x /= 2
			}
		} else {
			if x*2 > maxn {
				cnt++
			} else {
				if s[i] == 'L' {
					x *= 2
				} else {
					x = x*2 + 1
				}
			}
		}
	}
	fmt.Println(x)
}
