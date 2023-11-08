package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var d, s [102]int

	var m int
	fmt.Fscan(in, &m)
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &d[i])
		s[i] = d[i] + s[i-1]
	}
	mid := s[m] / 2
	i := m
	for s[i] > mid {
		i--
	}
	fmt.Println(i+1, mid-s[i]+1)
}
