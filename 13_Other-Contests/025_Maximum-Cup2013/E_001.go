package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, t int
	fmt.Fscan(in, &n, &t)
	var find func(string, string) bool
	find = func(a, b string) bool {
		return strings.Index(a, b) != -1
	}
	s1, s2 := 0, 0
	for i := 1; i <= n; i++ {
		var s, c string
		var r int
		fmt.Fscan(in, &s, &c, &r)
		if c[0] != 'N' && find(s, "Alicia") {
			s1 += r
		}
		s2 += r
	}
	res := float64(t-t/10) * float64(s1) / float64(s2)
	fmt.Println(res)
}
