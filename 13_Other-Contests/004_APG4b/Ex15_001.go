package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := input(n)
	b := input(n)
	c := input(n)
	sumA := sum(a)
	sumB := sum(b)
	sumC := sum(c)
	fmt.Println(sumA * sumB * sumC)
}

func input(n int) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&res[i])
	}
	return res
}

func sum(v []int) int {
	sum := 0
	for i := 0; i < len(v); i++ {
		sum += v[i]
	}
	return sum
}
