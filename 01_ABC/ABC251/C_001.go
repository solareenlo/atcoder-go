package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	m := make(map[string]bool)
	v := 0
	p := 0
	for i := 1; i <= n; i++ {
		var s string
		var t int
		fmt.Fscan(in, &s, &t)
		if _, ok := m[s]; !ok && t > v {
			v = t
			p = i
		}
		m[s] = true
	}
	fmt.Println(p)
}
