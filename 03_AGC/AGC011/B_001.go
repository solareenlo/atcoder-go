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

	a := make([]int, 100100)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	tmp := a[1 : n+1]
	sort.Ints(tmp)

	ans := 0
	for i := 1; i <= n; i++ {
		a[i] += a[i-1]
		ans++
		if a[i]*2 < a[i+1] {
			ans = 0
		}
	}
	fmt.Println(ans)
}
