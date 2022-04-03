package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	tmp := 0
	if a%10 < 0 {
		tmp = 1
	}
	fmt.Println(a/10 - tmp)
}
