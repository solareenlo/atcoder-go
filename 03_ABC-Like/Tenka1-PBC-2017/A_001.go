package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)

	cnt := 0
	for i := range a {
		if a[i] == '1' {
			cnt++
		}
	}
	fmt.Println(cnt)
}
