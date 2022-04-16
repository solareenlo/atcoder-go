package main

import "fmt"

func main() {
	var m string
	fmt.Scan(&m)

	n := "CODEFESTIVAL2016"
	tot := 0
	for i := 0; i < 16; i++ {
		if m[i] != n[i] {
			tot++
		}
	}
	fmt.Println(tot)
}
