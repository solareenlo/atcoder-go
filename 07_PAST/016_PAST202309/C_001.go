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
	tb := make([]int, 8)
	for n > 0 {
		n--
		var v int
		fmt.Fscan(in, &v)
		v--
		tb[v]++
	}
	sort.Ints(tb)
	fmt.Println(tb[0])
}
