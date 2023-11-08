package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 1000005

	var t [N]int

	var n, m int
	var S string
	fmt.Fscan(in, &n, &S, &m)
	S = " " + S
	s := strings.Split(S, "")
	tim, tt := 0, 0
	for i := 1; i <= m; i++ {
		var op, x int
		var y string
		fmt.Fscan(in, &op, &x, &y)
		if op == 1 {
			s[x] = string(y[0])
			t[x] = i
		} else if op == 2 {
			tim = i
			tt = 1
		} else {
			tim = i
			tt = 0
		}
	}
	for i := 1; i <= n; i++ {
		if t[i] < tim {
			if tt != 0 {
				fmt.Print(strings.ToLower(string(s[i])))
			} else {
				fmt.Print(strings.ToUpper(string(s[i])))
			}
		} else {
			fmt.Print(string(s[i]))
		}
	}
}
