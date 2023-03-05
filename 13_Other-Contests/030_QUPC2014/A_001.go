package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var a, b, c, d int
	fmt.Fscan(in, &a, &b, &c, &d)

	s := make([][]int, c)
	for i := range s {
		s[i] = make([]int, a)
	}
	for i := 0; i < c; i++ {
		for j := 0; j < a; j++ {
			fmt.Fscan(in, &s[i][j])
		}
		sort.Ints(s[i])
	}

	now := make([]int, c)
	for i := 0; i < c; i++ {
		now[i] = s[i][a-b]
	}
	sort.Ints(now)
	fmt.Println(now[c-d])
}
