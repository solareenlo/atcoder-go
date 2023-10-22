package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, c int
	fmt.Fscan(in, &n, &c)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a[1:])
	ans := 0
	for l, r := 1, n; l <= r; r-- {
		if a[l]+a[r]+1 <= c {
			l++
		}
		ans++
	}
	fmt.Println(ans)
}
