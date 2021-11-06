package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	cnt := 0
	for i := 100; i < x; i += i / 100 {
		cnt++
	}
	fmt.Println(cnt)
}
