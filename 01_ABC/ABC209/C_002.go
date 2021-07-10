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
	fmt.Scan(&n)

	c := make([]int, n)
	for i := range c {
		fmt.Fscan(in, &c[i])
	}
	sort.Ints(c)

	res := int64(1)
	for i := range c {
		res = res * int64(c[i]-i) % int64(1000000007)
	}
	fmt.Println(res)
}
