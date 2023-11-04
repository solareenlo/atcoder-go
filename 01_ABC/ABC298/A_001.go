package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	flag := false
	var S string
	fmt.Scan(&S)
	for _, x := range S {
		if x == 'x' {
			fmt.Println("No")
			return
		}
		if x == 'o' {
			flag = true
		}
	}
	if flag {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
