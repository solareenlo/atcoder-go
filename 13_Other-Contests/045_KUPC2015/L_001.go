package main

import (
	"fmt"
	"sort"
)

func Nim(a int, b int) int {
	return ((a ^ b) - 1) ^ a ^ b
}

func SG(x []int) int {
	var i, j, o int
	for i = 0; i < len(x); i++ {
		o ^= x[i]
	}
	for i = 0; i < len(x); i++ {
		for j = 0; j < i; j++ {
			o ^= Nim(x[i], x[j])
		}
	}
	return o
}

func main() {
	var n, v int
	var x, y []int
	fmt.Scan(&n)
	for ; n > 0; n-- {
		fmt.Scan(&v)
		if v > 0 {
			x = append(x, v-1)
		} else {
			y = append(y, -(v + 1))
		}
	}
	sort.Ints(x)
	sort.Ints(y)
	if SG(x)^SG(y) == 0 {
		fmt.Println("Ben")
	} else {
		fmt.Println("Alyssa")
	}
}
