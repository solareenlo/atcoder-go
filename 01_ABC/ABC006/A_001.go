package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	res := "NO"
	if n%3 == 0 {
		res = "YES"
	}
	fmt.Println(res)
}
