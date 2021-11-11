package main

import "fmt"

func main() {
	var d, t, s int
	fmt.Scan(&d, &t, &s)

	T := float64(d) / float64(s)
	if T <= float64(t) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
