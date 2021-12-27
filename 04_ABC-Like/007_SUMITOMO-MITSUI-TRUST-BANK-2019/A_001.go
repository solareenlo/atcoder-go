package main

import "fmt"

func main() {
	var m1, d1, m2, d2 int
	fmt.Scan(&m1, &d1, &m2, &d2)

	if m1 != m2 {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
