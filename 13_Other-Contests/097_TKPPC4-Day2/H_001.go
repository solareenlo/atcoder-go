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

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a)

	ans := 0.0
	for i := 0; i < k; i++ {
		ans += 1.0 + float64(a[i])/float64(a[i+n-k])
	}
	fmt.Println(ans)
}
