package main

import "fmt"

func main() {
	var k, a, b int
	fmt.Scan(&k, &a, &b)

	dist := "NG"
	for i := a; i <= b; i++ {
		if i%k == 0 {
			dist = "OK"
			break
		}
	}
	fmt.Println(dist)
}
