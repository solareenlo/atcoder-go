package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	t := make([]int, n)
	for i := range t {
		fmt.Fscan(in, &t[i])
	}
	sort.Ints(t)

	sum := 0
	for i := 0; i < n; i++ {
		sum += t[n-1-i] * (i + 1)
	}
	fmt.Println(sum)

	m := map[int]int{}
	for i := 0; i < n; i++ {
		m[t[i]]++
	}

	res := 1
	for _, v := range m {
		for i := 1; i <= v; i++ {
			res *= i
			res %= 1_000_000_007
		}
	}

	fmt.Println(res)
}
