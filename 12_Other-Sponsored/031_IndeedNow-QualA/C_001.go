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
	v := make([]int, n)
	for i := range v {
		fmt.Fscan(in, &v[i])
	}
	sort.Slice(v, func(i, j int) bool {
		return v[i] > v[j]
	})
	v = append(v, 0)
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		var s int
		fmt.Fscan(in, &s)
		if v[s] != 0 {
			fmt.Println(v[s] + 1)
		} else {
			fmt.Println(v[s])
		}
	}
}
