package main

import "fmt"

func main() {
	var n, a int
	fmt.Scan(&n)

	s := make(map[int]bool)
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		s[a] = true
	}

	if len(s) == n {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
