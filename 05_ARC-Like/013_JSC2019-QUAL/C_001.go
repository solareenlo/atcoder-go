package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	var s string
	fmt.Fscan(in, &n, &s)
	s = " " + s

	const mod = 1_000_000_007
	ans := 1
	for i := 1; i <= n; i++ {
		ans = ans * i % mod
	}

	z := 0
	flag := 0
	for i := 1; i <= n*2; i++ {
		tmp := 0
		if s[i] == 'B' {
			tmp = 1
		}
		if tmp == flag {
			if z == 0 {
				ans = 0
			}
			ans = ans * z % mod
			z--
			flag ^= 1
		} else {
			z++
			flag ^= 1
		}
	}

	if z != 0 {
		ans = 0
	}
	fmt.Println(ans)
}
