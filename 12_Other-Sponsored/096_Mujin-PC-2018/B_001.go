package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	var s string
	fmt.Scan(&s)
	for _, c := range s {
		if a == 0 {
			fmt.Println("Yes")
			return
		}
		a += int(44 - c)
	}
	if a != 0 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
