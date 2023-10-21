package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a[1:])
	sum := 0
	for i := 1; i <= k; i++ {
		sum += a[i] + i - 1
	}
	fmt.Println(sum)
}
