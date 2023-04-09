package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, d int
	fmt.Fscan(in, &n, &d)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a)
	ans := 0
	for i := n - 1; i >= 0; i-- {
		if (n-i)%d == 0 {
			ans += a[i]
		}
	}
	fmt.Println(ans)
}
