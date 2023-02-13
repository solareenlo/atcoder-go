package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	mini := min(a, b, c)
	maxi := max(a, b, c)
	fmt.Println(maxi - mini)
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
