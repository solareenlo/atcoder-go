package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	m := map[int]struct{}{}
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		m[a] = struct{}{}
	}

	if len(m) == n {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
