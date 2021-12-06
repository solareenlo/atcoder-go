package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	type st struct{ s, t string }
	m := map[st]bool{}
	for i := 0; i < n; i++ {
		var s, t string
		fmt.Scan(&s, &t)
		m[st{s, t}] = true
	}

	if len(m) == n {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
