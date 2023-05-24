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
	var s, t string
	fmt.Fscan(in, &s, &t)

	ss := make([]int, 0)
	tt := make([]int, 0)
	for i := 0; i < n; i++ {
		if s[i] == '0' {
			ss = append(ss, i%2)
		}
		if t[i] == '0' {
			tt = append(tt, i%2)
		}
	}

	if len(ss) != len(tt) || count(ss, 1) != count(tt, 1) {
		fmt.Println(-1)
		return
	}

	i, j := 0, 0
	ans := 0
	n = len(ss)
	for i < n {
		for i < n && ss[i] == 0 {
			i++
		}
		for j < n && tt[j] == 0 {
			j++
		}
		if i < n && j < n {
			ans += abs(i - j)
		}
		i++
		j++
	}

	fmt.Println(ans)
}

func count(slice []int, target int) int {
	count := 0
	for _, num := range slice {
		if num == target {
			count++
		}
	}
	return count
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
