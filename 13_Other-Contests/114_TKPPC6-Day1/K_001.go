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

	a := make([]int, n+1)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	d := make([]int, n+1)
	d[0] = a[0]
	for i := 1; i <= n; i++ {
		d[i] = (a[i] - a[i-1] + 10) % 10
	}
	sort.Slice(d, func(i, j int) bool {
		return d[i] > d[j]
	})

	sum := 0
	for i := 0; i < n+1; i++ {
		sum += d[i]
	}
	sum /= 10

	var d2 [100001]int
	for i := 0; i < sum; i++ {
		d2[i] = 10
	}

	ans := 0
	for i := 0; i < n+1; i++ {
		ans += abs(d[i] - d2[i])
	}
	fmt.Println(ans / 2)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
