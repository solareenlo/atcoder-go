package main

import "fmt"

func main() {
	a := [13]int{0, 3, 3, 4, 1, 1, 4, 4, 2, 2, 2, 5, 3}
	var n int
	fmt.Scan(&n)
	n /= 2
	fmt.Println(a[n%13])
}
