package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	b := make([]int, 10)
	f := make([]int, 10)
	for i := 0; i < 10; i++ {
		fmt.Fscan(in, &b[i])
		f[b[i]] = i
	}
	var n int
	fmt.Fscan(in, &n)
	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s[i])
	}

	sort.Slice(s, func(i, j int) bool {
		m1 := len(s[i])
		m2 := len(s[j])
		if m1 != m2 {
			return m1 < m2
		}
		for k := 0; k < m1; k++ {
			if s[i][k] != s[j][k] {
				return f[s[i][k]-'0'] < f[s[j][k]-'0']
			}
		}
		return false
	})

	for i := 0; i < n; i++ {
		fmt.Println(s[i])
	}
}
