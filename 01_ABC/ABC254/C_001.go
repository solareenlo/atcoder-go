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

	const N = 200005
	e := make([][]int, N)
	for i := 1; i <= n; i++ {
		var j int
		fmt.Fscan(in, &j)
		e[i%k] = append(e[i%k], j)
	}

	for i := 0; i < k; i++ {
		sort.Ints(e[i])
	}

	var a [N]int
	for i := 1; i <= n; i++ {
		a[i] = e[i%k][(i-1)/k]
	}
	for i := 1; i < n; i++ {
		if a[i] > a[i+1] {
			fmt.Println("No")
			os.Exit(0)
		}
	}
	fmt.Println("Yes")

}
