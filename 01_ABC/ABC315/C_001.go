package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 300009

	var f, s [N]int

	var n int
	fmt.Fscan(in, &n)
	m := 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &f[i], &s[i])
		if s[i] > s[m] {
			m = i
		}
	}
	mt := 0
	for i := 1; i <= n; i++ {
		if f[i] == f[m] {
			s[i] /= 2
		}
		if i != m && s[mt] < s[i] {
			mt = i
		}
	}
	fmt.Println(s[m]*2 + s[mt])
}
