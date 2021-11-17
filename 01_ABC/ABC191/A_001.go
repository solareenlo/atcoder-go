package main

import "fmt"

func main() {
	var v, t, s, d float64
	fmt.Scan(&v, &t, &s, &d)

	dv := d / v
	if dv < t || s < dv {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
