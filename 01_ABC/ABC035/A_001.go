package main

import "fmt"

func main() {
	var w, h float64
	fmt.Scan(&w, &h)
	if h/w == 3.0/4.0 {
		fmt.Println("4:3")
	} else {
		fmt.Println("16:9")
	}
}
