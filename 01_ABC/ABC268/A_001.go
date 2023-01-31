package main

import "fmt"

func main() {
	m := make(map[int]bool)
	for i := 0; i < 5; i++ {
		var a int
		fmt.Scan(&a)
		m[a] = true
	}
	fmt.Println(len(m))
}
