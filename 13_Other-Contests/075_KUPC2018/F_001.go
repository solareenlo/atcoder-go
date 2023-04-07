package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s[i])
		s[i] *= 2
	}
	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		a--
		b--
		s[a] += c
		s[b] += c
	}
	sort.Ints(s)
	s = reverseOrderInt(s)
	var e [2]int
	for i := 0; i < n; i++ {
		e[i%2] += s[i]
	}
	if e[0] > e[1] {
		fmt.Println("First")
	} else {
		fmt.Println("Second")
	}
}

func reverseOrderInt(a []int) []int {
	n := len(a)
	res := make([]int, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
