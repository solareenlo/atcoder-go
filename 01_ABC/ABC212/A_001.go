package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	res := ""
	if 0 < a && b == 0 {
		res = "Gold"
	} else if a == 0 && 0 < b {
		res = "Silver"
	} else if 0 < a && 0 < b {
		res = "Alloy"
	}
	fmt.Println(res)
}
