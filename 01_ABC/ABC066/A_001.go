package main

import (
	"fmt"
	"sort"
)

func main() {
	b := make([]int, 3)
	fmt.Scan(&b[0], &b[1], &b[2])
	sort.Ints(b)
	fmt.Println(b[0] + b[1])
}
