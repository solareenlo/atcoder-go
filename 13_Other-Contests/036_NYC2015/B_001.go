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
	tmp := a[1:]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] < tmp[j]
	})
	s := 0
	ans := 0
	for i := 1; i <= n; i++ {
		if s < a[i] {
			s += a[i]
			ans++
		}
	}
	fmt.Println(ans)
}
