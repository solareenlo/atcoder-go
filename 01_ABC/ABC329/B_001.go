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
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a)
	mx := a[n-1]
	for i := n - 1; i >= 0; i-- {
		if a[i] != mx {
			fmt.Println(a[i])
			break
		}
	}
}
