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
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	sum := 0
	cnt := 0
	for i := 0; i < n; i++ {
		sum += a[i]
		cnt++
		if sum >= m {
			fmt.Println(cnt)
			return
		}
	}
}
