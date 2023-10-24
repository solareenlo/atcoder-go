package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	if len(s) < 5 {
		fmt.Println("No")
	} else {
		if s[:5] == "MUJIN" {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
