package main

import "fmt"

func main() {
	var M int
	fmt.Scan(&M)
	if M <= 9 && M >= 4 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
