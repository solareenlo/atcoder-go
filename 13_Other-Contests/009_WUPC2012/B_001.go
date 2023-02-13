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
	s := make([]string, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &s[i])
	}

	ans := s[1] + s[2]
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i != j && ans > s[i]+s[j] {
				ans = s[i] + s[j]
			}
		}
	}
	fmt.Println(ans)
}
