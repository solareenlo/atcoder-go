package main

import "fmt"

func main() {
	var y, m, d int
	fmt.Scan(&y, &m, &d)
	if m == 12 && d == 25 {
		fmt.Println(y - 2018)
	} else {
		fmt.Println("NOT CHRISTMAS DAY")
	}
}
