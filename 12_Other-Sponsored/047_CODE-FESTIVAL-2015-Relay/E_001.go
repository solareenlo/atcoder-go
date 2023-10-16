package main

import "fmt"

func main() {
	var h2, t2, h1, t1 int
	fmt.Scan(&h2, &t2)
	fmt.Scan(&h1, &t1)
	if h2 >= 12 {
		fmt.Println("No")
		return
	}
	h1 *= 60
	h2 *= 60
	if h1+t1-360-30 <= h2+t2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
