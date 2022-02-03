package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	set := map[int]struct{}{}
	var a int
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		set[a] = struct{}{}
	}

	res := len(set)
	if res%2 == 0 {
		res--
	}
	fmt.Println(res)
}
