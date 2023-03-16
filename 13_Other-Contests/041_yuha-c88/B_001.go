package main

import "fmt"

func main() {
	var k int
	fmt.Scan(&k)
	if k%4 == 0 {
		fmt.Println("GO")
	} else {
		fmt.Println("SEN")
	}
}
