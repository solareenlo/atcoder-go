package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var x [101]string
	var c [101]int

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &x[i])
	}
	s := make(map[string]int)
	for i := 1; i <= m; i++ {
		var y string
		fmt.Fscan(in, &y)
		s[y] = i
	}
	for i := 0; i <= m; i++ {
		fmt.Fscan(in, &c[i])
	}
	ans := 0
	for i := 1; i <= n; i++ {
		ans += c[s[x[i]]]
	}
	fmt.Println(ans)
}
