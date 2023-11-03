package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MAXN = 400400

	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a[1:])
	if m < 0 {
		m = -m
	}
	l := 0
	for i := 1; i <= n; i++ {
		for a[l+1]+m < a[i] {
			l++
		}
		if a[l+1]+m == a[i] {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
