package main

import "fmt"

func main() {
	var b int
	fmt.Scan(&b)
	cnt := 0
	for b%3 != 2 {
		b = (b*2 + 2) / 3
		cnt++
	}
	fmt.Println(cnt)
}
