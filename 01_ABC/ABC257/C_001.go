package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type pair struct {
	first, second int
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	var str string
	fmt.Scan(&n, &str)

	a := make([]pair, 8<<15)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i].first)
		a[i].second = int('1' - str[i-1])
	}
	tmp := a[1 : n+1]
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].first == tmp[j].first {
			return tmp[i].second < tmp[j].second
		}
		return tmp[i].first < tmp[j].first
	})

	s := make([]int, 8<<15)
	for i := 1; i <= n; i++ {
		s[i] += s[i-1] + 1 - a[i].second
	}

	m := 0
	for i := 0; i <= n; i++ {
		m = max(m, i-s[i]+s[n]-s[i])
	}
	fmt.Println(m)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
