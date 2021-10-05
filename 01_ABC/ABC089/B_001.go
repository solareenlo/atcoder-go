package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	m := map[string]struct{}{}
	var s string
	for i := 0; i < n; i++ {
		fmt.Scan(&s)
		m[s] = struct{}{}
	}

	if len(m) == 4 {
		fmt.Println("Four")
	} else {
		fmt.Println("Three")
	}
}
