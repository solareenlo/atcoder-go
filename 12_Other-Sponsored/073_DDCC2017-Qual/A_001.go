package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)
	if a[0] == a[1] && a[2] == a[3] && a[0] != a[3] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
