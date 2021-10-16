package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	ts := make([]int, 5)
	for i := range ts {
		fmt.Scan(&ts[i])
	}
	sort.Ints(ts)

	if n <= ts[0] {
		fmt.Println(5)
	} else if n%ts[0] != 0 {
		fmt.Println(5 + (n / ts[0]))
	} else {
		fmt.Println(5 + (n / ts[0]) - 1)
	}
}
