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
	v := make([]int, 0)
	for i := 0; i < n/2; i++ {
		v = append(v, a[i]+a[n-1-i])
	}
	sort.Ints(v)
	fmt.Println(v[len(v)-1] - v[0])
}
