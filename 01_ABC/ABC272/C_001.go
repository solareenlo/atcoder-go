package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var a int
	fmt.Fscan(in, &a)

	b := make([]int, a)
	for i := 0; i < a; i++ {
		fmt.Fscan(in, &b[i])
	}
	sort.Ints(b)

	ans := -1
	for i := a - 2; i < a; i++ {
		for j := 0; j < i; j++ {
			if (b[i]+b[j])%2 == 0 {
				ans = max(ans, b[i]+b[j])
			}
		}
	}

	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
