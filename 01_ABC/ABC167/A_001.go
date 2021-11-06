package main

import "fmt"

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	tmp := t[:len(t)-1]

	if s == tmp {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
