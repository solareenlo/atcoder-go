package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(min(a*b, a*c, b*c), max(a*b, a*c, b*c))
}

func min(a ...int) int {
	res := a[0]
	for i := range a {
		if res > a[i] {
			res = a[i]
		}
	}
	return res
}

func max(a ...int) int {
	res := a[0]
	for i := range a {
		if res < a[i] {
			res = a[i]
		}
	}
	return res
}
