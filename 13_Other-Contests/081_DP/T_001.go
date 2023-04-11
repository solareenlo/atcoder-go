package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const mod = 1000000007
	const N = 3145

	var n int
	var s string
	fmt.Fscan(in, &n, &s)
	s = " " + s
	var f [N]int
	for i := 1; i < n+1; i++ {
		f[i] = i
	}
	for i := 1; i < n; i++ {
		if s[i] == '<' {
			for j := 1; j < n; j++ {
				f[j] = (f[n-i+1] - f[j] + mod) % mod
			}
		}
		for j := 1; j < n+1; j++ {
			f[j+1] = (f[j+1] + f[j]) % mod
		}
	}
	fmt.Println(f[1])
}
