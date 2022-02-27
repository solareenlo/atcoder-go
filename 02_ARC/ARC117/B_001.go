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
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	tmp := a[1 : n+1]
	sort.Ints(tmp)

	res := 1
	for i := 1; i <= n; i++ {
		res *= (a[i] - a[i-1] + 1)
		res %= 1000000007
	}

	fmt.Println(res)
}
