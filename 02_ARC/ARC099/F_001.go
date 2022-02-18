package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAXN = 250010
const P = 233
const iP = 7204522363551799129

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	var s string
	fmt.Fscan(in, &n, &s)
	s = " " + s

	hsh1 := make([]int, MAXN)
	hsh2 := make([]int, MAXN)
	hsh2[0] = 1
	mp := map[int]int{}
	for i := 1; i <= n; i++ {
		hsh1[i] = hsh1[i-1]
		hsh2[i] = hsh2[i-1]
		if s[i] == '+' {
			hsh1[i] += hsh2[i]
		} else if s[i] == '-' {
			hsh1[i] -= hsh2[i]
		} else if s[i] == '<' {
			hsh2[i] *= iP
		} else {
			hsh2[i] *= P
		}
		mp[hsh1[i]]++
	}

	ans := 0
	for i := 1; i <= n; i++ {
		ans += mp[hsh1[n]*hsh2[i-1]+hsh1[i-1]]
		mp[hsh1[i]]--
	}
	fmt.Println(ans)
}
