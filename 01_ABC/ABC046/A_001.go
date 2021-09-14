package main

import "fmt"

func main() {
	set := map[int]struct{}{}
	for i := 0; i < 3; i++ {
		var tmp int
		fmt.Scan(&tmp)
		set[tmp] = struct{}{}
	}
	fmt.Println(len(set))
}
