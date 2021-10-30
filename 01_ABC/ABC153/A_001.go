package main

import "fmt"

func main() {
	var h, a int
	fmt.Scan(&h, &a)

	if h%a != 0 {
		fmt.Println(h/a + 1)
	} else {
		fmt.Println(h / a)
	}
}
