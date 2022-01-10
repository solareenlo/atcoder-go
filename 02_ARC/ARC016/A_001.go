package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	for i := 1; i < n+1; i++ {
		if i != m {
			fmt.Println(i)
			break
		}
	}
}
