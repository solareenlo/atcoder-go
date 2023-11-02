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
	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s[i])
	}
	t := make([]string, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &t[i])
	}

	ans := 0
	for i := 0; i < n; i++ {
		x := false
		for j := 0; j < m; j++ {
			if s[i][len(s[i])-3:] == t[j] {
				x = true
			}
		}
		if x {
			ans++
		}
	}
	fmt.Println(ans)
}
