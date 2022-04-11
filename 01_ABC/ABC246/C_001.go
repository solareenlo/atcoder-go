package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k, x int
	fmt.Fscan(in, &n, &k, &x)

	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		for ; a[i] > x && k > 0; k-- {
			a[i] -= x
		}
	}
	sort.Ints(a)

	ans := 0
	for i := 1; i <= n-k; i++ {
		ans += a[i]
	}
	fmt.Println(ans)
}
