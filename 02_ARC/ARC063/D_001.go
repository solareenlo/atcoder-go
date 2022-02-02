package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, t, mini int
	fmt.Fscan(in, &n, &t, &mini)

	m := map[int]int{}
	for i := 1; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		if a > mini {
			m[a-mini]++
		}
		mini = min(mini, a)
	}

	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	fmt.Println(m[keys[len(keys)-1]])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
