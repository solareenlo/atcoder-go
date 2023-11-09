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
	ans := -1
	for i := 0; i < n-1; i++ {
		if abs(a[i]-a[i+1]) > 1 {
			ans = a[i] + 1
			break
		}
	}
	fmt.Println(ans)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
