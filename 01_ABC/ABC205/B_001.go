package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	s := make(map[int]struct{})
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		s[a] = struct{}{}
	}

	if len(s) == n {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
