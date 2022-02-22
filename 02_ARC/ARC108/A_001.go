package main

import "fmt"

func main() {
	var s, p int
	fmt.Scan(&s, &p)

	for i := 1; i <= 1000000; i++ {
		if i*(s-i) == p {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")
}
