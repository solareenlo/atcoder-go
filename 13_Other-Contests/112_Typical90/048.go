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

	ab := make([]int, 2*n)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		ab[2*i] = a - b
		ab[2*i+1] = b
	}
	sort.Ints(ab)

	ans := 0
	for i := 0; i < k; i++ {
		ans += ab[2*n-1-i]
	}
	fmt.Println(ans)
}
