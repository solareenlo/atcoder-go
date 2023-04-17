package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	if m%4 != 0 && (m%4 != 3 || n%2 != 0) {
		fmt.Println("Angel")
	} else {
		fmt.Println("Devil")
	}
}
